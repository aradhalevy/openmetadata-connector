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

// UsageStats struct for UsageStats
type UsageStats struct {
	Count          int32    `json:"count"`
	PercentileRank *float64 `json:"percentileRank,omitempty"`
}

// NewUsageStats instantiates a new UsageStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageStats(count int32) *UsageStats {
	this := UsageStats{}
	this.Count = count
	return &this
}

// NewUsageStatsWithDefaults instantiates a new UsageStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageStatsWithDefaults() *UsageStats {
	this := UsageStats{}
	return &this
}

// GetCount returns the Count field value
func (o *UsageStats) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *UsageStats) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *UsageStats) SetCount(v int32) {
	o.Count = v
}

// GetPercentileRank returns the PercentileRank field value if set, zero value otherwise.
func (o *UsageStats) GetPercentileRank() float64 {
	if o == nil || o.PercentileRank == nil {
		var ret float64
		return ret
	}
	return *o.PercentileRank
}

// GetPercentileRankOk returns a tuple with the PercentileRank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageStats) GetPercentileRankOk() (*float64, bool) {
	if o == nil || o.PercentileRank == nil {
		return nil, false
	}
	return o.PercentileRank, true
}

// HasPercentileRank returns a boolean if a field has been set.
func (o *UsageStats) HasPercentileRank() bool {
	if o != nil && o.PercentileRank != nil {
		return true
	}

	return false
}

// SetPercentileRank gets a reference to the given float64 and assigns it to the PercentileRank field.
func (o *UsageStats) SetPercentileRank(v float64) {
	o.PercentileRank = &v
}

func (o UsageStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	if o.PercentileRank != nil {
		toSerialize["percentileRank"] = o.PercentileRank
	}
	return json.Marshal(toSerialize)
}

type NullableUsageStats struct {
	value *UsageStats
	isSet bool
}

func (v NullableUsageStats) Get() *UsageStats {
	return v.value
}

func (v *NullableUsageStats) Set(val *UsageStats) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageStats) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageStats(val *UsageStats) *NullableUsageStats {
	return &NullableUsageStats{value: val, isSet: true}
}

func (v NullableUsageStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}