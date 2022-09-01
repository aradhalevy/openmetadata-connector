/*
OpenMetadata Apis

# Overview  OpenMetadata supports REST APIs for getting metadata in and out of metadata store. The API resources are grouped under following categories: - **Data assets** - includes resources for data entities, such as `databases`, `tables`, and `topics`. Resources for data assets created from data, such as `dashboards`, `reports`, `metrics`, and `ML Features` are part of this collection. `pipelines`, `notebooks`, etc. that are used for creating data assets are also available as resources as of this collection. - **Teams and Users** - includes `users`, `teams`, a special type of user called `bots` that performs many automated tasks such as ingestion. - **Services** - are services that OpenMetadata integrates with. Currently `databaseService` is the only service under this collection that represents data sources. In the future, services related to Dashboards, Reports, ETL pipelines will be added under this collection. - **Glossary** - OpenMetadata supports hierarchical tags that can be used to build business vocabulary to describe and classify data available under `tags` resource.

API version: 0.11.0
Contact: openmetadata-dev@googlegroups.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// FieldChange struct for FieldChange
type FieldChange struct {
	Name     *string      `json:"name,omitempty"`
	NewValue *interface{} `json:"newValue,omitempty"`
	OldValue *interface{} `json:"oldValue,omitempty"`
}

// NewFieldChange instantiates a new FieldChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldChange() *FieldChange {
	this := FieldChange{}
	return &this
}

// NewFieldChangeWithDefaults instantiates a new FieldChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldChangeWithDefaults() *FieldChange {
	this := FieldChange{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FieldChange) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldChange) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FieldChange) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FieldChange) SetName(v string) {
	o.Name = &v
}

// GetNewValue returns the NewValue field value if set, zero value otherwise.
func (o *FieldChange) GetNewValue() interface{} {
	if o == nil || o.NewValue == nil {
		var ret interface{}
		return ret
	}
	return *o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldChange) GetNewValueOk() (*interface{}, bool) {
	if o == nil || o.NewValue == nil {
		return nil, false
	}
	return o.NewValue, true
}

// HasNewValue returns a boolean if a field has been set.
func (o *FieldChange) HasNewValue() bool {
	if o != nil && o.NewValue != nil {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given interface{} and assigns it to the NewValue field.
func (o *FieldChange) SetNewValue(v interface{}) {
	o.NewValue = &v
}

// GetOldValue returns the OldValue field value if set, zero value otherwise.
func (o *FieldChange) GetOldValue() interface{} {
	if o == nil || o.OldValue == nil {
		var ret interface{}
		return ret
	}
	return *o.OldValue
}

// GetOldValueOk returns a tuple with the OldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldChange) GetOldValueOk() (*interface{}, bool) {
	if o == nil || o.OldValue == nil {
		return nil, false
	}
	return o.OldValue, true
}

// HasOldValue returns a boolean if a field has been set.
func (o *FieldChange) HasOldValue() bool {
	if o != nil && o.OldValue != nil {
		return true
	}

	return false
}

// SetOldValue gets a reference to the given interface{} and assigns it to the OldValue field.
func (o *FieldChange) SetOldValue(v interface{}) {
	o.OldValue = &v
}

func (o FieldChange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NewValue != nil {
		toSerialize["newValue"] = o.NewValue
	}
	if o.OldValue != nil {
		toSerialize["oldValue"] = o.OldValue
	}
	return json.Marshal(toSerialize)
}

type NullableFieldChange struct {
	value *FieldChange
	isSet bool
}

func (v NullableFieldChange) Get() *FieldChange {
	return v.value
}

func (v *NullableFieldChange) Set(val *FieldChange) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldChange) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldChange(val *FieldChange) *NullableFieldChange {
	return &NullableFieldChange{value: val, isSet: true}
}

func (v NullableFieldChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}