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

// ControllersUpsertTenantRoleBindingRequest struct for ControllersUpsertTenantRoleBindingRequest
type ControllersUpsertTenantRoleBindingRequest struct {
	ActorId *string `json:"actor_id,omitempty"`
	ActorType *string `json:"actor_type,omitempty"`
	RoleName *string `json:"role_name,omitempty"`
}

// NewControllersUpsertTenantRoleBindingRequest instantiates a new ControllersUpsertTenantRoleBindingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersUpsertTenantRoleBindingRequest() *ControllersUpsertTenantRoleBindingRequest {
	this := ControllersUpsertTenantRoleBindingRequest{}
	return &this
}

// NewControllersUpsertTenantRoleBindingRequestWithDefaults instantiates a new ControllersUpsertTenantRoleBindingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersUpsertTenantRoleBindingRequestWithDefaults() *ControllersUpsertTenantRoleBindingRequest {
	this := ControllersUpsertTenantRoleBindingRequest{}
	return &this
}

// GetActorId returns the ActorId field value if set, zero value otherwise.
func (o *ControllersUpsertTenantRoleBindingRequest) GetActorId() string {
	if o == nil || o.ActorId == nil {
		var ret string
		return ret
	}
	return *o.ActorId
}

// GetActorIdOk returns a tuple with the ActorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersUpsertTenantRoleBindingRequest) GetActorIdOk() (*string, bool) {
	if o == nil || o.ActorId == nil {
		return nil, false
	}
	return o.ActorId, true
}

// HasActorId returns a boolean if a field has been set.
func (o *ControllersUpsertTenantRoleBindingRequest) HasActorId() bool {
	if o != nil && o.ActorId != nil {
		return true
	}

	return false
}

// SetActorId gets a reference to the given string and assigns it to the ActorId field.
func (o *ControllersUpsertTenantRoleBindingRequest) SetActorId(v string) {
	o.ActorId = &v
}

// GetActorType returns the ActorType field value if set, zero value otherwise.
func (o *ControllersUpsertTenantRoleBindingRequest) GetActorType() string {
	if o == nil || o.ActorType == nil {
		var ret string
		return ret
	}
	return *o.ActorType
}

// GetActorTypeOk returns a tuple with the ActorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersUpsertTenantRoleBindingRequest) GetActorTypeOk() (*string, bool) {
	if o == nil || o.ActorType == nil {
		return nil, false
	}
	return o.ActorType, true
}

// HasActorType returns a boolean if a field has been set.
func (o *ControllersUpsertTenantRoleBindingRequest) HasActorType() bool {
	if o != nil && o.ActorType != nil {
		return true
	}

	return false
}

// SetActorType gets a reference to the given string and assigns it to the ActorType field.
func (o *ControllersUpsertTenantRoleBindingRequest) SetActorType(v string) {
	o.ActorType = &v
}

// GetRoleName returns the RoleName field value if set, zero value otherwise.
func (o *ControllersUpsertTenantRoleBindingRequest) GetRoleName() string {
	if o == nil || o.RoleName == nil {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersUpsertTenantRoleBindingRequest) GetRoleNameOk() (*string, bool) {
	if o == nil || o.RoleName == nil {
		return nil, false
	}
	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *ControllersUpsertTenantRoleBindingRequest) HasRoleName() bool {
	if o != nil && o.RoleName != nil {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *ControllersUpsertTenantRoleBindingRequest) SetRoleName(v string) {
	o.RoleName = &v
}

func (o ControllersUpsertTenantRoleBindingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActorId != nil {
		toSerialize["actor_id"] = o.ActorId
	}
	if o.ActorType != nil {
		toSerialize["actor_type"] = o.ActorType
	}
	if o.RoleName != nil {
		toSerialize["role_name"] = o.RoleName
	}
	return json.Marshal(toSerialize)
}

type NullableControllersUpsertTenantRoleBindingRequest struct {
	value *ControllersUpsertTenantRoleBindingRequest
	isSet bool
}

func (v NullableControllersUpsertTenantRoleBindingRequest) Get() *ControllersUpsertTenantRoleBindingRequest {
	return v.value
}

func (v *NullableControllersUpsertTenantRoleBindingRequest) Set(val *ControllersUpsertTenantRoleBindingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersUpsertTenantRoleBindingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersUpsertTenantRoleBindingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersUpsertTenantRoleBindingRequest(val *ControllersUpsertTenantRoleBindingRequest) *NullableControllersUpsertTenantRoleBindingRequest {
	return &NullableControllersUpsertTenantRoleBindingRequest{value: val, isSet: true}
}

func (v NullableControllersUpsertTenantRoleBindingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersUpsertTenantRoleBindingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


