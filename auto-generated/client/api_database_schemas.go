/*
OpenMetadata Apis

# Overview  OpenMetadata supports REST APIs for getting metadata in and out of metadata store. The API resources are grouped under following categories: - **Data assets** - includes resources for data entities, such as `databases`, `tables`, and `topics`. Resources for data assets created from data, such as `dashboards`, `reports`, `metrics`, and `ML Features` are part of this collection. `pipelines`, `notebooks`, etc. that are used for creating data assets are also available as resources as of this collection. - **Teams and Users** - includes `users`, `teams`, a special type of user called `bots` that performs many automated tasks such as ingestion. - **Services** - are services that OpenMetadata integrates with. Currently `databaseService` is the only service under this collection that represents data sources. In the future, services related to Dashboards, Reports, ETL pipelines will be added under this collection. - **Glossary** - OpenMetadata supports hierarchical tags that can be used to build business vocabulary to describe and classify data available under `tags` resource.

API version: 0.11.0
Contact: openmetadata-dev@googlegroups.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// DatabaseSchemasApiService DatabaseSchemasApi service
type DatabaseSchemasApiService service

type ApiCreateDBSchemaRequest struct {
	ctx                  context.Context
	ApiService           *DatabaseSchemasApiService
	createDatabaseSchema *CreateDatabaseSchema
}

func (r ApiCreateDBSchemaRequest) CreateDatabaseSchema(createDatabaseSchema CreateDatabaseSchema) ApiCreateDBSchemaRequest {
	r.createDatabaseSchema = &createDatabaseSchema
	return r
}

func (r ApiCreateDBSchemaRequest) Execute() (*DatabaseSchema, *http.Response, error) {
	return r.ApiService.CreateDBSchemaExecute(r)
}

/*
CreateDBSchema Create a schema

Create a schema under an existing `service`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateDBSchemaRequest
*/
func (a *DatabaseSchemasApiService) CreateDBSchema(ctx context.Context) ApiCreateDBSchemaRequest {
	return ApiCreateDBSchemaRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return DatabaseSchema
func (a *DatabaseSchemasApiService) CreateDBSchemaExecute(r ApiCreateDBSchemaRequest) (*DatabaseSchema, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DatabaseSchema
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseSchemasApiService.CreateDBSchema")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/databaseSchemas"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createDatabaseSchema
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateOrUpdateDBSchemaRequest struct {
	ctx                  context.Context
	ApiService           *DatabaseSchemasApiService
	createDatabaseSchema *CreateDatabaseSchema
}

func (r ApiCreateOrUpdateDBSchemaRequest) CreateDatabaseSchema(createDatabaseSchema CreateDatabaseSchema) ApiCreateOrUpdateDBSchemaRequest {
	r.createDatabaseSchema = &createDatabaseSchema
	return r
}

func (r ApiCreateOrUpdateDBSchemaRequest) Execute() (*DatabaseSchema, *http.Response, error) {
	return r.ApiService.CreateOrUpdateDBSchemaExecute(r)
}

/*
CreateOrUpdateDBSchema Create or update schema

Create a database schema, if it does not exist or update an existing database schema.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateOrUpdateDBSchemaRequest
*/
func (a *DatabaseSchemasApiService) CreateOrUpdateDBSchema(ctx context.Context) ApiCreateOrUpdateDBSchemaRequest {
	return ApiCreateOrUpdateDBSchemaRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return DatabaseSchema
func (a *DatabaseSchemasApiService) CreateOrUpdateDBSchemaExecute(r ApiCreateOrUpdateDBSchemaRequest) (*DatabaseSchema, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DatabaseSchema
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseSchemasApiService.CreateOrUpdateDBSchema")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/databaseSchemas"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createDatabaseSchema
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteDBSchemaRequest struct {
	ctx        context.Context
	ApiService *DatabaseSchemasApiService
	id         string
	recursive  *bool
	hardDelete *bool
}

// Recursively delete this entity and it&#39;s children. (Default &#x60;false&#x60;)
func (r ApiDeleteDBSchemaRequest) Recursive(recursive bool) ApiDeleteDBSchemaRequest {
	r.recursive = &recursive
	return r
}

// Hard delete the entity. (Default &#x3D; &#x60;false&#x60;)
func (r ApiDeleteDBSchemaRequest) HardDelete(hardDelete bool) ApiDeleteDBSchemaRequest {
	r.hardDelete = &hardDelete
	return r
}

func (r ApiDeleteDBSchemaRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteDBSchemaExecute(r)
}

/*
DeleteDBSchema Delete a schema

Delete a schema by `id`. Schema can only be deleted if it has no tables.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiDeleteDBSchemaRequest
*/
func (a *DatabaseSchemasApiService) DeleteDBSchema(ctx context.Context, id string) ApiDeleteDBSchemaRequest {
	return ApiDeleteDBSchemaRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *DatabaseSchemasApiService) DeleteDBSchemaExecute(r ApiDeleteDBSchemaRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseSchemasApiService.DeleteDBSchema")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/databaseSchemas/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.recursive != nil {
		localVarQueryParams.Add("recursive", parameterToString(*r.recursive, ""))
	}
	if r.hardDelete != nil {
		localVarQueryParams.Add("hardDelete", parameterToString(*r.hardDelete, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetDBSchemaByFQNRequest struct {
	ctx        context.Context
	ApiService *DatabaseSchemasApiService
	fqn        string
	fields     *string
	include    *string
}

// Fields requested in the returned resource
func (r ApiGetDBSchemaByFQNRequest) Fields(fields string) ApiGetDBSchemaByFQNRequest {
	r.fields = &fields
	return r
}

// Include all, deleted, or non-deleted entities.
func (r ApiGetDBSchemaByFQNRequest) Include(include string) ApiGetDBSchemaByFQNRequest {
	r.include = &include
	return r
}

func (r ApiGetDBSchemaByFQNRequest) Execute() (*DatabaseSchema, *http.Response, error) {
	return r.ApiService.GetDBSchemaByFQNExecute(r)
}

/*
GetDBSchemaByFQN Get a schema by name

Get a database schema by fully qualified name.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fqn
 @return ApiGetDBSchemaByFQNRequest
*/
func (a *DatabaseSchemasApiService) GetDBSchemaByFQN(ctx context.Context, fqn string) ApiGetDBSchemaByFQNRequest {
	return ApiGetDBSchemaByFQNRequest{
		ApiService: a,
		ctx:        ctx,
		fqn:        fqn,
	}
}

// Execute executes the request
//  @return DatabaseSchema
func (a *DatabaseSchemasApiService) GetDBSchemaByFQNExecute(r ApiGetDBSchemaByFQNRequest) (*DatabaseSchema, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DatabaseSchema
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseSchemasApiService.GetDBSchemaByFQN")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/databaseSchemas/name/{fqn}"
	localVarPath = strings.Replace(localVarPath, "{"+"fqn"+"}", url.PathEscape(parameterToString(r.fqn, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
	}
	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetDBSchemaByIDRequest struct {
	ctx        context.Context
	ApiService *DatabaseSchemasApiService
	id         string
	fields     *string
	include    *string
}

// Fields requested in the returned resource
func (r ApiGetDBSchemaByIDRequest) Fields(fields string) ApiGetDBSchemaByIDRequest {
	r.fields = &fields
	return r
}

// Include all, deleted, or non-deleted entities.
func (r ApiGetDBSchemaByIDRequest) Include(include string) ApiGetDBSchemaByIDRequest {
	r.include = &include
	return r
}

func (r ApiGetDBSchemaByIDRequest) Execute() (*DatabaseSchema, *http.Response, error) {
	return r.ApiService.GetDBSchemaByIDExecute(r)
}

/*
GetDBSchemaByID Get a schema

Get a database schema by `id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetDBSchemaByIDRequest
*/
func (a *DatabaseSchemasApiService) GetDBSchemaByID(ctx context.Context, id string) ApiGetDBSchemaByIDRequest {
	return ApiGetDBSchemaByIDRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//  @return DatabaseSchema
func (a *DatabaseSchemasApiService) GetDBSchemaByIDExecute(r ApiGetDBSchemaByIDRequest) (*DatabaseSchema, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DatabaseSchema
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseSchemasApiService.GetDBSchemaByID")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/databaseSchemas/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
	}
	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetSpecificDBSchemaVersionRequest struct {
	ctx        context.Context
	ApiService *DatabaseSchemasApiService
	id         string
	version    string
}

func (r ApiGetSpecificDBSchemaVersionRequest) Execute() (*DatabaseSchema, *http.Response, error) {
	return r.ApiService.GetSpecificDBSchemaVersionExecute(r)
}

/*
GetSpecificDBSchemaVersion Get a version of the schema

Get a version of the database schema by given `id`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Database schema Id
 @param version Database schema version number in the form `major`.`minor`
 @return ApiGetSpecificDBSchemaVersionRequest
*/
func (a *DatabaseSchemasApiService) GetSpecificDBSchemaVersion(ctx context.Context, id string, version string) ApiGetSpecificDBSchemaVersionRequest {
	return ApiGetSpecificDBSchemaVersionRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
		version:    version,
	}
}

// Execute executes the request
//  @return DatabaseSchema
func (a *DatabaseSchemasApiService) GetSpecificDBSchemaVersionExecute(r ApiGetSpecificDBSchemaVersionRequest) (*DatabaseSchema, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DatabaseSchema
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseSchemasApiService.GetSpecificDBSchemaVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/databaseSchemas/{id}/versions/{version}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", url.PathEscape(parameterToString(r.version, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListAllDBSchemaVersionRequest struct {
	ctx        context.Context
	ApiService *DatabaseSchemasApiService
	id         string
}

func (r ApiListAllDBSchemaVersionRequest) Execute() (*EntityHistory, *http.Response, error) {
	return r.ApiService.ListAllDBSchemaVersionExecute(r)
}

/*
ListAllDBSchemaVersion List schema versions

Get a list of all the versions of a schema identified by `id`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Database schema Id
 @return ApiListAllDBSchemaVersionRequest
*/
func (a *DatabaseSchemasApiService) ListAllDBSchemaVersion(ctx context.Context, id string) ApiListAllDBSchemaVersionRequest {
	return ApiListAllDBSchemaVersionRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//  @return EntityHistory
func (a *DatabaseSchemasApiService) ListAllDBSchemaVersionExecute(r ApiListAllDBSchemaVersionRequest) (*EntityHistory, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EntityHistory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseSchemasApiService.ListAllDBSchemaVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/databaseSchemas/{id}/versions"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListDBSchemasRequest struct {
	ctx        context.Context
	ApiService *DatabaseSchemasApiService
	fields     *string
	database   *string
	limit      *int32
	before     *string
	after      *string
	include    *string
}

// Fields requested in the returned resource
func (r ApiListDBSchemasRequest) Fields(fields string) ApiListDBSchemasRequest {
	r.fields = &fields
	return r
}

// Filter schemas by database name
func (r ApiListDBSchemasRequest) Database(database string) ApiListDBSchemasRequest {
	r.database = &database
	return r
}

// Limit the number schemas returned. (1 to 1000000, default &#x3D; 10)
func (r ApiListDBSchemasRequest) Limit(limit int32) ApiListDBSchemasRequest {
	r.limit = &limit
	return r
}

// Returns list of schemas before this cursor
func (r ApiListDBSchemasRequest) Before(before string) ApiListDBSchemasRequest {
	r.before = &before
	return r
}

// Returns list of schemas after this cursor
func (r ApiListDBSchemasRequest) After(after string) ApiListDBSchemasRequest {
	r.after = &after
	return r
}

// Include all, deleted, or non-deleted entities.
func (r ApiListDBSchemasRequest) Include(include string) ApiListDBSchemasRequest {
	r.include = &include
	return r
}

func (r ApiListDBSchemasRequest) Execute() (*DatabaseSchemaList, *http.Response, error) {
	return r.ApiService.ListDBSchemasExecute(r)
}

/*
ListDBSchemas List database schemas

Get a list of database schemas, optionally filtered by `database` it belongs to. Use `fields` parameter to get only necessary fields. Use cursor-based pagination to limit the number entries in the list using `limit` and `before` or `after` query params.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListDBSchemasRequest
*/
func (a *DatabaseSchemasApiService) ListDBSchemas(ctx context.Context) ApiListDBSchemasRequest {
	return ApiListDBSchemasRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return DatabaseSchemaList
func (a *DatabaseSchemasApiService) ListDBSchemasExecute(r ApiListDBSchemasRequest) (*DatabaseSchemaList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DatabaseSchemaList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseSchemasApiService.ListDBSchemas")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/databaseSchemas"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
	}
	if r.database != nil {
		localVarQueryParams.Add("database", parameterToString(*r.database, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.before != nil {
		localVarQueryParams.Add("before", parameterToString(*r.before, ""))
	}
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPatchDBSchemaRequest struct {
	ctx        context.Context
	ApiService *DatabaseSchemasApiService
	id         string
	body       *map[string]interface{}
}

// JsonPatch with array of operations
func (r ApiPatchDBSchemaRequest) Body(body map[string]interface{}) ApiPatchDBSchemaRequest {
	r.body = &body
	return r
}

func (r ApiPatchDBSchemaRequest) Execute() (*http.Response, error) {
	return r.ApiService.PatchDBSchemaExecute(r)
}

/*
PatchDBSchema Update a database schema

Update an existing database schema using JsonPatch.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiPatchDBSchemaRequest
*/
func (a *DatabaseSchemasApiService) PatchDBSchema(ctx context.Context, id string) ApiPatchDBSchemaRequest {
	return ApiPatchDBSchemaRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *DatabaseSchemasApiService) PatchDBSchemaExecute(r ApiPatchDBSchemaRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseSchemasApiService.PatchDBSchema")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/databaseSchemas/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-patch+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}