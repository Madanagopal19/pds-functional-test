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

// ControllersWhoAmIResponse struct for ControllersWhoAmIResponse
type ControllersWhoAmIResponse struct {
	ServiceAccount *ControllersWhoAmIServiceAccount `json:"service_account,omitempty"`
	User *ControllersWhoAmIUser `json:"user,omitempty"`
}

// NewControllersWhoAmIResponse instantiates a new ControllersWhoAmIResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersWhoAmIResponse() *ControllersWhoAmIResponse {
	this := ControllersWhoAmIResponse{}
	return &this
}

// NewControllersWhoAmIResponseWithDefaults instantiates a new ControllersWhoAmIResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersWhoAmIResponseWithDefaults() *ControllersWhoAmIResponse {
	this := ControllersWhoAmIResponse{}
	return &this
}

// GetServiceAccount returns the ServiceAccount field value if set, zero value otherwise.
func (o *ControllersWhoAmIResponse) GetServiceAccount() ControllersWhoAmIServiceAccount {
	if o == nil || o.ServiceAccount == nil {
		var ret ControllersWhoAmIServiceAccount
		return ret
	}
	return *o.ServiceAccount
}

// GetServiceAccountOk returns a tuple with the ServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersWhoAmIResponse) GetServiceAccountOk() (*ControllersWhoAmIServiceAccount, bool) {
	if o == nil || o.ServiceAccount == nil {
		return nil, false
	}
	return o.ServiceAccount, true
}

// HasServiceAccount returns a boolean if a field has been set.
func (o *ControllersWhoAmIResponse) HasServiceAccount() bool {
	if o != nil && o.ServiceAccount != nil {
		return true
	}

	return false
}

// SetServiceAccount gets a reference to the given ControllersWhoAmIServiceAccount and assigns it to the ServiceAccount field.
func (o *ControllersWhoAmIResponse) SetServiceAccount(v ControllersWhoAmIServiceAccount) {
	o.ServiceAccount = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ControllersWhoAmIResponse) GetUser() ControllersWhoAmIUser {
	if o == nil || o.User == nil {
		var ret ControllersWhoAmIUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersWhoAmIResponse) GetUserOk() (*ControllersWhoAmIUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ControllersWhoAmIResponse) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given ControllersWhoAmIUser and assigns it to the User field.
func (o *ControllersWhoAmIResponse) SetUser(v ControllersWhoAmIUser) {
	o.User = &v
}

func (o ControllersWhoAmIResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceAccount != nil {
		toSerialize["service_account"] = o.ServiceAccount
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableControllersWhoAmIResponse struct {
	value *ControllersWhoAmIResponse
	isSet bool
}

func (v NullableControllersWhoAmIResponse) Get() *ControllersWhoAmIResponse {
	return v.value
}

func (v *NullableControllersWhoAmIResponse) Set(val *ControllersWhoAmIResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersWhoAmIResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersWhoAmIResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersWhoAmIResponse(val *ControllersWhoAmIResponse) *NullableControllersWhoAmIResponse {
	return &NullableControllersWhoAmIResponse{value: val, isSet: true}
}

func (v NullableControllersWhoAmIResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersWhoAmIResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


