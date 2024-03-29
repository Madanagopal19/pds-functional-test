/*
PDS API

Portworx Data Services API Server

API version: 3.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pds

import (
	"encoding/json"
)

// ModelsDeploymentResources struct for ModelsDeploymentResources
type ModelsDeploymentResources struct {
	CpuLimit *string `json:"cpu_limit,omitempty"`
	CpuRequest *string `json:"cpu_request,omitempty"`
	MemoryLimit *string `json:"memory_limit,omitempty"`
	MemoryRequest *string `json:"memory_request,omitempty"`
	StorageRequest *string `json:"storage_request,omitempty"`
}

// NewModelsDeploymentResources instantiates a new ModelsDeploymentResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsDeploymentResources() *ModelsDeploymentResources {
	this := ModelsDeploymentResources{}
	return &this
}

// NewModelsDeploymentResourcesWithDefaults instantiates a new ModelsDeploymentResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsDeploymentResourcesWithDefaults() *ModelsDeploymentResources {
	this := ModelsDeploymentResources{}
	return &this
}

// GetCpuLimit returns the CpuLimit field value if set, zero value otherwise.
func (o *ModelsDeploymentResources) GetCpuLimit() string {
	if o == nil || o.CpuLimit == nil {
		var ret string
		return ret
	}
	return *o.CpuLimit
}

// GetCpuLimitOk returns a tuple with the CpuLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentResources) GetCpuLimitOk() (*string, bool) {
	if o == nil || o.CpuLimit == nil {
		return nil, false
	}
	return o.CpuLimit, true
}

// HasCpuLimit returns a boolean if a field has been set.
func (o *ModelsDeploymentResources) HasCpuLimit() bool {
	if o != nil && o.CpuLimit != nil {
		return true
	}

	return false
}

// SetCpuLimit gets a reference to the given string and assigns it to the CpuLimit field.
func (o *ModelsDeploymentResources) SetCpuLimit(v string) {
	o.CpuLimit = &v
}

// GetCpuRequest returns the CpuRequest field value if set, zero value otherwise.
func (o *ModelsDeploymentResources) GetCpuRequest() string {
	if o == nil || o.CpuRequest == nil {
		var ret string
		return ret
	}
	return *o.CpuRequest
}

// GetCpuRequestOk returns a tuple with the CpuRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentResources) GetCpuRequestOk() (*string, bool) {
	if o == nil || o.CpuRequest == nil {
		return nil, false
	}
	return o.CpuRequest, true
}

// HasCpuRequest returns a boolean if a field has been set.
func (o *ModelsDeploymentResources) HasCpuRequest() bool {
	if o != nil && o.CpuRequest != nil {
		return true
	}

	return false
}

// SetCpuRequest gets a reference to the given string and assigns it to the CpuRequest field.
func (o *ModelsDeploymentResources) SetCpuRequest(v string) {
	o.CpuRequest = &v
}

// GetMemoryLimit returns the MemoryLimit field value if set, zero value otherwise.
func (o *ModelsDeploymentResources) GetMemoryLimit() string {
	if o == nil || o.MemoryLimit == nil {
		var ret string
		return ret
	}
	return *o.MemoryLimit
}

// GetMemoryLimitOk returns a tuple with the MemoryLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentResources) GetMemoryLimitOk() (*string, bool) {
	if o == nil || o.MemoryLimit == nil {
		return nil, false
	}
	return o.MemoryLimit, true
}

// HasMemoryLimit returns a boolean if a field has been set.
func (o *ModelsDeploymentResources) HasMemoryLimit() bool {
	if o != nil && o.MemoryLimit != nil {
		return true
	}

	return false
}

// SetMemoryLimit gets a reference to the given string and assigns it to the MemoryLimit field.
func (o *ModelsDeploymentResources) SetMemoryLimit(v string) {
	o.MemoryLimit = &v
}

// GetMemoryRequest returns the MemoryRequest field value if set, zero value otherwise.
func (o *ModelsDeploymentResources) GetMemoryRequest() string {
	if o == nil || o.MemoryRequest == nil {
		var ret string
		return ret
	}
	return *o.MemoryRequest
}

// GetMemoryRequestOk returns a tuple with the MemoryRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentResources) GetMemoryRequestOk() (*string, bool) {
	if o == nil || o.MemoryRequest == nil {
		return nil, false
	}
	return o.MemoryRequest, true
}

// HasMemoryRequest returns a boolean if a field has been set.
func (o *ModelsDeploymentResources) HasMemoryRequest() bool {
	if o != nil && o.MemoryRequest != nil {
		return true
	}

	return false
}

// SetMemoryRequest gets a reference to the given string and assigns it to the MemoryRequest field.
func (o *ModelsDeploymentResources) SetMemoryRequest(v string) {
	o.MemoryRequest = &v
}

// GetStorageRequest returns the StorageRequest field value if set, zero value otherwise.
func (o *ModelsDeploymentResources) GetStorageRequest() string {
	if o == nil || o.StorageRequest == nil {
		var ret string
		return ret
	}
	return *o.StorageRequest
}

// GetStorageRequestOk returns a tuple with the StorageRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentResources) GetStorageRequestOk() (*string, bool) {
	if o == nil || o.StorageRequest == nil {
		return nil, false
	}
	return o.StorageRequest, true
}

// HasStorageRequest returns a boolean if a field has been set.
func (o *ModelsDeploymentResources) HasStorageRequest() bool {
	if o != nil && o.StorageRequest != nil {
		return true
	}

	return false
}

// SetStorageRequest gets a reference to the given string and assigns it to the StorageRequest field.
func (o *ModelsDeploymentResources) SetStorageRequest(v string) {
	o.StorageRequest = &v
}

func (o ModelsDeploymentResources) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CpuLimit != nil {
		toSerialize["cpu_limit"] = o.CpuLimit
	}
	if o.CpuRequest != nil {
		toSerialize["cpu_request"] = o.CpuRequest
	}
	if o.MemoryLimit != nil {
		toSerialize["memory_limit"] = o.MemoryLimit
	}
	if o.MemoryRequest != nil {
		toSerialize["memory_request"] = o.MemoryRequest
	}
	if o.StorageRequest != nil {
		toSerialize["storage_request"] = o.StorageRequest
	}
	return json.Marshal(toSerialize)
}

type NullableModelsDeploymentResources struct {
	value *ModelsDeploymentResources
	isSet bool
}

func (v NullableModelsDeploymentResources) Get() *ModelsDeploymentResources {
	return v.value
}

func (v *NullableModelsDeploymentResources) Set(val *ModelsDeploymentResources) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsDeploymentResources) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsDeploymentResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsDeploymentResources(val *ModelsDeploymentResources) *NullableModelsDeploymentResources {
	return &NullableModelsDeploymentResources{value: val, isSet: true}
}

func (v NullableModelsDeploymentResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsDeploymentResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


