変更点
======

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

- 新規部品
    * `PushedAuthReqAction` 型と同型の定数
    * `PushedAuthReqRequest` 構造体
    * `PushedAuthReqResponse` 構造体

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
