# Source

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DescriptiveName** | **string** | The descriptive name of the source | 
**SourceId** | **string** | Unique ID of the source | 
**TokenName** | **string** | The name of the nats JWT that has been generated for this source | 
**TokenExpiry** | **float32** | When the NATS JWT expires (unix time) | 
**PublicNkey** | **string** | The public NKey associated with the NATS JWT | 
**Type** | **string** | What source to configure. Currently either \&quot;stdlib\&quot; or \&quot;aws\&quot; | 
**Config** | Pointer to **map[string]interface{}** | Config for this source. See the source documentation for what source-specific config is available/required. This will be supplied directly to viper via a config file at &#x60;/etc/srcman/config/source.yaml&#x60; | [optional] 
**AdditionalConfig** | Pointer to **map[string]string** | Additional config options that should be passed to the source. The keys of this object should be file names, and the values should be their content. These files will be made available to the source at runtime. Check the source&#39;s documentation for what to configure here if required | [optional] 

## Methods

### NewSource

`func NewSource(descriptiveName string, sourceId string, tokenName string, tokenExpiry float32, publicNkey string, type_ string, ) *Source`

NewSource instantiates a new Source object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceWithDefaults

`func NewSourceWithDefaults() *Source`

NewSourceWithDefaults instantiates a new Source object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescriptiveName

`func (o *Source) GetDescriptiveName() string`

GetDescriptiveName returns the DescriptiveName field if non-nil, zero value otherwise.

### GetDescriptiveNameOk

`func (o *Source) GetDescriptiveNameOk() (*string, bool)`

GetDescriptiveNameOk returns a tuple with the DescriptiveName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptiveName

`func (o *Source) SetDescriptiveName(v string)`

SetDescriptiveName sets DescriptiveName field to given value.


### GetSourceId

`func (o *Source) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *Source) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *Source) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetTokenName

`func (o *Source) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *Source) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *Source) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.


### GetTokenExpiry

`func (o *Source) GetTokenExpiry() float32`

GetTokenExpiry returns the TokenExpiry field if non-nil, zero value otherwise.

### GetTokenExpiryOk

`func (o *Source) GetTokenExpiryOk() (*float32, bool)`

GetTokenExpiryOk returns a tuple with the TokenExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpiry

`func (o *Source) SetTokenExpiry(v float32)`

SetTokenExpiry sets TokenExpiry field to given value.


### GetPublicNkey

`func (o *Source) GetPublicNkey() string`

GetPublicNkey returns the PublicNkey field if non-nil, zero value otherwise.

### GetPublicNkeyOk

`func (o *Source) GetPublicNkeyOk() (*string, bool)`

GetPublicNkeyOk returns a tuple with the PublicNkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNkey

`func (o *Source) SetPublicNkey(v string)`

SetPublicNkey sets PublicNkey field to given value.


### GetType

`func (o *Source) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Source) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Source) SetType(v string)`

SetType sets Type field to given value.


### GetConfig

`func (o *Source) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Source) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Source) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Source) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetAdditionalConfig

`func (o *Source) GetAdditionalConfig() map[string]string`

GetAdditionalConfig returns the AdditionalConfig field if non-nil, zero value otherwise.

### GetAdditionalConfigOk

`func (o *Source) GetAdditionalConfigOk() (*map[string]string, bool)`

GetAdditionalConfigOk returns a tuple with the AdditionalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalConfig

`func (o *Source) SetAdditionalConfig(v map[string]string)`

SetAdditionalConfig sets AdditionalConfig field to given value.

### HasAdditionalConfig

`func (o *Source) HasAdditionalConfig() bool`

HasAdditionalConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


