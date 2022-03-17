# ResponsePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Max** | Pointer to **int64** |  | [optional] 
**Ttl** | Pointer to **int64** | A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately 290 years. | [optional] 

## Methods

### NewResponsePermission

`func NewResponsePermission() *ResponsePermission`

NewResponsePermission instantiates a new ResponsePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsePermissionWithDefaults

`func NewResponsePermissionWithDefaults() *ResponsePermission`

NewResponsePermissionWithDefaults instantiates a new ResponsePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMax

`func (o *ResponsePermission) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ResponsePermission) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ResponsePermission) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *ResponsePermission) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetTtl

`func (o *ResponsePermission) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ResponsePermission) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ResponsePermission) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ResponsePermission) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


