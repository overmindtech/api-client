# \ManagementApi

All URIs are relative to *https://www.df.overmind-demo.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HealthzGet**](ManagementApi.md#HealthzGet) | **Get** /healthz | Health check



## HealthzGet

> HealthzGet(ctx).Execute()

Health check



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
    r, err := apiClient.ManagementApi.HealthzGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.HealthzGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthzGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

