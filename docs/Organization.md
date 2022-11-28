# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **string** | Unique ID of the org | [optional] 
**Name** | Pointer to **string** | Name of the organization you would like to create. This is the name that an end-user would type in the pre-login prompt to identify which organization they wanted to log in through. Unique logical identifier. May contain lowercase alphabetical characters, numbers, underscores (_), and dashes (-). Can start with a number. Must be between 3 and 50 characters. | [optional] 
**DisplayName** | Pointer to **string** | User-friendly name of the organizations that can be displayed in the login flow and email templates. | [optional] 
**PublicNkey** | Pointer to **string** | The public Nkey which signs all NATS \&quot;user\&quot; tokens | [optional] 

## Methods

### NewOrganization

`func NewOrganization() *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *Organization) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Organization) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Organization) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Organization) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetName

`func (o *Organization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Organization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Organization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Organization) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *Organization) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Organization) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Organization) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Organization) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPublicNkey

`func (o *Organization) GetPublicNkey() string`

GetPublicNkey returns the PublicNkey field if non-nil, zero value otherwise.

### GetPublicNkeyOk

`func (o *Organization) GetPublicNkeyOk() (*string, bool)`

GetPublicNkeyOk returns a tuple with the PublicNkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNkey

`func (o *Organization) SetPublicNkey(v string)`

SetPublicNkey sets PublicNkey field to given value.

### HasPublicNkey

`func (o *Organization) HasPublicNkey() bool`

HasPublicNkey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


