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

// checks if the TokenRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenRequestData{}

// TokenRequestData struct for TokenRequestData
type TokenRequestData struct {
	// The Public NKey of the user that is requesting a token
	UserPubKey string `json:"user_pub_key"`
	// Friendly user name
	UserName string `json:"user_name"`
}

// NewTokenRequestData instantiates a new TokenRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenRequestData(userPubKey string, userName string) *TokenRequestData {
	this := TokenRequestData{}
	this.UserPubKey = userPubKey
	this.UserName = userName
	return &this
}

// NewTokenRequestDataWithDefaults instantiates a new TokenRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRequestDataWithDefaults() *TokenRequestData {
	this := TokenRequestData{}
	return &this
}

// GetUserPubKey returns the UserPubKey field value
func (o *TokenRequestData) GetUserPubKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserPubKey
}

// GetUserPubKeyOk returns a tuple with the UserPubKey field value
// and a boolean to check if the value has been set.
func (o *TokenRequestData) GetUserPubKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserPubKey, true
}

// SetUserPubKey sets field value
func (o *TokenRequestData) SetUserPubKey(v string) {
	o.UserPubKey = v
}

// GetUserName returns the UserName field value
func (o *TokenRequestData) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *TokenRequestData) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *TokenRequestData) SetUserName(v string) {
	o.UserName = v
}

func (o TokenRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_pub_key"] = o.UserPubKey
	toSerialize["user_name"] = o.UserName
	return toSerialize, nil
}

type NullableTokenRequestData struct {
	value *TokenRequestData
	isSet bool
}

func (v NullableTokenRequestData) Get() *TokenRequestData {
	return v.value
}

func (v *NullableTokenRequestData) Set(val *TokenRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenRequestData(val *TokenRequestData) *NullableTokenRequestData {
	return &NullableTokenRequestData{value: val, isSet: true}
}

func (v NullableTokenRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


