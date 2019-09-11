Go 用 Authlete ライブラリ
=========================

概要
----

[Authlete](https://www.authlete.com) Web API 用の公式 Go ライブラリです。

ライセンス
----------

  Apache License, Version 2.0

ソースコード
------------

  <code>https://github.com/authlete/authlete-go</code>

パッケージ
----------

    import (
        "github.com/authlete/authlete-go/api"
        "github.com/authlete/authlete-go/conf"
        "github.com/authlete/authlete-go/dto"
        "github.com/authlete/authlete-go/types"
    )

クイックスタート
----------------

下記のコードは「認可コードフロー」をシミュレートするものです。コード内の `CLIENT_ID`,
`SERVICE_API_KEY`, `SERVICE_API_SECRET` を適宜あなたのもので置き換えてください。
コードでは、クライアントアプリケーションのクライアントタイプが public であること
（そうでない場合はトークンエンドポイントでクライアント認証が要求される）、
登録されているリダイレクト URI の数が一つであること（そうでない場合は `redirect_uri`
リクエストパラメーターが要求される）を想定しています。

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
    // Authlete API にアクセスするための設定
    cnf := createAuthleteConfiguration()

    // Authlete API を呼び出すためのインスタンス
    api := api.New(cnf)

    //------------------------------------------------------------
    // 認可リクエストとレスポンス
    //------------------------------------------------------------

    // /api/auth/authorization API を呼ぶ。この API は、与えられた
    // 認可リクエストをパースし、その結果を返す。
    authRes, _ := callAuthorizationApi(api)

    // /api/auth/authorization/issue API を呼ぶ。この API は、認可
    // リクエストに対するレスポンスを生成する。
    authIssueRes, _ := callAuthorizationIssueApi(api, authRes)

    // ユーザーエージェントに返す認可レスポンス。
    fmt.Printf("HTTP/1.1 302 Found\n")
    fmt.Printf("Location: %s\n\n", authIssueRes.ResponseContent)

    //------------------------------------------------------------
    // トークンリクエストとレスポンス
    //------------------------------------------------------------

    // /api/auth/token API を呼ぶ。この API は、与えられたトークン
    // リクエストをパースし、トークンレスポンスを生成する。
    tokenRes, _ := callTokenApi(api, authIssueRes.AuthorizationCode)

    // クライアントに返すトークンレスポンス。
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
    // /api/auth/authorization API へのリクエストを用意する。
    // `Parameters` は認可リクエストを表す。
    req := dto.AuthorizationRequest{}
    req.Parameters = fmt.Sprintf(
        `response_type=code&client_id=%s`, client_id)

    // /api/auth/authorization API を呼ぶ。
    return api.Authorization(&req)
}

func callAuthorizationIssueApi(
    api api.AuthleteApi, res *dto.AuthorizationResponse) (
        *dto.AuthorizationIssueResponse, *api.AuthleteError) {
    // /api/auth/authorization/issue API へのリクエストを用意する。
    req := dto.AuthorizationIssueRequest{}
    req.Ticket = res.Ticket
    req.Subject = user_id

    // /api/auth/authorization/issue API を呼ぶ。
    return api.AuthorizationIssue(&req)
}

func callTokenApi(api api.AuthleteApi, code string) (
    *dto.TokenResponse, *api.AuthleteError) {
    // /api/auth/token API へのリクエストを用意する。
    // `Parameters` はトークンリクエストを表す。
    req := dto.TokenRequest{}
    req.Parameters = fmt.Sprintf(
        `client_id=%s&grant_type=authorization_code&code=%s`, client_id, code)

    // /api/auth/token API を呼ぶ。
    return api.Token(&req)
}
```

説明
----

#### AuthleteApi の取得方法

Authlete Web API とやりとりするメソッドは全て `api.AuthleteApi`
インターフェースに集められています。 このインターフェースのインスタンスは
`api.New(conf.AuthleteConfiguration)` 関数で生成することができます。
この関数は、引数として `conf.AuthleteConfiguration`
インターフェースのインスタンスを要求します。

```go
// AuthleteConfiguration のインスタンスを用意する。
// AuthleteSimpleConfiguration はインターフェース実装の一つ。
cnf := conf.AuthleteSimpleConfiguration{}
cnf.SetBaseUrl(...)
cnf.SetServiceOwnerApiKey(...)
cnf.SetServiceOwnerApiSecret(...)
cnf.SetServiceApiKey(...)
cnf.SetServiceApiSecret(...)

