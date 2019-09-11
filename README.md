Authlete Library for Go
=======================

Overview
--------

This is the official Go library for [Authlete](https://www.authlete.com/)
Web APIs.

License
-------

  Apache License, Version 2.0

Source Code
-----------

  <code>https://github.com/authlete/authlete-go</code>

Packages
--------

    import (
        "github.com/authlete/authlete-go/api"
        "github.com/authlete/authlete-go/conf"
        "github.com/authlete/authlete-go/dto"
        "github.com/authlete/authlete-go/types"
    )

Quick Start
-----------

The following code simulates "Authorization Code Flow". Replace `CLIENT_ID`,
`SERVICE_API_KEY` and `SERVICE_API_SECRET` in the code with your own properly.
The code assumes that the client type of the client application is 'public'
(otherwise client authentication would be required at the token endpoint) and
the number of registered redirect URIs is one (otherwise `redirect_uri` request
parameter would be required).

```go
package main

import (
    "fmt"
    "github.com/authlete/authlete-go/api"
    "github.com/authlete/authlete-go/conf"
    "github.com/authlete/authlete-go/dto"
)

var (
    base_url = `https://api.authlete.com/`
    service_api_key = `SERVICE_API_KEY`
    service_api_secret = `SERVICE_API_SECRET`
    client_id = `CLIENT_ID`
    user_id = `USER_ID`
)

func main() {
    // Configuration to access Authlete APIs.
    cnf := createAuthleteConfiguration()

    // Authlete API caller.
    api := api.New(cnf)

    //------------------------------------------------------------
    // Authorization Request and Response
    //------------------------------------------------------------

    // Call /api/auth/authorization API. The API parses a given
    // authorization request and returns the parsed result.
    authRes, _ := callAuthorizationApi(api)

    // Call /api/auth/authorization/issue API. The API generates
    // a response for the authorization request.
    authIssueRes, _ := callAuthorizationIssueApi(api, authRes)

    // Authorization response returned to the user agent.
    fmt.Printf("HTTP/1.1 302 Found\n")
    fmt.Printf("Location: %s\n\n", authIssueRes.ResponseContent)

    //------------------------------------------------------------
    // Token Request and Response
    //------------------------------------------------------------

    // Call /api/auth/token API. The API parses a given token
    // request and generates a response for the token request.
    tokenRes, _ := callTokenApi(api, authIssueRes.AuthorizationCode)

    // Token response returned to the client.
    fmt.Printf("HTTP/1.1 200 OK\n")
    fmt.Printf("Content-Type: application/json\n\n")
    fmt.Println(tokenRes.ResponseContent)
}

func createAuthleteConfiguration() conf.AuthleteConfiguration {
    cnf := conf.AuthleteSimpleConfiguration{}
    cnf.SetBaseUrl(base_url)
    cnf.SetServiceApiKey(service_api_key)
    cnf.SetServiceApiSecret(service_api_secret)

    return &cnf
}

func callAuthorizationApi(api api.AuthleteApi) (
    *dto.AuthorizationResponse, *api.AuthleteError) {
    // Prepare a request to /api/auth/authorization API.
    // `Parameters` represents an authorization request.
    req := dto.AuthorizationRequest{}
    req.Parameters = fmt.Sprintf(
        `response_type=code&client_id=%s`, client_id)

    // Call /api/auth/authorization API.
    return api.Authorization(&req)
}

func callAuthorizationIssueApi(
    api api.AuthleteApi, res *dto.AuthorizationResponse) (
        *dto.AuthorizationIssueResponse, *api.AuthleteError) {
    // Prepare a request to /api/auth/authorization/issue API.
    req := dto.AuthorizationIssueRequest{}
    req.Ticket = res.Ticket
    req.Subject = user_id

    // Call /api/auth/authorization/issue API.
    return api.AuthorizationIssue(&req)
}

func callTokenApi(api api.AuthleteApi, code string) (
    *dto.TokenResponse, *api.AuthleteError) {
    // Prepare a request to /api/auth/token API.
    // `Parameters` represents a token request.
    req := dto.TokenRequest{}
    req.Parameters = fmt.Sprintf(
        `client_id=%s&grant_type=authorization_code&code=%s`, client_id, code)

    // Call /api/auth/token API.
    return api.Token(&req)
}
```

Description
-----------

#### How To Get AuthleteApi

All the methods to communicate with Authlete Web APIs are gathered in
`api.AuthleteApi` interface. An instance of the interface can be created by
`api.New(conf.AuthleteConfiguration)` function. The function requires an
instance of `conf.AuthleteConfiguration` interface.

```go
// Prepare an instance of AuthleteConfiguration.
// AuthleteSimpleConfiguration is an implementation of the interface.
cnf := conf.AuthleteSimpleConfiguration{}
cnf.SetBaseUrl(...)
cnf.SetServiceOwnerApiKey(...)
cnf.SetServiceOwnerApiSecret(...)
cnf.SetServiceApiKey(...)
cnf.SetServiceApiSecret(...)

