# OperatorLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conn** | Pointer to **int64** |  | [optional] 
**Consumer** | Pointer to **int64** |  | [optional] 
**Data** | Pointer to **int64** |  | [optional] 
**DiskStorage** | Pointer to **int64** |  | [optional] 
**Exports** | Pointer to **int64** |  | [optional] 
**Imports** | Pointer to **int64** |  | [optional] 
**Leaf** | Pointer to **int64** |  | [optional] 
**MaxBytesRequired** | Pointer to **bool** |  | [optional] 
**MemStorage** | Pointer to **int64** |  | [optional] 
**Payload** | Pointer to **int64** |  | [optional] 
**Streams** | Pointer to **int64** |  | [optional] 
**Subs** | Pointer to **int64** |  | [optional] 
**Wildcards** | Pointer to **bool** |  | [optional] 

## Methods

### NewOperatorLimits

`func NewOperatorLimits() *OperatorLimits`

NewOperatorLimits instantiates a new OperatorLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatorLimitsWithDefaults

`func NewOperatorLimitsWithDefaults() *OperatorLimits`

NewOperatorLimitsWithDefaults instantiates a new OperatorLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConn

`func (o *OperatorLimits) GetConn() int64`

GetConn returns the Conn field if non-nil, zero value otherwise.

### GetConnOk

`func (o *OperatorLimits) GetConnOk() (*int64, bool)`

GetConnOk returns a tuple with the Conn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConn

`func (o *OperatorLimits) SetConn(v int64)`

SetConn sets Conn field to given value.

### HasConn

`func (o *OperatorLimits) HasConn() bool`

HasConn returns a boolean if a field has been set.

### GetConsumer

`func (o *OperatorLimits) GetConsumer() int64`

GetConsumer returns the Consumer field if non-nil, zero value otherwise.

### GetConsumerOk

`func (o *OperatorLimits) GetConsumerOk() (*int64, bool)`

GetConsumerOk returns a tuple with the Consumer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumer

`func (o *OperatorLimits) SetConsumer(v int64)`

SetConsumer sets Consumer field to given value.

### HasConsumer

`func (o *OperatorLimits) HasConsumer() bool`

HasConsumer returns a boolean if a field has been set.

### GetData

`func (o *OperatorLimits) GetData() int64`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OperatorLimits) GetDataOk() (*int64, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OperatorLimits) SetData(v int64)`

SetData sets Data field to given value.

### HasData

`func (o *OperatorLimits) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDiskStorage

`func (o *OperatorLimits) GetDiskStorage() int64`

GetDiskStorage returns the DiskStorage field if non-nil, zero value otherwise.

### GetDiskStorageOk

`func (o *OperatorLimits) GetDiskStorageOk() (*int64, bool)`

GetDiskStorageOk returns a tuple with the DiskStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStorage

`func (o *OperatorLimits) SetDiskStorage(v int64)`

SetDiskStorage sets DiskStorage field to given value.

### HasDiskStorage

`func (o *OperatorLimits) HasDiskStorage() bool`

HasDiskStorage returns a boolean if a field has been set.

### GetExports

`func (o *OperatorLimits) GetExports() int64`

GetExports returns the Exports field if non-nil, zero value otherwise.

### GetExportsOk

`func (o *OperatorLimits) GetExportsOk() (*int64, bool)`

GetExportsOk returns a tuple with the Exports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExports

`func (o *OperatorLimits) SetExports(v int64)`

SetExports sets Exports field to given value.

### HasExports

`func (o *OperatorLimits) HasExports() bool`

HasExports returns a boolean if a field has been set.

### GetImports

`func (o *OperatorLimits) GetImports() int64`

GetImports returns the Imports field if non-nil, zero value otherwise.

### GetImportsOk

`func (o *OperatorLimits) GetImportsOk() (*int64, bool)`

GetImportsOk returns a tuple with the Imports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImports

`func (o *OperatorLimits) SetImports(v int64)`

SetImports sets Imports field to given value.

### HasImports

`func (o *OperatorLimits) HasImports() bool`

HasImports returns a boolean if a field has been set.

### GetLeaf

`func (o *OperatorLimits) GetLeaf() int64`

GetLeaf returns the Leaf field if non-nil, zero value otherwise.

### GetLeafOk

`func (o *OperatorLimits) GetLeafOk() (*int64, bool)`

GetLeafOk returns a tuple with the Leaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaf

`func (o *OperatorLimits) SetLeaf(v int64)`

SetLeaf sets Leaf field to given value.

### HasLeaf

`func (o *OperatorLimits) HasLeaf() bool`

HasLeaf returns a boolean if a field has been set.

### GetMaxBytesRequired

`func (o *OperatorLimits) GetMaxBytesRequired() bool`

GetMaxBytesRequired returns the MaxBytesRequired field if non-nil, zero value otherwise.

### GetMaxBytesRequiredOk

`func (o *OperatorLimits) GetMaxBytesRequiredOk() (*bool, bool)`

GetMaxBytesRequiredOk returns a tuple with the MaxBytesRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBytesRequired

`func (o *OperatorLimits) SetMaxBytesRequired(v bool)`

SetMaxBytesRequired sets MaxBytesRequired field to given value.

### HasMaxBytesRequired

`func (o *OperatorLimits) HasMaxBytesRequired() bool`

HasMaxBytesRequired returns a boolean if a field has been set.

### GetMemStorage

`func (o *OperatorLimits) GetMemStorage() int64`

GetMemStorage returns the MemStorage field if non-nil, zero value otherwise.

### GetMemStorageOk

`func (o *OperatorLimits) GetMemStorageOk() (*int64, bool)`

GetMemStorageOk returns a tuple with the MemStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemStorage

`func (o *OperatorLimits) SetMemStorage(v int64)`

SetMemStorage sets MemStorage field to given value.

### HasMemStorage

`func (o *OperatorLimits) HasMemStorage() bool`

HasMemStorage returns a boolean if a field has been set.

### GetPayload

`func (o *OperatorLimits) GetPayload() int64`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *OperatorLimits) GetPayloadOk() (*int64, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *OperatorLimits) SetPayload(v int64)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *OperatorLimits) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetStreams

`func (o *OperatorLimits) GetStreams() int64`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *OperatorLimits) GetStreamsOk() (*int64, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *OperatorLimits) SetStreams(v int64)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *OperatorLimits) HasStreams() bool`

HasStreams returns a boolean if a field has been set.

### GetSubs

`func (o *OperatorLimits) GetSubs() int64`

GetSubs returns the Subs field if non-nil, zero value otherwise.

### GetSubsOk

`func (o *OperatorLimits) GetSubsOk() (*int64, bool)`

GetSubsOk returns a tuple with the Subs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubs

`func (o *OperatorLimits) SetSubs(v int64)`

SetSubs sets Subs field to given value.

### HasSubs

`func (o *OperatorLimits) HasSubs() bool`

HasSubs returns a boolean if a field has been set.

### GetWildcards

`func (o *OperatorLimits) GetWildcards() bool`

GetWildcards returns the Wildcards field if non-nil, zero value otherwise.

### GetWildcardsOk

`func (o *OperatorLimits) GetWildcardsOk() (*bool, bool)`

GetWildcardsOk returns a tuple with the Wildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcards

`func (o *OperatorLimits) SetWildcards(v bool)`

SetWildcards sets Wildcards field to given value.

### HasWildcards

`func (o *OperatorLimits) HasWildcards() bool`

HasWildcards returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


