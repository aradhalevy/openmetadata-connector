// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package databasetypes

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	zerolog "github.com/rs/zerolog"

	models "fybrik.io/openmetadata-connector/datacatalog-go-models"
	utils "fybrik.io/openmetadata-connector/pkg/utils"
	vault "fybrik.io/openmetadata-connector/pkg/vault"
)

type s3 struct {
	dataBase
	vaultClientConfiguration map[interface{}]interface{}
}

var translate = map[string]string{
	Region:         AwsRegion,
	Endpoint:       EndPointURL,
	AccessKeyID:    AwsAccessKeyID,
	SecretAccessID: AwsSecretAccessKey,
}
var translateInv = map[string]string{
	AwsRegion:          Region,
	EndPointURL:        Endpoint,
	AwsAccessKeyID:     AccessKeyID,
	AwsSecretAccessKey: SecretAccessID,
}

func NewS3(vaultClientConfiguration map[interface{}]interface{}, logger *zerolog.Logger) *s3 {
	return &s3{dataBase: dataBase{name: Datalake, logger: logger},
		vaultClientConfiguration: vaultClientConfiguration}
}

func (s *s3) getS3Credentials(vaultClientConfiguration map[interface{}]interface{},
	credentialsPath *string) (string, string, error) {
	client := vault.NewVaultClient(vaultClientConfiguration, s.logger)
	token, err := client.GetToken()
	if err != nil {
		s.logger.Warn().Msg(GetTokenFailed)
		return "", "", errors.New(GetTokenFailed)
	}
	secret, err := client.GetSecret(token, *credentialsPath)
	if err != nil {
		s.logger.Warn().Msg(GetSecretFailed)
		return "", "", errors.New(GetSecretFailed)
	}
	accessKey, secretKey, err := client.ExtractS3CredentialsFromSecret(secret)
	if err != nil {
		s.logger.Warn().Msg("S3 credentials extraction failed")
		return "", "", err
	}
	return accessKey, secretKey, nil
}

func (s *s3) TranslateFybrikConfigToOpenMetadataConfig(config map[string]interface{}, credentialsPath *string) map[string]interface{} {
	ret := make(map[string]interface{})
	configSourceMap := make(map[string]interface{})
	ret[Type] = s.OMTypeName()
	bucketName, found := config[Bucket]
	if found {
		ret[BucketName] = bucketName
	}

	securityMap := make(map[string]interface{})
	securityMap[AwsRegion] = "eu-de" // awsRegion field is mandatory, although it is persumably ignored if endpoint is provided
	for key, value := range config {
		translation, found := translate[key]
		if found {
			securityMap[translation] = value
		}
	}

	if s.vaultClientConfiguration != nil && credentialsPath != nil {
		awsAccessKeyID, awsSecretAccessKey, err := s.getS3Credentials(s.vaultClientConfiguration, credentialsPath)
		if err == nil && awsAccessKeyID != "" && awsSecretAccessKey != "" {
			securityMap[AwsAccessKeyID] = awsAccessKeyID
			securityMap[AwsSecretAccessKey] = awsSecretAccessKey
		}
	}

	configSourceMap[SecurityConfig] = securityMap
	ret[ConfigSource] = configSourceMap
	return ret
}

func (s *s3) TranslateOpenMetadataConfigToFybrikConfig(tableName string, credentials string,
	config map[string]interface{}) map[string]interface{} {
	ret := make(map[string]interface{})
	ret[ObjectKey] = tableName

	configSource, ok := utils.InterfaceToMap(config[ConfigSource])
	if !ok {
		s.logger.Warn().Msg(fmt.Sprintf(FailedToConvert, ConfigSource))
		return nil
	}
	securityConfig, ok := utils.InterfaceToMap(configSource[SecurityConfig])
	if !ok {
		s.logger.Warn().Msg(fmt.Sprintf(FailedToConvert, SecurityConfig))
		return nil
	}

	for key, value := range securityConfig {
		if translation, found := translateInv[key]; found {
			ret[translation] = value
		}
	}
	if value, found := config[BucketName]; found {
		ret[Bucket] = value
	}

	// if credentials were extracted from vault, we don't want to return
	// the actual access key and secret key
	if credentials != "" {
		delete(ret, AccessKeyID)
		delete(ret, SecretAccessID)
	}

	return ret
}

