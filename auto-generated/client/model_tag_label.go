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

// TagLabel struct for TagLabel
type TagLabel struct {
	Description *string `json:"description,omitempty"`
	Href        *string `json:"href,omitempty"`
	LabelType   string  `json:"labelType"`
	Source      string  `json:"source"`
	State       string  `json:"state"`
	TagFQN      string  `json:"tagFQN"`
}

// NewTagLabel instantiates a new TagLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagLabel(labelType string, source string, state string, tagFQN string) *TagLabel {
	this := TagLabel{}
	this.LabelType = labelType
	this.Source = source
	this.State = state
	this.TagFQN = tagFQN
	return &this
}

// NewTagLabelWithDefaults instantiates a new TagLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagLabelWithDefaults() *TagLabel {
	this := TagLabel{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TagLabel) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagLabel) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TagLabel) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TagLabel) SetDescription(v string) {
	o.Description = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *TagLabel) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagLabel) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *TagLabel) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *TagLabel) SetHref(v string) {
	o.Href = &v
}

// GetLabelType returns the LabelType field value
func (o *TagLabel) GetLabelType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LabelType
}

// GetLabelTypeOk returns a tuple with the LabelType field value
// and a boolean to check if the value has been set.
func (o *TagLabel) GetLabelTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelType, true
}

// SetLabelType sets field value
func (o *TagLabel) SetLabelType(v string) {
	o.LabelType = v
}

// GetSource returns the Source field value
func (o *TagLabel) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *TagLabel) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *TagLabel) SetSource(v string) {
	o.Source = v
}

// GetState returns the State field value
func (o *TagLabel) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *TagLabel) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *TagLabel) SetState(v string) {
	o.State = v
}

// GetTagFQN returns the TagFQN field value
func (o *TagLabel) GetTagFQN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TagFQN
}

// GetTagFQNOk returns a tuple with the TagFQN field value
// and a boolean to check if the value has been set.
func (o *TagLabel) GetTagFQNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagFQN, true
}

// SetTagFQN sets field value
func (o *TagLabel) SetTagFQN(v string) {
	o.TagFQN = v
}

func (o TagLabel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if true {
		toSerialize["labelType"] = o.LabelType
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["tagFQN"] = o.TagFQN
	}
	return json.Marshal(toSerialize)
}

type NullableTagLabel struct {
	value *TagLabel
	isSet bool
}

func (v NullableTagLabel) Get() *TagLabel {
	return v.value
}

func (v *NullableTagLabel) Set(val *TagLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableTagLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableTagLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagLabel(val *TagLabel) *NullableTagLabel {
	return &NullableTagLabel{value: val, isSet: true}
}

func (v NullableTagLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}