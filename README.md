# Go API client for overmind

API for managing your Overmind account

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.5
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import overmind "github.com/overmindtech/api-server/overmind"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), overmind.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), overmind.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), overmind.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), overmind.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://www.df.overmind-demo.com/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AdminApi* | [**AdminCreateAccount**](docs/AdminApi.md#admincreateaccount) | **Post** /admin/accounts | Accounts - Create
*AdminApi* | [**AdminCreateSource**](docs/AdminApi.md#admincreatesource) | **Post** /admin/accounts/{account_name}/sources | Sources - Create
*AdminApi* | [**AdminCreateToken**](docs/AdminApi.md#admincreatetoken) | **Post** /admin/accounts/{account_name}/tokens | Generate a NATS token
*AdminApi* | [**AdminDeleteAccount**](docs/AdminApi.md#admindeleteaccount) | **Delete** /admin/accounts/{account_name} | Accounts - Delete
*AdminApi* | [**AdminDeleteSource**](docs/AdminApi.md#admindeletesource) | **Delete** /admin/accounts/{account_name}/sources/{source_id} | Sources - Delete
*AdminApi* | [**AdminGetAccount**](docs/AdminApi.md#admingetaccount) | **Get** /admin/accounts/{account_name} | Accounts - Get details
*AdminApi* | [**AdminGetSource**](docs/AdminApi.md#admingetsource) | **Get** /admin/accounts/{account_name}/sources/{source_id} | Sources - Get details
*AdminApi* | [**AdminListAccounts**](docs/AdminApi.md#adminlistaccounts) | **Get** /admin/accounts | Accounts - List
*AdminApi* | [**AdminListSources**](docs/AdminApi.md#adminlistsources) | **Get** /admin/accounts/{account_name}/sources | Sources - List
*AdminApi* | [**AdminUpdateSource**](docs/AdminApi.md#adminupdatesource) | **Put** /admin/accounts/{account_name}/sources/{source_id} | Sources - Update
*CoreApi* | [**CreateSource**](docs/CoreApi.md#createsource) | **Post** /core/sources | Sources - Create
*CoreApi* | [**CreateToken**](docs/CoreApi.md#createtoken) | **Post** /core/tokens | Generate a NATS token
*CoreApi* | [**DeleteSource**](docs/CoreApi.md#deletesource) | **Delete** /core/sources/{source_id} | Sources - Delete
*CoreApi* | [**GetAccount**](docs/CoreApi.md#getaccount) | **Get** /core/account | Account - Get details
*CoreApi* | [**GetSource**](docs/CoreApi.md#getsource) | **Get** /core/sources/{source_id} | Sources - Get details
*CoreApi* | [**KeepaliveSources**](docs/CoreApi.md#keepalivesources) | **Post** /core/sources/keepalive | Sources - Keepalive
*CoreApi* | [**ListSources**](docs/CoreApi.md#listsources) | **Get** /core/sources | Sources - List
*CoreApi* | [**UpdateSource**](docs/CoreApi.md#updatesource) | **Put** /core/sources/{source_id} | Sources - Update
*ManagementApi* | [**HealthzGet**](docs/ManagementApi.md#healthzget) | **Get** /healthz | Health check


## Documentation For Models

 - [Account](docs/Account.md)
 - [AdminCreateAccountRequest](docs/AdminCreateAccountRequest.md)
 - [AdminCreateSourceRequest](docs/AdminCreateSourceRequest.md)
 - [Source](docs/Source.md)
 - [TokenRequestData](docs/TokenRequestData.md)


## Documentation For Authorization



### OAuth2


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - **admin:read**: Read information about all accounts and sources from the admin API
 - **admin:write**: Update accounts and sources in the admin API
 - **account:read**: Read information about your account
 - **account:create**: Create new accounts
 - **source:read**: Read source information
 - **source:write**: Create, update and delete sources
 - **request:receive**: Receive requests and respond to them
 - **request:send**: Ability to send requests to subjects like request.scope.{scope}
 - **reverselink:ingest**: Ability to communicate with imported requests and items subjects from other accounts
 - **reverselink:request**: Can make requests to the revlink service
 - **reverselink:respond**: Can respond to reverse linking requests
 - **gateway:stream**: Can stream gateway responses to be picked up by revlink

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


### OAuth2


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://om-dogfood.eu.auth0.com/authorization
- **Scopes**: 
 - **admin:read**: Read information about all accounts and sources from the admin API
 - **admin:write**: Update accounts and sources in the admin API
 - **account:read**: Read information about your account
 - **account:create**: Create new accounts
 - **source:read**: Read source information
 - **source:write**: Create, update and delete sources
 - **request:receive**: Receive requests and respond to them
 - **request:send**: Ability to send requests to subjects like request.scope.{scope}
 - **reverselink:ingest**: Ability to communicate with imported requests and items subjects from other accounts
 - **reverselink:request**: Can make requests to the revlink service
 - **reverselink:respond**: Can respond to reverse linking requests
 - **gateway:stream**: Can stream gateway responses to be picked up by revlink

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