// Get an implementation of AuthleteApi interface.
api := api.New(&cnf)
```

There are three implementations of `conf.AuthleteConfiguration` interface:
`conf.AuthleteSimpleConfiguration` (used in the above example),
`conf.AuthleteEnvConfiguration` and `conf.AuthleteTomlConfiguration`.

`conf.AuthleteEnvConfiguration` reads settings from the following environment
variables.

- `AUTHLETE_BASE_URL`
- `AUTHLETE_SERVICEOWNER_APIKEY`
- `AUTHLETE_SERVICEOWNER_APISECRET`
- `AUTHLETE_SERVICE_APIKEY`
- `AUTHLETE_SERVICE_APISECRET`

Each method of `conf.AuthleteEnvConfiguration` reads its corresponding
environment variable, so what you have to do in Go code is just to create
an instance of the struct.

```go
cnf := conf.AuthleteEnvConfiguration{}
```

On the other hand, `conf.AuthleteTomlConfiguration` reads a TOML file. The
format of the file `conf.AuthleteTomlConfiguration` expects is as follows.

```toml
BaseUrl               = "..."
ServiceOwnerApiKey    = "..."
ServiceOwnerApiSecret = "..."
ServiceApiKey         = "..."
ServiceApiSecret      = "..."
```

`Load(file string)` method loads a TOML file and sets up internal variables.
The method needs to be called before use.

```go
cnf := AuthleteTomlConfiguration{}
cnf.Load(`authlete.toml`)
```

#### AuthleteApi Settings

`Settings()` method of `api.AuthleteApi` interface returns an instance of
`api.Settings`. You can set network timeout via the instance.

```go
import "time"