// Check whether the fields in the 'configSource' sectin are equivalent.
// They don't have to be identical. We allow additional fields (e.g. 'aws_token')
// in the OM configuration, but we require that all fields in the request configuration
// also appear in the OM configuration, and that the values be identical
func (s *s3) equivalentConfigSource(fromService, fromRequest map[string]interface{}) bool {
	// ignore some fields, such as 'aws_token' which would appear only serviceSecurityConfig
	serviceSecurityConfig, ok := utils.InterfaceToMap(fromService[SecurityConfig])
	if !ok {
		s.logger.Warn().Msg(fmt.Sprintf(FailedToConvert, "OM "+SecurityConfig))
		return false
	}
	requestSecurityConfig, ok := utils.InterfaceToMap(fromRequest[SecurityConfig])
	if !ok {
		s.logger.Warn().Msg(fmt.Sprintf(FailedToConvert, "Request "+SecurityConfig))
		return false
	}
	for property, value := range requestSecurityConfig {
		if !reflect.DeepEqual(serviceSecurityConfig[property], value) {
			return false
		}
	}
	return true
}

func (s *s3) EquivalentServiceConfigurations(requestConfig, serviceConfig map[string]interface{}) bool {
	for property, value := range requestConfig {
		if property == ConfigSource {
			servicePropertyMap, ok := utils.InterfaceToMap(serviceConfig[property])
			if !ok {
				s.logger.Warn().Msg(fmt.Sprintf(FailedToConvert, property))
				return false
			}
			valueMap, ok := utils.InterfaceToMap(value)
			if !ok {
				s.logger.Warn().Msg(fmt.Sprintf(FailedToConvert, Value))
				return false
			}
			if !s.equivalentConfigSource(servicePropertyMap, valueMap) {
				return false
			}
		} else if !reflect.DeepEqual(serviceConfig[property], value) {
			return false
		}
	}
	return true
}
func (s *s3) DatabaseName(createAssetRequest *models.CreateAssetRequest) string {
	return Default
}

func (s *s3) DatabaseFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) string {
	return utils.AppendStrings(serviceName, s.DatabaseName(createAssetRequest))
}

func (s *s3) DatabaseSchemaName(createAssetRequest *models.CreateAssetRequest) string {
	connectionProperties, ok := utils.InterfaceToMap(createAssetRequest.Details.GetConnection().
		AdditionalProperties[S3])
	if !ok {
		s.logger.Warn().Msg(fmt.Sprintf(FailedToConvert, AdditionalProperties))
		return ""
	}
	bucket, found := connectionProperties[Bucket]
	if found {
		return bucket.(string)
	}

	assetID := *createAssetRequest.DestinationAssetID
	split := strings.Split(assetID, ".")
	if len(split) > 1 {
		return split[len(split)-2]
	}

	s.logger.Warn().Msg("Could not determine the name of the DatabaseSchema (bucket)")
	return ""
}

func (s *s3) DatabaseSchemaFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) string {
	return utils.AppendStrings(s.DatabaseFQN(serviceName, createAssetRequest),
		s.DatabaseSchemaName(createAssetRequest))
}

func (s *s3) TableName(createAssetRequest *models.CreateAssetRequest) string {
	connectionProperties, ok := utils.InterfaceToMap(createAssetRequest.Details.GetConnection().AdditionalProperties[S3])
	if !ok {
		s.logger.Warn().Msg(fmt.Sprintf(FailedToConvert, AdditionalProperties))
		return ""
	}
	objectKey, found := connectionProperties[ObjectKey]
	if found {
		return objectKey.(string)
	}
	split := strings.Split(*createAssetRequest.DestinationAssetID, ".")
	return split[len(split)-1]
}

func (s *s3) TableFQN(serviceName string, createAssetRequest *models.CreateAssetRequest) string {
	return utils.AppendStrings(s.DatabaseSchemaFQN(serviceName, createAssetRequest),
		s.TableName(createAssetRequest))
}
