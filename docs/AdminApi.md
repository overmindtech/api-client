# \AdminApi

All URIs are relative to *https://www.df.overmind-demo.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminCreateAccount**](AdminApi.md#AdminCreateAccount) | **Post** /admin/accounts | Accounts - Create
[**AdminCreateSource**](AdminApi.md#AdminCreateSource) | **Post** /admin/accounts/{account_name}/sources | Sources - Create
[**AdminCreateToken**](AdminApi.md#AdminCreateToken) | **Post** /admin/accounts/{account_name}/tokens | Generate a NATS token
[**AdminDeleteAccount**](AdminApi.md#AdminDeleteAccount) | **Delete** /admin/accounts/{account_name} | Accounts - Delete
[**AdminDeleteSource**](AdminApi.md#AdminDeleteSource) | **Delete** /admin/accounts/{account_name}/sources/{source_id} | Sources - Delete
[**AdminGetAccount**](AdminApi.md#AdminGetAccount) | **Get** /admin/accounts/{account_name} | Accounts - Get details
[**AdminGetSource**](AdminApi.md#AdminGetSource) | **Get** /admin/accounts/{account_name}/sources/{source_id} | Sources - Get details
[**AdminKeepaliveSources**](AdminApi.md#AdminKeepaliveSources) | **Post** /admin/accounts/{account_name}/sources/keepalive | Sources - Keepalive
[**AdminListAccounts**](AdminApi.md#AdminListAccounts) | **Get** /admin/accounts | Accounts - List
[**AdminListSources**](AdminApi.md#AdminListSources) | **Get** /admin/accounts/{account_name}/sources | Sources - List
[**AdminUpdateSource**](AdminApi.md#AdminUpdateSource) | **Put** /admin/accounts/{account_name}/sources/{source_id} | Sources - Update



## AdminCreateAccount

> Account AdminCreateAccount(ctx).AdminCreateAccountRequest(adminCreateAccountRequest).Execute()

Accounts - Create



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/overmindtech/api-server/overmind"
)

func main() {
    adminCreateAccountRequest := *openapiclient.NewAdminCreateAccountRequest("test-account") // AdminCreateAccountRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.AdminCreateAccount(context.Background()).AdminCreateAccountRequest(adminCreateAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminCreateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminCreateAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.AdminCreateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminCreateAccountRequest** | [**AdminCreateAccountRequest**](AdminCreateAccountRequest.md) |  | 

### Return type

[**Account**](Account.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminCreateSource

> Source AdminCreateSource(ctx, accountName).AdminCreateSourceRequest(adminCreateSourceRequest).Execute()

Sources - Create



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/overmindtech/api-server/overmind"
)

func main() {
    accountName := "accountName_example" // string | The name of the account
    adminCreateSourceRequest := *openapiclient.NewAdminCreateSourceRequest("stdlib", "aws") // AdminCreateSourceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.AdminCreateSource(context.Background(), accountName).AdminCreateSourceRequest(adminCreateSourceRequest).Execute()
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
**accountName** | **string** | The name of the account | 

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

> string AdminCreateToken(ctx, accountName).TokenRequestData(tokenRequestData).Execute()

Generate a NATS token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/overmindtech/api-server/overmind"
)

func main() {
    accountName := "accountName_example" // string | The name of the account
    tokenRequestData := *openapiclient.NewTokenRequestData("UserPubKey_example", "UserName_example") // TokenRequestData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.AdminCreateToken(context.Background(), accountName).TokenRequestData(tokenRequestData).Execute()
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
**accountName** | **string** | The name of the account | 

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


## AdminDeleteAccount

> AdminDeleteAccount(ctx, accountName).Execute()

Accounts - Delete



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/overmindtech/api-server/overmind"
)

func main() {
    accountName := "accountName_example" // string | The name of the account

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminApi.AdminDeleteAccount(context.Background(), accountName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminDeleteAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | The name of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminDeleteAccountRequest struct via the builder pattern


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

> AdminDeleteSource(ctx, accountName, sourceId).Execute()

Sources - Delete



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/overmindtech/api-server/overmind"
)

func main() {
    accountName := "accountName_example" // string | The name of the account
    sourceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the source

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminApi.AdminDeleteSource(context.Background(), accountName, sourceId).Execute()
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
**accountName** | **string** | The name of the account | 
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


## AdminGetAccount

> Account AdminGetAccount(ctx, accountName).Execute()

Accounts - Get details



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/overmindtech/api-server/overmind"
)

func main() {
    accountName := "accountName_example" // string | The name of the account

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.AdminGetAccount(context.Background(), accountName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminGetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminGetAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.AdminGetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | The name of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Account**](Account.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminGetSource

> Source AdminGetSource(ctx, accountName, sourceId).Execute()

Sources - Get details



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/overmindtech/api-server/overmind"
)

func main() {
    accountName := "accountName_example" // string | The name of the account
    sourceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the source

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.AdminGetSource(context.Background(), accountName, sourceId).Execute()
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
**accountName** | **string** | The name of the account | 
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


## AdminKeepaliveSources

> AdminKeepaliveSources(ctx, accountName).Execute()

Sources - Keepalive



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/overmindtech/api-server/overmind"
)

func main() {
    accountName := "accountName_example" // string | The name of the account

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminApi.AdminKeepaliveSources(context.Background(), accountName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminKeepaliveSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | The name of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminKeepaliveSourcesRequest struct via the builder pattern


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


## AdminListAccounts

> []Account AdminListAccounts(ctx).Execute()

Accounts - List



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/overmindtech/api-server/overmind"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.AdminListAccounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminListAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminListAccounts`: []Account
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.AdminListAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminListAccountsRequest struct via the builder pattern


### Return type

[**[]Account**](Account.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminListSources

> []Source AdminListSources(ctx, accountName).Execute()

Sources - List



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/overmindtech/api-server/overmind"
)

func main() {
    accountName := "accountName_example" // string | The name of the account

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.AdminListSources(context.Background(), accountName).Execute()
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
**accountName** | **string** | The name of the account | 

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


## AdminUpdateSource

> AdminUpdateSource(ctx, accountName, sourceId).AdminCreateSourceRequest(adminCreateSourceRequest).Execute()

Sources - Update



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/overmindtech/api-server/overmind"
)

func main() {
    accountName := "accountName_example" // string | The name of the account
    sourceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the source
    adminCreateSourceRequest := *openapiclient.NewAdminCreateSourceRequest("stdlib", "aws") // AdminCreateSourceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminApi.AdminUpdateSource(context.Background(), accountName, sourceId).AdminCreateSourceRequest(adminCreateSourceRequest).Execute()
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
**accountName** | **string** | The name of the account | 
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

