/*
Data Catalog Service - Asset Details

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// UpdateAssetResponse struct for UpdateAssetResponse
type UpdateAssetResponse struct {
	// The updation status
	Status *string `json:"status,omitempty"`
}

// NewUpdateAssetResponse instantiates a new UpdateAssetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAssetResponse() *UpdateAssetResponse {
	this := UpdateAssetResponse{}
	return &this
}

// NewUpdateAssetResponseWithDefaults instantiates a new UpdateAssetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAssetResponseWithDefaults() *UpdateAssetResponse {
	this := UpdateAssetResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpdateAssetResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAssetResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdateAssetResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UpdateAssetResponse) SetStatus(v string) {
	o.Status = &v
}

func (o UpdateAssetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateAssetResponse struct {
	value *UpdateAssetResponse
	isSet bool
}

func (v NullableUpdateAssetResponse) Get() *UpdateAssetResponse {
	return v.value
}

func (v *NullableUpdateAssetResponse) Set(val *UpdateAssetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAssetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAssetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAssetResponse(val *UpdateAssetResponse) *NullableUpdateAssetResponse {
	return &NullableUpdateAssetResponse{value: val, isSet: true}
}

func (v NullableUpdateAssetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAssetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}