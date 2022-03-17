# TokenRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserPubKey** | Pointer to **string** | The Public NKey of the user that is requesting a token | [optional] 
**UserName** | Pointer to **string** | Friendly user name | [optional] 

## Methods

### NewTokenRequestData

`func NewTokenRequestData() *TokenRequestData`

NewTokenRequestData instantiates a new TokenRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenRequestDataWithDefaults

`func NewTokenRequestDataWithDefaults() *TokenRequestData`

NewTokenRequestDataWithDefaults instantiates a new TokenRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserPubKey

`func (o *TokenRequestData) GetUserPubKey() string`

GetUserPubKey returns the UserPubKey field if non-nil, zero value otherwise.

### GetUserPubKeyOk

`func (o *TokenRequestData) GetUserPubKeyOk() (*string, bool)`

GetUserPubKeyOk returns a tuple with the UserPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPubKey

`func (o *TokenRequestData) SetUserPubKey(v string)`

SetUserPubKey sets UserPubKey field to given value.

### HasUserPubKey

`func (o *TokenRequestData) HasUserPubKey() bool`

HasUserPubKey returns a boolean if a field has been set.

### GetUserName

`func (o *TokenRequestData) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *TokenRequestData) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *TokenRequestData) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *TokenRequestData) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


