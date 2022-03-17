# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultPermissions** | Pointer to [**Permissions**](Permissions.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Exports** | Pointer to [**[]Export**](Export.md) | Exports is a slice of exports | [optional] 
**Imports** | Pointer to [**[]Import**](Import.md) | Imports is a list of import structs | [optional] 
**InfoUrl** | Pointer to **string** |  | [optional] 
**Limits** | Pointer to [**OperatorLimits**](OperatorLimits.md) |  | [optional] 
**Mappings** | Pointer to [**map[string][]WeightedMapping**](array.md) |  | [optional] 
**Revocations** | Pointer to **map[string]int64** | RevocationList is used to store a mapping of public keys to unix timestamps | [optional] 
**SigningKeys** | Pointer to **[]string** | StringList is a wrapper for an array of strings | [optional] 
**Tags** | Pointer to **[]string** | TagList is a unique array of lower case strings All tag list methods lower case the strings in the arguments | [optional] 
**Type** | Pointer to **string** | ClaimType is used to indicate the type of JWT being stored in a Claim | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewAccount

`func NewAccount() *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultPermissions

`func (o *Account) GetDefaultPermissions() Permissions`

GetDefaultPermissions returns the DefaultPermissions field if non-nil, zero value otherwise.

### GetDefaultPermissionsOk

`func (o *Account) GetDefaultPermissionsOk() (*Permissions, bool)`

GetDefaultPermissionsOk returns a tuple with the DefaultPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPermissions

`func (o *Account) SetDefaultPermissions(v Permissions)`

SetDefaultPermissions sets DefaultPermissions field to given value.

### HasDefaultPermissions

`func (o *Account) HasDefaultPermissions() bool`

HasDefaultPermissions returns a boolean if a field has been set.

### GetDescription

`func (o *Account) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Account) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Account) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Account) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExports

`func (o *Account) GetExports() []Export`

GetExports returns the Exports field if non-nil, zero value otherwise.

### GetExportsOk

`func (o *Account) GetExportsOk() (*[]Export, bool)`

GetExportsOk returns a tuple with the Exports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExports

`func (o *Account) SetExports(v []Export)`

SetExports sets Exports field to given value.

### HasExports

`func (o *Account) HasExports() bool`

HasExports returns a boolean if a field has been set.

### GetImports

`func (o *Account) GetImports() []Import`

GetImports returns the Imports field if non-nil, zero value otherwise.

### GetImportsOk

`func (o *Account) GetImportsOk() (*[]Import, bool)`

GetImportsOk returns a tuple with the Imports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImports

`func (o *Account) SetImports(v []Import)`

SetImports sets Imports field to given value.

### HasImports

`func (o *Account) HasImports() bool`

HasImports returns a boolean if a field has been set.

### GetInfoUrl

`func (o *Account) GetInfoUrl() string`

GetInfoUrl returns the InfoUrl field if non-nil, zero value otherwise.

### GetInfoUrlOk

`func (o *Account) GetInfoUrlOk() (*string, bool)`

GetInfoUrlOk returns a tuple with the InfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoUrl

`func (o *Account) SetInfoUrl(v string)`

SetInfoUrl sets InfoUrl field to given value.

### HasInfoUrl

`func (o *Account) HasInfoUrl() bool`

HasInfoUrl returns a boolean if a field has been set.

### GetLimits

`func (o *Account) GetLimits() OperatorLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *Account) GetLimitsOk() (*OperatorLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *Account) SetLimits(v OperatorLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *Account) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetMappings

`func (o *Account) GetMappings() map[string][]WeightedMapping`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *Account) GetMappingsOk() (*map[string][]WeightedMapping, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *Account) SetMappings(v map[string][]WeightedMapping)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *Account) HasMappings() bool`

HasMappings returns a boolean if a field has been set.

### GetRevocations

`func (o *Account) GetRevocations() map[string]int64`

GetRevocations returns the Revocations field if non-nil, zero value otherwise.

### GetRevocationsOk

`func (o *Account) GetRevocationsOk() (*map[string]int64, bool)`

GetRevocationsOk returns a tuple with the Revocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocations

`func (o *Account) SetRevocations(v map[string]int64)`

SetRevocations sets Revocations field to given value.

### HasRevocations

`func (o *Account) HasRevocations() bool`

HasRevocations returns a boolean if a field has been set.

### GetSigningKeys

`func (o *Account) GetSigningKeys() []string`

GetSigningKeys returns the SigningKeys field if non-nil, zero value otherwise.

### GetSigningKeysOk

`func (o *Account) GetSigningKeysOk() (*[]string, bool)`

GetSigningKeysOk returns a tuple with the SigningKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKeys

`func (o *Account) SetSigningKeys(v []string)`

SetSigningKeys sets SigningKeys field to given value.

### HasSigningKeys

`func (o *Account) HasSigningKeys() bool`

HasSigningKeys returns a boolean if a field has been set.

### GetTags

`func (o *Account) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Account) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Account) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Account) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *Account) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Account) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Account) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Account) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *Account) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Account) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Account) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Account) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


