# AdminCreateSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DescriptiveName** | **string** | The descriptive name of the source | 
**Type** | **string** | What source to configure. Currently either \&quot;stdlib\&quot; or \&quot;aws\&quot; | 
**Config** | Pointer to **map[string]interface{}** | Config for this source. See the source documentation for what source-specific config is available/required. This will be supplied directly to viper via a config file at &#x60;/etc/srcman/config/source.yaml&#x60; | [optional] 
**AdditionalConfig** | Pointer to **map[string]string** | Additional config options that should be passed to the source. The keys of this object should be file names, and the values should be their content. These files will be made available to the source at runtime. Check the source&#39;s documentation for what to configure here if required | [optional] 

## Methods

### NewAdminCreateSourceRequest

`func NewAdminCreateSourceRequest(descriptiveName string, type_ string, ) *AdminCreateSourceRequest`

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


### GetConfig

`func (o *AdminCreateSourceRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AdminCreateSourceRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AdminCreateSourceRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AdminCreateSourceRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetAdditionalConfig

`func (o *AdminCreateSourceRequest) GetAdditionalConfig() map[string]string`

GetAdditionalConfig returns the AdditionalConfig field if non-nil, zero value otherwise.

### GetAdditionalConfigOk

`func (o *AdminCreateSourceRequest) GetAdditionalConfigOk() (*map[string]string, bool)`

GetAdditionalConfigOk returns a tuple with the AdditionalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalConfig

`func (o *AdminCreateSourceRequest) SetAdditionalConfig(v map[string]string)`

SetAdditionalConfig sets AdditionalConfig field to given value.

### HasAdditionalConfig

`func (o *AdminCreateSourceRequest) HasAdditionalConfig() bool`

HasAdditionalConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


