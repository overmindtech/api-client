# ServiceLatency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to **string** | Subject is a string that represents a NATS subject | [optional] 
**Sampling** | Pointer to **int64** |  | [optional] 

## Methods

### NewServiceLatency

`func NewServiceLatency() *ServiceLatency`

NewServiceLatency instantiates a new ServiceLatency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceLatencyWithDefaults

`func NewServiceLatencyWithDefaults() *ServiceLatency`

NewServiceLatencyWithDefaults instantiates a new ServiceLatency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ServiceLatency) GetResults() string`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ServiceLatency) GetResultsOk() (*string, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ServiceLatency) SetResults(v string)`

SetResults sets Results field to given value.

### HasResults

`func (o *ServiceLatency) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetSampling

`func (o *ServiceLatency) GetSampling() int64`

GetSampling returns the Sampling field if non-nil, zero value otherwise.

### GetSamplingOk

`func (o *ServiceLatency) GetSamplingOk() (*int64, bool)`

GetSamplingOk returns a tuple with the Sampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampling

`func (o *ServiceLatency) SetSampling(v int64)`

SetSampling sets Sampling field to given value.

### HasSampling

`func (o *ServiceLatency) HasSampling() bool`

HasSampling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


