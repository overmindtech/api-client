/*
NATS Token Exchange

Exchanges OAuth tokens for NATS tokens

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tokenx

import (
	"encoding/json"
)

// ServiceLatency ServiceLatency is used when observing and exported service for latency measurements. Sampling 1-100, represents sampling rate, defaults to 100. Results is the subject where the latency metrics are published. A metric will be defined by the nats-server's ServiceLatency. Time durations are in nanoseconds. see https://github.com/nats-io/nats-server/blob/main/server/accounts.go#L524 e.g. { \"app\": \"dlc22\", \"start\": \"2019-09-16T21:46:23.636869585-07:00\", \"svc\": 219732, \"nats\": { \"req\": 320415, \"resp\": 228268, \"sys\": 0 }, \"total\": 768415 }
type ServiceLatency struct {
	// Subject is a string that represents a NATS subject
	Results *string `json:"results,omitempty"`
	Sampling *int64 `json:"sampling,omitempty"`
}

// NewServiceLatency instantiates a new ServiceLatency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceLatency() *ServiceLatency {
	this := ServiceLatency{}
	return &this
}

// NewServiceLatencyWithDefaults instantiates a new ServiceLatency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceLatencyWithDefaults() *ServiceLatency {
	this := ServiceLatency{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ServiceLatency) GetResults() string {
	if o == nil || o.Results == nil {
		var ret string
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLatency) GetResultsOk() (*string, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ServiceLatency) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given string and assigns it to the Results field.
func (o *ServiceLatency) SetResults(v string) {
	o.Results = &v
}

// GetSampling returns the Sampling field value if set, zero value otherwise.
func (o *ServiceLatency) GetSampling() int64 {
	if o == nil || o.Sampling == nil {
		var ret int64
		return ret
	}
	return *o.Sampling
}

// GetSamplingOk returns a tuple with the Sampling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLatency) GetSamplingOk() (*int64, bool) {
	if o == nil || o.Sampling == nil {
		return nil, false
	}
	return o.Sampling, true
}

// HasSampling returns a boolean if a field has been set.
func (o *ServiceLatency) HasSampling() bool {
	if o != nil && o.Sampling != nil {
		return true
	}

	return false
}

// SetSampling gets a reference to the given int64 and assigns it to the Sampling field.
func (o *ServiceLatency) SetSampling(v int64) {
	o.Sampling = &v
}

func (o ServiceLatency) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	if o.Sampling != nil {
		toSerialize["sampling"] = o.Sampling
	}
	return json.Marshal(toSerialize)
}

type NullableServiceLatency struct {
	value *ServiceLatency
	isSet bool
}

func (v NullableServiceLatency) Get() *ServiceLatency {
	return v.value
}

func (v *NullableServiceLatency) Set(val *ServiceLatency) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceLatency) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceLatency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceLatency(val *ServiceLatency) *NullableServiceLatency {
	return &NullableServiceLatency{value: val, isSet: true}
}

func (v NullableServiceLatency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceLatency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