settings := api.Settings()
settings.Timeout = time.Duration(5) * time.Second
```

#### AuthleteApi Method Categories

Methods in `api.AuthleteApi` interface can be divided into some categories.

##### Methods for Authorization Endpoint Implementation

- `Authorization(request *dto.AuthorizationRequest) (*dto.AuthorizationResponse, *AuthleteError)`
- `AuthorizationFail(request *dto.AuthorizationFailRequest) (*dto.AuthorizationFailResponse, *AuthleteError)`
- `AuthorizationIssue(request *dto.AuthorizationIssueRequest) (*dto.AuthorizationIssueResponse, *AuthleteError)`

##### Methods for Token Endpoint Implementation

- `Token(request *dto.TokenRequest) (*dto.TokenResponse, *AuthleteError)`
- `TokenFail(request *dto.TokenFailRequest) (*dto.TokenFailResponse, *AuthleteError)`
- `TokenIssue(request *dto.TokenIssueRequest) (*dto.TokenIssueResponse, *AuthleteError)`

##### Methods for Service Management

- `CreateService(service *dto.Service) (*dto.Service, *AuthleteError)`
- `DeleteService(apiKey interface{}) *AuthleteError`
- `GetService(apiKey interface{}) (*dto.Service, *AuthleteError)`
- `GetServiceList(start uint32, end uint32) (*dto.ServiceListResponse, *AuthleteError)`
- `UpdateService(service *dto.Service) (*dto.Service, *AuthleteError)`

##### Methods for Client Application Management

- `CreateClient(client *dto.Client) (*dto.Client, *AuthleteError)`
- `DeleteClient(clientIdentifier interface{}) *AuthleteError`
- `GetClient(clientIdentifier interface{}) (*dto.Client, *AuthleteError)`
- `GetClientList(developer string, start uint32, end uint32) (*dto.ClientListResponse, *AuthleteError)`
- `UpdateClient(client *dto.Client) (*dto.Client, *AuthleteError)`
- `RefreshClientSecret(clientIdentifier interface{}) (*dto.ClientSecretRefreshResponse, *AuthleteError)`
- `UpdateClientSecret(clientIdentifier interface{}, clientSecret string) (*dto.ClientSecretUpdateResponse, *AuthleteError)`

##### Methods for Access Token Introspection

- `Introspection(request *dto.IntrospectionRequest) (*dto.IntrospectionResponse, *AuthleteError)`
- `StandardIntrospection(request *dto.StandardIntrospectionRequest) (*dto.StandardIntrospectionResponse, *AuthleteError)`
- `GetTokenList(clientIdentifier string, subject string, start uint32, end uint32) (*dto.TokenListResponse, *AuthleteError)`

##### Methods for Revocation Endpoint Implementation

- `Revocation(request *dto.RevocationRequest) (*dto.RevocationResponse, *AuthleteError)`

##### Methods for User Info Endpoint Implementation

- `UserInfo(request *dto.UserInfoRequest) (*dto.UserInfoResponse, *AuthleteError)`
- `UserInfoIssue(request *dto.UserInfoIssueRequest) (*dto.UserInfoIssueResponse, *AuthleteError)`

##### Methods for JWK Set Endpoint Implementation

- `GetServiceJwks(pretty bool, includePrivateKeys bool) (string, *AuthleteError)`

##### Methods for OpenID Connect Discovery

- `GetServiceConfiguration(pretty bool) (string, *AuthleteError)`

##### Methods for Token Operations

- `TokenCreate(request *dto.TokenCreateRequest) (*dto.TokenCreateResponse, *AuthleteError)`
- `TokenUpdate(request *dto.TokenUpdateRequest) (*dto.TokenUpdateResponse, *AuthleteError)`

##### Methods for Requestable Scopes per Client (deprecated; Client APIs suffice)

- `GetRequestableScopes(clientIdentifier interface{}) ([]string, *AuthleteError)`
- `SetRequestableScopes(clientIdentifier interface{}, scopes []string) ([]string, *AuthleteError)`
- `DeleteRequestableScopes(clientIdentifier interface{}) *AuthleteError`

##### Methods for Records of Granted Scopes

- `GetGrantedScopes(clientIdentifier interface{}, subject string) (*dto.GrantedScopesGetResponse, *AuthleteError)`
- `DeleteGrantedScopes(clientIdentifier interface{}, subject string) *AuthleteError`

##### Methods for Authorization Management on a User-Client Combination Basis

- `DeleteClientAuthorization(clientIdentifier interface{}, subject string) *AuthleteError`
- `GetClientAuthorizationList(request *dto.ClientAuthorizationGetListRequest) (*dto.AuthorizedClientListResponse, *AuthleteError)`
- `UpdateClientAuthorization(clientIdentifier interface{}, request *dto.ClientAuthorizationUpdateRequest) *AuthleteError`

##### Methods for JOSE

- `VerifyJose(request *dto.JoseVerifyRequest) (*dto.JoseVerifyResponse, *AuthleteError)`

##### Methods for CIBA (Client Initiated Backchannel Authentication)

- `BackchannelAuthentication(request *dto.BackchannelAuthenticationRequest) (*dto.BackchannelAuthenticationResponse, *AuthleteError)`
- `BackchannelAuthenticationIssue(request *dto.BackchannelAuthenticationIssueRequest) (*dto.BackchannelAuthenticationIssueResponse, *AuthleteError)`
- `BackchannelAuthenticationFail(request *dto.BackchannelAuthenticationFailRequest) (*dto.BackchannelAuthenticationFailResponse, *AuthleteError)`
- `BackchannelAuthenticationComplete(request *dto.BackchannelAuthenticationCompleteRequest) (*dto.BackchannelAuthenticationCompleteResponse, *AuthleteError)`

##### Methods for OpenID Connect Dynamic Client Registration

- `DynamicClientRegister(request *dto.ClientRegistrationRequest) (*dto.ClientRegistrationResponse, *AuthleteError)`
- `DynamicClientGet(request *dto.ClientRegistrationRequest) (*dto.ClientRegistrationResponse, *AuthleteError)`
- `DynamicClientUpdate(request *dto.ClientRegistrationRequest) (*dto.ClientRegistrationResponse, *AuthleteError)`
- `DynamicClientDelete(request *dto.ClientRegistrationRequest) (*dto.ClientRegistrationResponse, *AuthleteError)`

##### Methods for Device Flow

- `DeviceAuthorization(request *dto.DeviceAuthorizationRequest) (*dto.DeviceAuthorizationResponse, *AuthleteError)`
- `DeviceComplete(request *dto.DeviceCompleteRequest) (*dto.DeviceCompleteResponse, *AuthleteError)`
- `DeviceVerification(request *dto.DeviceVerificationRequest) (*dto.DeviceVerificationResponse, *AuthleteError)`

##### Methods for Pushed Request Object

- `RegisterRequestObject(request *dto.RequestObjectRequest) (*dto.RequestObjectResponse, *AuthleteError)`

Authlete Version
----------------

Some APIs and features don't work (even if they are defined in the `api.AuthleteApi`
interface) if Authlete API server you use doesn't support them. For example,
CIBA works only in Authlete 2.1 onwards. Please contact us if you want to use
newer Authlete versions.

Features available in Authlete 2.0 and onwards:

- Financial-grade API (FAPI)
- OAuth 2.0 Mutual TLS Client Authentication and Certificate Bound Access Tokens (MTLS)
- JWT-based Client Authentication (RFC 7523)
- Scope attributes
- UK Open Banking Security Profile

Features available in Authlete 2.1 and onwards:

- Client Initiated Backchannel Authentication (CIBA)
- JWT Secured Authorization Response Mode for OAuth 2.0 (JARM)
- Dynamic Client Registration (RFC 7591 & RFC 7592)
- OAuth 2.0 Device Authorization Grant (Device Flow)
- JWT-based Access Token

See [Spec Sheet](https://www.authlete.com/legal/spec_sheet/) for further details.

Contact
-------

Contact Form : https://www.authlete.com/contact/

| Purpose   | Email Address        |
|:----------|:---------------------|
| General   | info@authlete.com    |
| Sales     | sales@authlete.com   |
| PR        | pr@authlete.com      |
| Technical | support@authlete.com |
