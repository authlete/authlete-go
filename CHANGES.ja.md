変更点
======

v1.1.9 (2022 年 12 月 27 日)
----------------------------

- `Client` 構造体
    * `PkceRequired` フィールドを追加。
    * `PkceS256Required` フィールドを追加。

v1.1.8 (2022 年 09 月 22 日)
----------------------------

- `TokenRequest` 構造体
    * `JwtAtClaims` フィールドを追加。

v1.1.7 (2022 年 08 月 11 日)
----------------------------

- `GrantType` 型
    * `GrantType_JWT_BEARER` 定数を追加。

- `TokenAction` 型
    * `TokenAction_JWT_BEARER` 定数を追加。

- `TokenResponse` 構造体
    * `Assertion` フィールドを追加。

v1.1.6 (2022 年 08 月 09 日)
----------------------------

- `TokenAction` 型
    * `TokenAction_TOKEN_EXCHANGE` 定数を追加。

v1.1.5 (2022 年 08 月 09 日)
----------------------------

- `AuthorizationIssueRequest` 構造体
    * `ConsentedClaims` フィールドを追加。
    * `ClaimsForTx` フィールドを追加。
    * `VerifiedClaimsForTx` フィールドを追加。
    * `JwtAtClaims` フィールドを追加。
    * `AccessToken` フィールドを追加。

- `BackchannelAuthenticationCompleteRequest` 構造体
    * `ConsentedClaims` フィールドを追加。
    * `JwtAtClaims` フィールドを追加。
    * `AccessToken` フィールドを追加。

- `DeviceCompleteRequest` 構造体
    * `ConsentedClaims` フィールドを追加。
    * `JwtAtClaims` フィールドを追加。

- `GrantType` 型
    * `GrantType_TOKEN_EXCHANGE` 定数を追加。

- `TokenCreateRequest` 構造体
    * `Resources` フィールドを追加。
    * `ForExternalAttachment` フィールドを追加。
    * `JwtAtClaims` フィールドを追加。

- `TokenCreateResponse` 構造体
    * `JwtAccessToken` フィールドを追加。
    * `ForExternalAttachment` フィールドを追加。
    * `TokenId` フィールドを追加。

- `TokenIssueRequest` 構造体
    * `JwtAtClaims` フィールドを追加。
    * `AccessToken` フィールドを追加。

- `TokenRequest` 構造体
    * `AccessToken` フィールドを追加。

- `TokenResponse` 構造体
    * `GrantId` フィールドを追加。
    * `Audiences` フィールドを追加。
    * `RequestedTokenType` フィールドを追加。
    * `SubjectToken` フィールドを追加。
    * `SubjectTokenType` フィールドを追加。
    * `SubjectTokenInfo` フィールドを追加。
    * `ActorToken` フィールドを追加。
    * `ActorTokenType` フィールドを追加。
    * `ActorTokenInfo` フィールドを追加。

- `TokenUpdateRequest` 構造体
    * `ForExternalAttachment` フィールドを追加。
    * `TokenId` フィールドを追加。

- `TokenUpdateResponse` 構造体
    * `ForExternalAttachment` フィールドを追加。
    * `TokenId` フィールドを追加。

- 新規部品
    * `TokenInfo` 構造体
    * `TokenType` 型と同型の定数

v1.1.4 (2022 年 03 月 28 日)
----------------------------

- `Service` 構造体
    * `SupportedDocuments` フィールドを追加。
    * `DcrScopeUsedAsRequestable` フィールドを追加。
    * `LoopbackRedirectionUriVariable` フィールドを追加。

- `Settings` 構造体
    * `UserAgent` フィールドを追加。

v1.1.3 (2021 年 12 月 2 日)
----------------------------

- `AuthleteApi` インターフェース
    * `HskCreate()` メソッドを追加
    * `HskDelete()` メソッドを追加
    * `HskGet()` メソッドを追加
    * `HskGetList()` メソッドを追加
    * `Echo()` メソッドを追加

- `AuthorizationResponse` 構造体
    * `DynamicScopes` フィールドを追加。

- `BackchannelAuthenticationCompleteResponse` 構造体
    * `ServiceAttributes` フィールドを追加。
    * `ClientAttributes` フィールドを追加。

- `BackchannelAuthenticationResponse` 構造体
    * `DynamicScopes` フィールドを追加。
    * `ServiceAttributes` フィールドを追加。
    * `ClientAttributes` フィールドを追加。

