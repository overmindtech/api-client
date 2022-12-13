# \CoreApi

All URIs are relative to *https://api.overmind.tech*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSource**](CoreApi.md#CreateSource) | **Post** /core/sources | Sources - Create
[**CreateToken**](CoreApi.md#CreateToken) | **Post** /core/tokens | Generate a NATS token
[**DeleteSource**](CoreApi.md#DeleteSource) | **Delete** /core/sources/{source_id} | Sources - Delete
[**GetOrg**](CoreApi.md#GetOrg) | **Get** /core/org | Org - Get details
[**GetSource**](CoreApi.md#GetSource) | **Get** /core/sources/{source_id} | Sources - Get details
[**ListSources**](CoreApi.md#ListSources) | **Get** /core/sources | Sources - List
[**UpdateOrg**](CoreApi.md#UpdateOrg) | **Put** /core/org | Org - Update
[**UpdateSource**](CoreApi.md#UpdateSource) | **Put** /core/sources/{source_id} | Sources - Update



## CreateSource

> Source CreateSource(ctx, orgId).AdminCreateSourceRequest(adminCreateSourceRequest).Execute()

Sources - Create



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgId := "orgId_example" // string | The id of the org
    adminCreateSourceRequest := *openapiclient.NewAdminCreateSourceRequest() // AdminCreateSourceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreApi.CreateSource(context.Background(), orgId).AdminCreateSourceRequest(adminCreateSourceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.CreateSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSource`: Source
    fmt.Fprintf(os.Stdout, "Response from `CoreApi.CreateSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The id of the org | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adminCreateSourceRequest** | [**AdminCreateSourceRequest**](AdminCreateSourceRequest.md) |  | 

### Return type

[**Source**](Source.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateToken

> string CreateToken(ctx).TokenRequestData(tokenRequestData).Execute()

Generate a NATS token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokenRequestData := *openapiclient.NewTokenRequestData() // TokenRequestData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreApi.CreateToken(context.Background()).TokenRequestData(tokenRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.CreateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateToken`: string
    fmt.Fprintf(os.Stdout, "Response from `CoreApi.CreateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRequestData** | [**TokenRequestData**](TokenRequestData.md) |  | 

### Return type

**string**

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/jwt

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSource

> DeleteSource(ctx, sourceId).Execute()

Sources - Delete



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sourceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the source

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreApi.DeleteSource(context.Background(), sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.DeleteSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | ID of the source | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrg

> Organization GetOrg(ctx).Execute()

Org - Get details



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreApi.GetOrg(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.GetOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrg`: Organization
    fmt.Fprintf(os.Stdout, "Response from `CoreApi.GetOrg`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgRequest struct via the builder pattern


### Return type

[**Organization**](Organization.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSource

> Source GetSource(ctx, sourceId).Execute()

Sources - Get details



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sourceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the source

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreApi.GetSource(context.Background(), sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.GetSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSource`: Source
    fmt.Fprintf(os.Stdout, "Response from `CoreApi.GetSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | ID of the source | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Source**](Source.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSources

> []Organization ListSources(ctx).Execute()

Sources - List



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreApi.ListSources(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.ListSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSources`: []Organization
    fmt.Fprintf(os.Stdout, "Response from `CoreApi.ListSources`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSourcesRequest struct via the builder pattern


### Return type

[**[]Organization**](Organization.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrg

> Organization UpdateOrg(ctx).AdminCreateOrgRequest(adminCreateOrgRequest).Execute()

Org - Update



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    adminCreateOrgRequest := *openapiclient.NewAdminCreateOrgRequest() // AdminCreateOrgRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreApi.UpdateOrg(context.Background()).AdminCreateOrgRequest(adminCreateOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.UpdateOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrg`: Organization
    fmt.Fprintf(os.Stdout, "Response from `CoreApi.UpdateOrg`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminCreateOrgRequest** | [**AdminCreateOrgRequest**](AdminCreateOrgRequest.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSource

> UpdateSource(ctx, sourceId).AdminCreateSourceRequest(adminCreateSourceRequest).Execute()

Sources - Update



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sourceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the source
    adminCreateSourceRequest := *openapiclient.NewAdminCreateSourceRequest() // AdminCreateSourceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreApi.UpdateSource(context.Background(), sourceId).AdminCreateSourceRequest(adminCreateSourceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.UpdateSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | ID of the source | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adminCreateSourceRequest** | [**AdminCreateSourceRequest**](AdminCreateSourceRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