// AuthleteApi インターフェースの実装を取得する。
api := api.New(&cnf)
```

`conf.AuthleteConfiguration` には三つの実装があります。
`conf.AuthleteSimpleConfiguration` （上記の例で使用）、
`conf.AuthleteEnvConfiguration` および `conf.AuthleteTomlConfiguration` です。

`conf.AuthleteEnvConfiguration` は次の環境変数から設定を読み込みます。

- `AUTHLETE_BASE_URL`
- `AUTHLETE_SERVICEOWNER_APIKEY`
- `AUTHLETE_SERVICEOWNER_APISECRET`
- `AUTHLETE_SERVICE_APIKEY`
- `AUTHLETE_SERVICE_APISECRET`

`conf.AuthleteEnvConfiguration` の各メソッドがそれぞれに対応する環境変数を読み込むので、Go
コードに書かなければならないのは、次のようにインスタンスを生成する処理だけです。

```go
cnf := conf.AuthleteEnvConfiguration{}
```

一方、 `conf.AuthleteTomlConfiguration` は TOML ファイルから設定を読み込みます。
`conf.AuthleteTomlConfiguration` が想定しているファイルフォーマットは次の通りです。

```toml
BaseUrl               = "..."
ServiceOwnerApiKey    = "..."
ServiceOwnerApiSecret = "..."
ServiceApiKey         = "..."
ServiceApiSecret      = "..."
```

`Load(file string)` メソッドは TOML ファイルを読み、内部変数をセットアップします。
使用前にこのメソッドを呼ぶ必要があります。

```go
cnf := AuthleteTomlConfiguration{}
cnf.Load(`authlete.toml`)
```

#### AuthleteApi の設定

`api.AuthleteApi` インターフェースの `Settings()` メソッドは `api.Settings`
のインスタンスを返します。このインスタンスを介して、ネットワークのタイムアウトを設定することができます。

```go
import "time"