- `Client` 構造体
    * `Attributes` フィールドを追加。
    * `CustomMetadata` フィールドを追加。
    * `FrontChannelRequestObjectEncryptionRequired` フィールドを追加。
    * `RequestObjectEncryptionAlgMatchRequired` フィールドを追加。
    * `RequestObjectEncryptionEncMatchRequired` フィールドを追加。
    * `AuthorizationDataTypes` を `AuthorizationDetailsTypes` へ名称変更。

- `DeviceAuthorizationResponse` 構造体
    * `DynamicScopes` フィールドを追加。
    * `ServiceAttributes` フィールドを追加。
    * `ClientAttributes` フィールドを追加。

- `DeviceVerificationResponse` 構造体
    * `DynamicScopes` フィールドを追加。
    * `ServiceAttributes` フィールドを追加。
    * `ClientAttributes` フィールドを追加。

- `IntrospectionResponse` 構造体
    * `ServiceAttributes` フィールドを追加。
    * `ClientAttributes` フィールドを追加。

- `RevocationRequest` 構造体
    * `ClientCertificate` フィールドを追加。
    * `ClientCertificatePath` フィールドを追加。

- `Service` 構造体
    * `RefreshTokenDurationReset` フィールドを追加。
    * `NbfOptional` フィールドを追加。
    * `IssSuppressed` フィールドを追加。
    * `Attributes` フィールドを追加。
    * `SupportedCustomClientMetadata` フィールドを追加。
    * `TokenExpirationLinked` フィールドを追加。
    * `FrontChannelRequestObjectEncryptionRequired` フィールドを追加。
    * `RequestObjectEncryptionAlgMatchRequired` フィールドを追加。
    * `RequestObjectEncryptionEncMatchRequired` フィールドを追加。
    * `HsmEnabled` フィールドを追加。
    * `Hsks` フィールドを追加。
    * `SupportedAuthorizationDataTypes` を `SupportedAuthorizationDetailsTypes` へ名称変更。

- `TokenIssueResponse` 構造体
    * `ServiceAttributes` フィールドを追加。
    * `ClientAttributes` フィールドを追加。

- `TokenResponse` 構造体
    * `ServiceAttributes` フィールドを追加。
    * `ClientAttributes` フィールドを追加。

- `UserInfoResponse` 構造体
    * `ServiceAttributes` フィールドを追加。
    * `ClientAttributes` フィールドを追加。

- 新規部品
    * `DynamicScope` 構造体
    * `Hsk` 構造体
    * `HskAction` 型と同型の定数
    * `HskResponse` 構造体
    * `HskListAction` 型と同型の定数
    * `HskListResponse` 構造体
    * `HskCreateRequest` 構造体

v1.1.2 (2020 年 11 月 25 日)
----------------------------

- `StandardIntrospectionRequest` 構造体
    * `WithHiddenProperties` フィールドを追加。


v1.1.1 (2020 年 11 月 04 日)
----------------------------

- `Service` 構造体
    * `ClaimShortcutRestrictive` フィールドを追加。
    * `ScopeRequired` フィールドを追加。


v1.1.0 (2020 年 10 月 31 日)
----------------------------

- `AuthleteApi` インターフェース
    * `TokenDelete()` メソッドを追加。
    * `PushAuthorizationRequest()` メソッドを追加。
    * `RegisterRequestObject()` メソッドを削除。

- `AuthorizationFailReason` 型
    * `AuthorizationFailReason_INVALID_TARGET` を追加。

- `AuthorizationIssueRequest` 構造体
    * `IdtHeaderParams` フィールドを追加。

- `AuthorizationResponse` 構造体
    * `Resources` フィールドを追加。
    * `Purpose` フィールドを追加。

- `BackchannelAuthenticationCompleteRequest` 構造体
    * `IdtHeaderParams` フィールドを追加。

- `BackchannelAuthenticationCompleteResponse` 構造体
    * `Resources` フィールドを追加。

- `BackchannelAuthenticationFailReason` 型
    * `BackchannelAuthenticationFailReason_INVALID_TARGET` を追加。

- `BackchannelAuthenticationResponse` 構造体
    * `Resources` フィールドを追加。

- `Client` 構造体
    * `Jwks` フィールドを追加。
    * `DerivedSectorIdentifier` フィールドを追加。
    * `AuthorizationDataTypes` フィールドを追加。
    * `ParRequired` フィールドを追加。
    * `RequestObjectRequired` フィールドを追加。
    * `SectorIdentifier` フィールドを削除。

