# Limits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **int64** |  | [optional] 
**Payload** | Pointer to **int64** |  | [optional] 
**Src** | Pointer to **[]string** | TagList is a unique array of lower case strings All tag list methods lower case the strings in the arguments | [optional] 
**Subs** | Pointer to **int64** |  | [optional] 
**Times** | Pointer to [**[]TimeRange**](TimeRange.md) |  | [optional] 
**TimesLocation** | Pointer to **string** |  | [optional] 

## Methods

### NewLimits

`func NewLimits() *Limits`

NewLimits instantiates a new Limits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLimitsWithDefaults

`func NewLimitsWithDefaults() *Limits`

NewLimitsWithDefaults instantiates a new Limits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Limits) GetData() int64`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Limits) GetDataOk() (*int64, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Limits) SetData(v int64)`

SetData sets Data field to given value.

### HasData

`func (o *Limits) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPayload

`func (o *Limits) GetPayload() int64`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Limits) GetPayloadOk() (*int64, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Limits) SetPayload(v int64)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *Limits) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetSrc

`func (o *Limits) GetSrc() []string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *Limits) GetSrcOk() (*[]string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *Limits) SetSrc(v []string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *Limits) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### GetSubs

`func (o *Limits) GetSubs() int64`

GetSubs returns the Subs field if non-nil, zero value otherwise.

### GetSubsOk

`func (o *Limits) GetSubsOk() (*int64, bool)`

GetSubsOk returns a tuple with the Subs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubs

`func (o *Limits) SetSubs(v int64)`

SetSubs sets Subs field to given value.

### HasSubs

`func (o *Limits) HasSubs() bool`

HasSubs returns a boolean if a field has been set.

### GetTimes

`func (o *Limits) GetTimes() []TimeRange`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *Limits) GetTimesOk() (*[]TimeRange, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *Limits) SetTimes(v []TimeRange)`

SetTimes sets Times field to given value.

### HasTimes

`func (o *Limits) HasTimes() bool`

HasTimes returns a boolean if a field has been set.

### GetTimesLocation

`func (o *Limits) GetTimesLocation() string`

GetTimesLocation returns the TimesLocation field if non-nil, zero value otherwise.

### GetTimesLocationOk

`func (o *Limits) GetTimesLocationOk() (*string, bool)`

GetTimesLocationOk returns a tuple with the TimesLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesLocation

`func (o *Limits) SetTimesLocation(v string)`

SetTimesLocation sets TimesLocation field to given value.

### HasTimesLocation

`func (o *Limits) HasTimesLocation() bool`

HasTimesLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


