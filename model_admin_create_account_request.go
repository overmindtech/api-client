/*
Overmind API

API for managing your Overmind account

API version: 0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package overmind

import (
	"encoding/json"
)

// checks if the AdminCreateAccountRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminCreateAccountRequest{}

// AdminCreateAccountRequest Creates a new Account. The nkeys is autogenerated by the server
type AdminCreateAccountRequest struct {
	// Name of the account. This should only contain alphanumeric characters and dashes
	Name string `json:"name"`
}

// NewAdminCreateAccountRequest instantiates a new AdminCreateAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminCreateAccountRequest(name string) *AdminCreateAccountRequest {
	this := AdminCreateAccountRequest{}
	this.Name = name
	return &this
}

// NewAdminCreateAccountRequestWithDefaults instantiates a new AdminCreateAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminCreateAccountRequestWithDefaults() *AdminCreateAccountRequest {
	this := AdminCreateAccountRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AdminCreateAccountRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AdminCreateAccountRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AdminCreateAccountRequest) SetName(v string) {
	o.Name = v
}

func (o AdminCreateAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminCreateAccountRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableAdminCreateAccountRequest struct {
	value *AdminCreateAccountRequest
	isSet bool
}

func (v NullableAdminCreateAccountRequest) Get() *AdminCreateAccountRequest {
	return v.value
}

func (v *NullableAdminCreateAccountRequest) Set(val *AdminCreateAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminCreateAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminCreateAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminCreateAccountRequest(val *AdminCreateAccountRequest) *NullableAdminCreateAccountRequest {
	return &NullableAdminCreateAccountRequest{value: val, isSet: true}
}

func (v NullableAdminCreateAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminCreateAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