settings := api.Settings()
settings.Timeout = time.Duration(5) * time.Second
```

#### AuthleteApi メソッドのカテゴリー

`AuthleteApi` インターフェースのメソッド群はいくつかのカテゴリーに分けることができます。

##### 認可エンドポイント実装のためのメソッド群

- `Authorization(request *dto.AuthorizationRequest) (*dto.AuthorizationResponse, *AuthleteError)`
- `AuthorizationFail(request *dto.AuthorizationFailRequest) (*dto.AuthorizationFailResponse, *AuthleteError)`
- `AuthorizationIssue(request *dto.AuthorizationIssueRequest) (*dto.AuthorizationIssueResponse, *AuthleteError)`

##### トークンエンドポイント実装のためのメソッド群

- `Token(request *dto.TokenRequest) (*dto.TokenResponse, *AuthleteError)`
- `TokenFail(request *dto.TokenFailRequest) (*dto.TokenFailResponse, *AuthleteError)`
- `TokenIssue(request *dto.TokenIssueRequest) (*dto.TokenIssueResponse, *AuthleteError)`

##### サービス管理のためのメソッド群

- `CreateService(service *dto.Service) (*dto.Service, *AuthleteError)`
- `DeleteService(apiKey interface{}) *AuthleteError`
- `GetService(apiKey interface{}) (*dto.Service, *AuthleteError)`
- `GetServiceList(start uint32, end uint32) (*dto.ServiceListResponse, *AuthleteError)`
- `UpdateService(service *dto.Service) (*dto.Service, *AuthleteError)`

##### クライアントアプリケーション管理のためのメソッド群

- `CreateClient(client *dto.Client) (*dto.Client, *AuthleteError)`
- `DeleteClient(clientIdentifier interface{}) *AuthleteError`
- `GetClient(clientIdentifier interface{}) (*dto.Client, *AuthleteError)`
- `GetClientList(developer string, start uint32, end uint32) (*dto.ClientListResponse, *AuthleteError)`
- `UpdateClient(client *dto.Client) (*dto.Client, *AuthleteError)`
- `RefreshClientSecret(clientIdentifier interface{}) (*dto.ClientSecretRefreshResponse, *AuthleteError)`
- `UpdateClientSecret(clientIdentifier interface{}, clientSecret string) (*dto.ClientSecretUpdateResponse, *AuthleteError)`

##### アクセストークンの情報取得のためのメソッド群

- `Introspection(request *dto.IntrospectionRequest) (*dto.IntrospectionResponse, *AuthleteError)`
- `StandardIntrospection(request *dto.StandardIntrospectionRequest) (*dto.StandardIntrospectionResponse, *AuthleteError)`
- `GetTokenList(clientIdentifier string, subject string, start uint32, end uint32) (*dto.TokenListResponse, *AuthleteError)`

##### アクセストークン取り消しエンドポイント実装のためのメソッド群

- `Revocation(request *dto.RevocationRequest) (*dto.RevocationResponse, *AuthleteError)`

##### ユーザー情報エンドポイント実装のためのメソッド群

- `UserInfo(request *dto.UserInfoRequest) (*dto.UserInfoResponse, *AuthleteError)`
- `UserInfoIssue(request *dto.UserInfoIssueRequest) (*dto.UserInfoIssueResponse, *AuthleteError)`

##### JWK セットエンドポイント実装のためのメソッド群

- `GetServiceJwks(pretty bool, includePrivateKeys bool) (string, *AuthleteError)`

##### OpenID Connect Discovery のためのメソッド群

- `GetServiceConfiguration(pretty bool) (string, *AuthleteError)`

##### トークン操作のためのメソッド群

- `TokenCreate(request *dto.TokenCreateRequest) (*dto.TokenCreateResponse, *AuthleteError)`
- `TokenUpdate(request *dto.TokenUpdateRequest) (*dto.TokenUpdateResponse, *AuthleteError)`

##### クライアント毎の要求可能スコープ群に関するメソッド群 (非推奨; Client API で代替可能)

- `GetRequestableScopes(clientIdentifier interface{}) ([]string, *AuthleteError)`
- `SetRequestableScopes(clientIdentifier interface{}, scopes []string) ([]string, *AuthleteError)`
- `DeleteRequestableScopes(clientIdentifier interface{}) *AuthleteError`

##### 付与されたスコープの記録に関するメソッド群

- `GetGrantedScopes(clientIdentifier interface{}, subject string) (*dto.GrantedScopesGetResponse, *AuthleteError)`
- `DeleteGrantedScopes(clientIdentifier interface{}, subject string) *AuthleteError`

##### ユーザー・クライアント単位の認可管理に関するメソッド群

- `DeleteClientAuthorization(clientIdentifier interface{}, subject string) *AuthleteError`
- `GetClientAuthorizationList(request *dto.ClientAuthorizationGetListRequest) (*dto.AuthorizedClientListResponse, *AuthleteError)`
- `UpdateClientAuthorization(clientIdentifier interface{}, request *dto.ClientAuthorizationUpdateRequest) *AuthleteError`

##### JOSE に関するメソッド群

- `VerifyJose(request *dto.JoseVerifyRequest) (*dto.JoseVerifyResponse, *AuthleteError)`

##### CIBA (Client Initiated Backchannel Authentication) に関するメソッド群

- `BackchannelAuthentication(request *dto.BackchannelAuthenticationRequest) (*dto.BackchannelAuthenticationResponse, *AuthleteError)`
- `BackchannelAuthenticationIssue(request *dto.BackchannelAuthenticationIssueRequest) (*dto.BackchannelAuthenticationIssueResponse, *AuthleteError)`
- `BackchannelAuthenticationFail(request *dto.BackchannelAuthenticationFailRequest) (*dto.BackchannelAuthenticationFailResponse, *AuthleteError)`
- `BackchannelAuthenticationComplete(request *dto.BackchannelAuthenticationCompleteRequest) (*dto.BackchannelAuthenticationCompleteResponse, *AuthleteError)`

##### OpenID Connect Dynamic Client Registration に関するメソッド群

- `DynamicClientRegister(request *dto.ClientRegistrationRequest) (*dto.ClientRegistrationResponse, *AuthleteError)`
- `DynamicClientGet(request *dto.ClientRegistrationRequest) (*dto.ClientRegistrationResponse, *AuthleteError)`
- `DynamicClientUpdate(request *dto.ClientRegistrationRequest) (*dto.ClientRegistrationResponse, *AuthleteError)`
- `DynamicClientDelete(request *dto.ClientRegistrationRequest) (*dto.ClientRegistrationResponse, *AuthleteError)`

##### Device Flow に関するメソッド群

- `DeviceAuthorization(request *dto.DeviceAuthorizationRequest) (*dto.DeviceAuthorizationResponse, *AuthleteError)`
- `DeviceComplete(request *dto.DeviceCompleteRequest) (*dto.DeviceCompleteResponse, *AuthleteError)`
- `DeviceVerification(request *dto.DeviceVerificationRequest) (*dto.DeviceVerificationResponse, *AuthleteError)`

##### Pushed Request Object に関するメソッド群

- `RegisterRequestObject(request *dto.RequestObjectRequest) (*dto.RequestObjectResponse, *AuthleteError)`

Authlete バージョン
-------------------

幾つかの API や機能は、使用されている Authlete API サーバーがサポートしていなければ（たとえ
`api.AuthleteApi` インターフェースで定義されているとしても）使うことができません。例えば、CIBA は
Authlete 2.1 以降でのみ機能します。新しい Authlete バージョンを使用されたい場合は、ご連絡ください。

Authlete 2.0 以降で利用できる機能：

- Financial-grade API (FAPI)
- OAuth 2.0 Mutual TLS Client Authentication and Certificate Bound Access Tokens (MTLS)
- JWT-based Client Authentication (RFC 7523)
- Scope attributes
- UK Open Banking Security Profile

Authlete 2.1 以降で利用できる機能：

- Client Initiated Backchannel Authentication (CIBA)
- JWT Secured Authorization Response Mode for OAuth 2.0 (JARM)
- Dynamic Client Registration (RFC 7591 & RFC 7592)
- OAuth 2.0 Device Authorization Grant (Device Flow)
- JWT-based Access Token

詳細情報は [スペックシート](https://www.authlete.com/ja/legal/spec_sheet/)
を参照してください。

コンタクト
----------

コンタクトフォーム : https://www.authlete.com/contact/

| 目的 | メールアドレス       |
|:-----|:---------------------|
| 一般 | info@authlete.com    |
| 営業 | sales@authlete.com   |
| 広報 | pr@authlete.com      |
| 技術 | support@authlete.com |
