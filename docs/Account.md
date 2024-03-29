# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the account | 
**PublicNkey** | **string** | The public Nkey which signs all NATS \&quot;user\&quot; tokens | 

## Methods

### NewAccount

`func NewAccount(name string, publicNkey string, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Account) SetName(v string)`

SetName sets Name field to given value.


### GetPublicNkey

`func (o *Account) GetPublicNkey() string`

GetPublicNkey returns the PublicNkey field if non-nil, zero value otherwise.

### GetPublicNkeyOk

`func (o *Account) GetPublicNkeyOk() (*string, bool)`

GetPublicNkeyOk returns a tuple with the PublicNkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNkey

`func (o *Account) SetPublicNkey(v string)`

SetPublicNkey sets PublicNkey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


