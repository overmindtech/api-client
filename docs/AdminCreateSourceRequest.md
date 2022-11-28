# AdminCreateSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DescriptiveName** | Pointer to **string** | The descriptive name of the source | [optional] 
**IdToken** | Pointer to **string** | OAuth ID token that the source should use | [optional] 
**AccessToken** | Pointer to **string** | OAuth token that allows the source to respond to requests | [optional] 
**RefreshToken** | Pointer to **string** | OAuth refresh token that can be exchanged for a new access_token | [optional] 
**Replicas** | Pointer to **float32** | How many replicas of the source to run??? Do we need this? | [optional] 
**Image** | Pointer to **string** | Docker image of the source | [optional] 
**Config** | Pointer to **map[string]string** | Config for this source. See the source documentation for what config is available/required | [optional] 

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

### GetIdToken

`func (o *AdminCreateSourceRequest) GetIdToken() string`

GetIdToken returns the IdToken field if non-nil, zero value otherwise.

### GetIdTokenOk

`func (o *AdminCreateSourceRequest) GetIdTokenOk() (*string, bool)`

GetIdTokenOk returns a tuple with the IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdToken

`func (o *AdminCreateSourceRequest) SetIdToken(v string)`

SetIdToken sets IdToken field to given value.

### HasIdToken

`func (o *AdminCreateSourceRequest) HasIdToken() bool`

HasIdToken returns a boolean if a field has been set.

### GetAccessToken

`func (o *AdminCreateSourceRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AdminCreateSourceRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AdminCreateSourceRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *AdminCreateSourceRequest) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetRefreshToken

`func (o *AdminCreateSourceRequest) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *AdminCreateSourceRequest) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *AdminCreateSourceRequest) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *AdminCreateSourceRequest) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetReplicas

`func (o *AdminCreateSourceRequest) GetReplicas() float32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *AdminCreateSourceRequest) GetReplicasOk() (*float32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *AdminCreateSourceRequest) SetReplicas(v float32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *AdminCreateSourceRequest) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetImage

`func (o *AdminCreateSourceRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AdminCreateSourceRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AdminCreateSourceRequest) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *AdminCreateSourceRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

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


