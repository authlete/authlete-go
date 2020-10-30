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