- `ClientExtension` 構造体
    * `AccessTokenDuration` フィールドを追加。
    * `RefreshTokenDuration` フィールドを追加。

- `DeviceAuthorizationResponse` 構造体
    * `Resources` フィールドを追加。

- `DeviceCompleteRequest` 構造体
    * `IdtHeaderParams` フィールドを追加。

- `DeviceVerificationResponse` 構造体
    * `Resources` フィールドを追加。

- `IntrospectionRequest` 構造体
    * `Dpop` フィールドを追加。
    * `Htm` フィールドを追加。
    * `Htu` フィールドを追加。

- `IntrospectionResponse` 構造体
    * `Resources` フィールドを追加。
    * `AccessTokenResources` フィールドを追加。

- `Service` 構造体
    * `RefreshTokenDurationKept` フィールドを追加。
    * `EndSessionEndpoint` フィールドを追加。
    * `PushedAuthReqDuration` フィールドを追加。
    * `SupportedAuthorizationDataTypes` フィールドを追加。
    * `SupportedTrustFrameworks` フィールドを追加。
    * `SupportedEvidence` フィールドを追加。
    * `SupportedIdentityDocuments` フィールドを追加。
    * `SupportedVerificationMethods` フィールドを追加。
    * `SupportedVerifiedClaims` フィールドを追加。
    * `MissingClientIdAllowed` フィールドを追加。
    * `ParRequired` フィールドを追加。
    * `RequestObjectRequired` フィールドを追加。
    * `TraditionalRequestObjectProcessingApplied` フィールドを追加。
    * `Scopes` を `SupportedScopes` へ名称変更。
    * `RequestObjectEndpoint` を `PushedAuthReqEndpoint` へ名称変更。

- `TokenCreateRequest` 構造体
    * `CertificateThumbprint` フィールドを追加。
    * `DpopKeyThumbprint` フィールドを追加。

- `TokenFailReason` 型
    * `TokenFailReason_INVALID_TARGET` を追加。

- `TokenIssueResponse` 型
    * `AccessTokenResources` フィールドを追加。

- `TokenRequest` 構造体
    * `Dpop` フィールドを追加。
    * `Htm` フィールドを追加。
    * `Htu` フィールドを追加。

- `TokenResponse` 構造体
    * `Resources` フィールドを追加。
    * `AccessTokenResources` フィールドを追加。

- `TokenUpdateRequest` 構造体
    * `CertificateThumbprint` フィールドを追加。
    * `DpopKeyThumbprint` フィールドを追加。

- `UserInfoRequest` 構造体
    * `Dpop` フィールドを追加。
    * `Htm` フィールドを追加。
    * `Htu` フィールドを追加。

- `UserInfoResponse` 構造体
    * `UserInfoClaims` フィールドを追加。

- 新規部品
    * `PushedAuthReqAction` 型と同型の定数
    * `PushedAuthReqRequest` 構造体
    * `PushedAuthReqResponse` 構造体

- 削除された部品
    * `RequestObjectAction` 型
    * `RequestObjectRequest` 構造体
    * `RequestObjectResponse` 構造体

v1.0.5 (2019 年 09 月 25 日)
----------------------------

- 与えられたパラメーター群を /api/service/jwks/get API に渡すように
  `api.GetServiceJwks(bool, bool)` メソッドの実装を修正。

v1.0.4 (2019 年 09 月 19 日)
----------------------------

- `TokenResponse` のフィールド `RefreshTokenExpiresAt` と `RefreshTokenDuration` の
  型を `string` から `uint64` に変更。

v1.0.3 (2019 年 09 月 17 日)
----------------------------

- `TokenRequest` に `Properties` を追加。

v1.0.2 (2019 年 09 月 12 日)
----------------------------

- JSON マーシャリング用に `omitempty` を `struct` フィールド群に追加。
  これをしておかないと、 `null` のかわりに空文字列が生成されてしまう。

- `bakchannel_*.go` ファイル群の名前を `backchannel_*.go` に変更。

v1.0.1 (2019 年 09 月 11 日)
----------------------------

- `authlete.go` ファイルを追加。 _"no Go files"_ エラーを回避するため、少なくとも
  一つ以上の（テストファイルではない）Go ファイルがパッケージルート直下に必要。

v1.0.0 (2019 年 09 月 11 日)
----------------------------

- 最初のリリース
