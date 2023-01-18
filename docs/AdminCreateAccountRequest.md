# AdminCreateAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the account. This should only contain alphanumeric characters and dashes | [optional] 

## Methods

### NewAdminCreateAccountRequest

`func NewAdminCreateAccountRequest() *AdminCreateAccountRequest`

NewAdminCreateAccountRequest instantiates a new AdminCreateAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminCreateAccountRequestWithDefaults

`func NewAdminCreateAccountRequestWithDefaults() *AdminCreateAccountRequest`

NewAdminCreateAccountRequestWithDefaults instantiates a new AdminCreateAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AdminCreateAccountRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminCreateAccountRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminCreateAccountRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdminCreateAccountRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


