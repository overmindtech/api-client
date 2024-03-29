/*
Overmind API

API for managing your Overmind account

API version: 0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package overmind

import (
	"encoding/json"
)

// checks if the Account type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Account{}

// Account A NATS Account
type Account struct {
	// The name of the account
	Name string `json:"name"`
	// The public Nkey which signs all NATS \"user\" tokens
	PublicNkey string `json:"public_nkey"`
}

// NewAccount instantiates a new Account object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccount(name string, publicNkey string) *Account {
	this := Account{}
	this.Name = name
	this.PublicNkey = publicNkey
	return &this
}

// NewAccountWithDefaults instantiates a new Account object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountWithDefaults() *Account {
	this := Account{}
	return &this
}

// GetName returns the Name field value
func (o *Account) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Account) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Account) SetName(v string) {
	o.Name = v
}

// GetPublicNkey returns the PublicNkey field value
func (o *Account) GetPublicNkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicNkey
}

// GetPublicNkeyOk returns a tuple with the PublicNkey field value
// and a boolean to check if the value has been set.
func (o *Account) GetPublicNkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicNkey, true
}

// SetPublicNkey sets field value
func (o *Account) SetPublicNkey(v string) {
	o.PublicNkey = v
}

func (o Account) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Account) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["public_nkey"] = o.PublicNkey
	return toSerialize, nil
}

type NullableAccount struct {
	value *Account
	isSet bool
}

func (v NullableAccount) Get() *Account {
	return v.value
}

func (v *NullableAccount) Set(val *Account) {
	v.value = val
	v.isSet = true
}

func (v NullableAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccount(val *Account) *NullableAccount {
	return &NullableAccount{value: val, isSet: true}
}

func (v NullableAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


