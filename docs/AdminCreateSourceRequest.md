# AdminCreateSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DescriptiveName** | Pointer to **string** | The descriptive name of the source | [optional] 
**Type** | Pointer to **string** | What source to configure. Currently either \&quot;stdlib\&quot; or \&quot;aws\&quot; | [optional] 
**Config** | Pointer to **map[string]string** | Config for this source. See the source documentation for what source-specific config is available/required | [optional] 

## Methods

### NewAdminCreateSourceRequest

`func NewAdminCreateSourceRequest() *AdminCreateSourceRequest`

NewAdminCreateSourceRequest instantiates a new AdminCreateSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminCreateSourceRequestWithDefaults

`func NewAdminCreateSourceRequestWithDefaults() *AdminCreateSourceRequest`

NewAdminCreateSourceRequestWithDefaults instantiates a new AdminCreateSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescriptiveName

`func (o *AdminCreateSourceRequest) GetDescriptiveName() string`

GetDescriptiveName returns the DescriptiveName field if non-nil, zero value otherwise.

### GetDescriptiveNameOk

`func (o *AdminCreateSourceRequest) GetDescriptiveNameOk() (*string, bool)`

GetDescriptiveNameOk returns a tuple with the DescriptiveName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptiveName

`func (o *AdminCreateSourceRequest) SetDescriptiveName(v string)`

SetDescriptiveName sets DescriptiveName field to given value.

### HasDescriptiveName

`func (o *AdminCreateSourceRequest) HasDescriptiveName() bool`

HasDescriptiveName returns a boolean if a field has been set.

### GetType

`func (o *AdminCreateSourceRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdminCreateSourceRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdminCreateSourceRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AdminCreateSourceRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfig

`func (o *AdminCreateSourceRequest) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AdminCreateSourceRequest) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AdminCreateSourceRequest) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AdminCreateSourceRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


