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

// ControllersRefreshTokenRequest struct for ControllersRefreshTokenRequest
type ControllersRefreshTokenRequest struct {
	RefreshToken *string `json:"refreshToken,omitempty"`
}

// NewControllersRefreshTokenRequest instantiates a new ControllersRefreshTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersRefreshTokenRequest() *ControllersRefreshTokenRequest {
	this := ControllersRefreshTokenRequest{}
	return &this
}

// NewControllersRefreshTokenRequestWithDefaults instantiates a new ControllersRefreshTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersRefreshTokenRequestWithDefaults() *ControllersRefreshTokenRequest {
	this := ControllersRefreshTokenRequest{}
	return &this
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *ControllersRefreshTokenRequest) GetRefreshToken() string {
	if o == nil || o.RefreshToken == nil {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersRefreshTokenRequest) GetRefreshTokenOk() (*string, bool) {
	if o == nil || o.RefreshToken == nil {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *ControllersRefreshTokenRequest) HasRefreshToken() bool {
	if o != nil && o.RefreshToken != nil {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *ControllersRefreshTokenRequest) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

func (o ControllersRefreshTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RefreshToken != nil {
		toSerialize["refreshToken"] = o.RefreshToken
	}
	return json.Marshal(toSerialize)
}

type NullableControllersRefreshTokenRequest struct {
	value *ControllersRefreshTokenRequest
	isSet bool
}

func (v NullableControllersRefreshTokenRequest) Get() *ControllersRefreshTokenRequest {
	return v.value
}

func (v *NullableControllersRefreshTokenRequest) Set(val *ControllersRefreshTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersRefreshTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersRefreshTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersRefreshTokenRequest(val *ControllersRefreshTokenRequest) *NullableControllersRefreshTokenRequest {
	return &NullableControllersRefreshTokenRequest{value: val, isSet: true}
}

func (v NullableControllersRefreshTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersRefreshTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


