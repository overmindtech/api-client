# \AdminApi

All URIs are relative to *https://api.overmind.tech*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminCreateOrg**](AdminApi.md#AdminCreateOrg) | **Post** /admin/orgs | Orgs - Create
[**AdminCreateSource**](AdminApi.md#AdminCreateSource) | **Post** /admin/orgs/{org_id}/sources | Sources - Create
[**AdminCreateToken**](AdminApi.md#AdminCreateToken) | **Post** /admin/orgs/{org_id}/tokens | Generate a NATS token
[**AdminDeleteOrg**](AdminApi.md#AdminDeleteOrg) | **Delete** /admin/orgs/{org_id} | Orgs - Delete
[**AdminDeleteSource**](AdminApi.md#AdminDeleteSource) | **Delete** /admin/orgs/{org_id}/sources/{source_id} | Sources - Delete
[**AdminGetOrg**](AdminApi.md#AdminGetOrg) | **Get** /admin/orgs/{org_id} | Orgs - Get details
[**AdminGetSource**](AdminApi.md#AdminGetSource) | **Get** /admin/orgs/{org_id}/sources/{source_id} | Sources - Get details
[**AdminListOrgs**](AdminApi.md#AdminListOrgs) | **Get** /admin/orgs | Orgs - List
[**AdminListSources**](AdminApi.md#AdminListSources) | **Get** /admin/orgs/{org_id}/sources | Sources - List
[**AdminUpdateOrg**](AdminApi.md#AdminUpdateOrg) | **Put** /admin/orgs/{org_id} | Orgs - Update
[**AdminUpdateSource**](AdminApi.md#AdminUpdateSource) | **Put** /admin/orgs/{org_id}/sources/{source_id} | Sources - Update



## AdminCreateOrg

> Organization AdminCreateOrg(ctx).AdminCreateOrgRequest(adminCreateOrgRequest).Execute()

Orgs - Create



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
    resp, r, err := apiClient.AdminApi.AdminCreateOrg(context.Background()).AdminCreateOrgRequest(adminCreateOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminCreateOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminCreateOrg`: Organization
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.AdminCreateOrg`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminCreateOrgRequest struct via the builder pattern


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


## AdminCreateSource

> Source AdminCreateSource(ctx, orgId).AdminCreateSourceRequest(adminCreateSourceRequest).Execute()

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
    resp, r, err := apiClient.AdminApi.AdminCreateSource(context.Background(), orgId).AdminCreateSourceRequest(adminCreateSourceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminCreateSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminCreateSource`: Source
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.AdminCreateSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The id of the org | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminCreateSourceRequest struct via the builder pattern


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


## AdminCreateToken

> string AdminCreateToken(ctx, orgId).TokenRequestData(tokenRequestData).Execute()

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
    orgId := "orgId_example" // string | The id of the org
    tokenRequestData := *openapiclient.NewTokenRequestData() // TokenRequestData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.AdminCreateToken(context.Background(), orgId).TokenRequestData(tokenRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminCreateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminCreateToken`: string
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.AdminCreateToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The id of the org | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminCreateTokenRequest struct via the builder pattern


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


## AdminDeleteOrg

> AdminDeleteOrg(ctx, orgId).Execute()

Orgs - Delete



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.AdminDeleteOrg(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminDeleteOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The id of the org | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminDeleteOrgRequest struct via the builder pattern


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


## AdminDeleteSource

> AdminDeleteSource(ctx, orgId, sourceId).Execute()

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
    orgId := "orgId_example" // string | The id of the org
    sourceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the source

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.AdminDeleteSource(context.Background(), orgId, sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminDeleteSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The id of the org | 
**sourceId** | **string** | ID of the source | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminDeleteSourceRequest struct via the builder pattern


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


## AdminGetOrg

> Organization AdminGetOrg(ctx, orgId).Execute()

Orgs - Get details



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.AdminGetOrg(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminGetOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminGetOrg`: Organization
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.AdminGetOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The id of the org | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminGetOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## AdminGetSource

> Source AdminGetSource(ctx, orgId, sourceId).Execute()

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
    orgId := "orgId_example" // string | The id of the org
    sourceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the source

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.AdminGetSource(context.Background(), orgId, sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminGetSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminGetSource`: Source
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.AdminGetSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The id of the org | 
**sourceId** | **string** | ID of the source | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminGetSourceRequest struct via the builder pattern


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


## AdminListOrgs

> []Organization AdminListOrgs(ctx).Execute()

Orgs - List



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
    resp, r, err := apiClient.AdminApi.AdminListOrgs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminListOrgs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminListOrgs`: []Organization
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.AdminListOrgs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminListOrgsRequest struct via the builder pattern


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


## AdminListSources

> []Source AdminListSources(ctx, orgId).Execute()

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
    orgId := "orgId_example" // string | The id of the org

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.AdminListSources(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminListSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminListSources`: []Source
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.AdminListSources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The id of the org | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminListSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Source**](Source.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminUpdateOrg

> Organization AdminUpdateOrg(ctx, orgId).AdminCreateOrgRequest(adminCreateOrgRequest).Execute()

Orgs - Update



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
    adminCreateOrgRequest := *openapiclient.NewAdminCreateOrgRequest() // AdminCreateOrgRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.AdminUpdateOrg(context.Background(), orgId).AdminCreateOrgRequest(adminCreateOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminUpdateOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminUpdateOrg`: Organization
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.AdminUpdateOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The id of the org | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminUpdateOrgRequest struct via the builder pattern


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


## AdminUpdateSource

> AdminUpdateSource(ctx, orgId, sourceId).AdminCreateSourceRequest(adminCreateSourceRequest).Execute()

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
    orgId := "orgId_example" // string | The id of the org
    sourceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the source
    adminCreateSourceRequest := *openapiclient.NewAdminCreateSourceRequest() // AdminCreateSourceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.AdminUpdateSource(context.Background(), orgId, sourceId).AdminCreateSourceRequest(adminCreateSourceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminUpdateSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The id of the org | 
**sourceId** | **string** | ID of the source | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminUpdateSourceRequest struct via the builder pattern


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

