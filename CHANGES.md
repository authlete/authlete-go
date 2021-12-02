CHANGES
=======

- `AuthleteApi` interface
    * Added `HskCreate()` method.
    * Added `HskDelete()` method.
    * Added `HskGet()` method.
    * Added `HskGetList()` method.
    * Added `Echo()` method.

- `AuthorizationResponse` struct
    * Added `DynamicScopes` field.

- `BackchannelAuthenticationCompleteResponse` struct
    * Added `ServiceAttributes` field.
    * Added `ClientAttributes` field.

- `BackchannelAuthenticationResponse` struct
    * Added `DynamicScopes` field.
    * Added `ServiceAttributes` field.
    * Added `ClientAttributes` field.

- `Client` struct
    * Added `Attributes` field.
    * Added `CustomMetadata` field.
    * Added `FrontChannelRequestObjectEncryptionRequired` field.
    * Added `RequestObjectEncryptionAlgMatchRequired` field.
    * Added `RequestObjectEncryptionEncMatchRequired` field.
    * Renamed `AuthorizationDataTypes` to `AuthorizationDetailsTypes`.

- `DeviceAuthorizationResponse` struct
    * Added `DynamicScopes` field.
    * Added `ServiceAttributes` field.
    * Added `ClientAttributes` field.

- `DeviceVerificationResponse` struct
    * Added `DynamicScopes` field.
    * Added `ServiceAttributes` field.
    * Added `ClientAttributes` field.

- `IntrospectionResponse` struct
    * Added `ServiceAttributes` field.
    * Added `ClientAttributes` field.

- `RevocationRequest` struct
    * Added `ClientCertificate` field.
    * Added `ClientCertificatePath` field.

- `Service` struct
    * Added `RefreshTokenDurationReset` field.
    * Added `NbfOptional` field.
    * Added `IssSuppressed` field.
    * Added `Attributes` field.
    * Added `SupportedCustomClientMetadata` field.
    * Added `TokenExpirationLinked` field.
    * Added `FrontChannelRequestObjectEncryptionRequired` field.
    * Added `RequestObjectEncryptionAlgMatchRequired` field.
    * Added `RequestObjectEncryptionEncMatchRequired` field.
    * Added `HsmEnabled` field.
    * Added `Hsks` field.
    * Renamed `SupportedAuthorizationDataTypes` to `SupportedAuthorizationDetailsTypes`.

- `TokenIssueResponse` struct
    * Added `ServiceAttributes` field.
    * Added `ClientAttributes` field.

- `TokenResponse` struct
    * Added `ServiceAttributes` field.
    * Added `ClientAttributes` field.

- `UserInfoResponse` struct
    * Added `ServiceAttributes` field.
    * Added `ClientAttributes` field.

- New components
    * `DynamicScope` struct
    * `Hsk` struct
    * `HskAction` type and some of its constants
    * `HskResponse` struct
    * `HskListAction` type and some of its constants
    * `HskListResponse` struct
    * `HskCreateRequest` struct

v1.1.2 (2020-11-25)
-------------------

- `StandardIntrospectionRequest` struct
    * Added `WithHiddenProperties` field.


v1.1.1 (2020-11-04)
-------------------

- `Service` struct
    * Added `ClaimShortcutRestrictive` field.
    * Added `ScopeRequired` field.


v1.1.0 (2020-10-31)
-------------------

- `AuthleteApi` interface
    * Added `TokenDelete()` method.
    * Added `PushAuthorizationRequest()` method.
    * Removed `RegisterRequestObject()` method.

- `AuthorizationFailReason` type
    * Added `AuthorizationFailReason_INVALID_TARGET`.

- `AuthorizationIssueRequest` struct
    * Added `IdtHeaderParams` field.

- `AuthorizationResponse` struct
    * Added `Resources` field.
    * Added `Purpose` field.

- `BackchannelAuthenticationCompleteRequest` struct
    * Added `IdtHeaderParams` field.

- `BackchannelAuthenticationCompleteResponse` struct
    * Added `Resources` field.

- `BackchannelAuthenticationFailReason` type
    * Added `BackchannelAuthenticationFailReason_INVALID_TARGET`.

- `BackchannelAuthenticationResponse` struct
    * Added `Resources` field.

- `Client` struct
    * Added `Jwks` field.
    * Added `DerivedSectorIdentifier` field.
    * Added `AuthorizationDataTypes` field.
    * Added `ParRequired` field.
    * Added `RequestObjectRequired` field.
    * Removed `SectorIdentifier` field.

- `ClientExtension` struct
    * Added `AccessTokenDuration` field.
    * Added `RefreshTokenDuration` field.

- `DeviceAuthorizationResponse` struct
    * Added `Resources` field.

- `DeviceCompleteRequest` struct
    * Added `IdtHeaderParams` field.

- `DeviceVerificationResponse` struct
    * Added `Resources` field.

- `IntrospectionRequest` struct
    * Added `Dpop` field.
    * Added `Htm` field.
    * Added `Htu` field.

- `IntrospectionResponse` struct
    * Added `Resources` field.
    * Added `AccessTokenResources` field.

- `Service` struct
    * Added `RefreshTokenDurationKept` field.
    * Added `EndSessionEndpoint` field.
    * Added `PushedAuthReqDuration` field.
    * Added `SupportedAuthorizationDataTypes` field.
    * Added `SupportedTrustFrameworks` field.
    * Added `SupportedEvidence` field.
    * Added `SupportedIdentityDocuments` field.
    * Added `SupportedVerificationMethods` field.
    * Added `SupportedVerifiedClaims` field.
    * Added `MissingClientIdAllowed` field.
    * Added `ParRequired` field.
    * Added `RequestObjectRequired` field.
    * Added `TraditionalRequestObjectProcessingApplied` field.
    * Renamed `Scopes` to `SupportedScopes`.
    * Renamed `RequestObjectEndpoint` to `PushedAuthReqEndpoint`.

- `TokenCreateRequest` struct
    * Added `CertificateThumbprint` field.
    * Added `DpopKeyThumbprint` field.

- `TokenFailReason` type
    * Added `TokenFailReason_INVALID_TARGET`.

- `TokenIssueResponse` struct
    * Added `AccessTokenResources` field.

- `TokenRequest` struct
    * Added `Dpop` field.
    * Added `Htm` field.
    * Added `Htu` field.

- `TokenResponse` struct
    * Added `Resources` field.
    * Added `AccessTokenResources` field.

- `TokenUpdateRequest` struct
    * Added `CertificateThumbprint` field.
    * Added `DpopKeyThumbprint` field.

- `UserInfoRequest` struct
    * Added `Dpop` field.
    * Added `Htm` field.
    * Added `Htu` field.

- `UserInfoResponse` struct
    * Added `UserInfoClaims` field.

- New components
    * `PushedAuthReqAction` type and some of its constants
    * `PushedAuthReqRequest` struct
    * `PushedAuthReqResponse` struct

- Deleted components
    * `RequestObjectAction` type
    * `RequestObjectRequest` struct
    * `RequestObjectResponse` struct

v1.0.5 (2019-09-25)
-------------------

- Fixed the implementation of `api.GetServiceJwks(bool, bool)` method
  so that it passes given parameters to /api/service/jwks/get API.

v1.0.4 (2019-09-19)
-------------------

- Changed types of `RefresthTokenExpiresAt` and `RefreshTokenDuration`
  fields in `TokenResponse` from `string` to `uint64`.

v1.0.3 (2019-09-17)
-------------------

- Added `Properties` to `TokenRequest`.

v1.0.2 (2019-09-12)
-------------------

- Added `omitempty` to `struct` fields for JSON marshaling. Without this,
  empty strings are generated instead of `null`.

- Renamed `bakchannel_*.go` files to `backchannel_*.go`.

v1.0.1 (2019-09-11)
-------------------

- Added `authlete.go`. At least one (non-test) Go file must exist directly
  under the package root in order to prevent _"no Go files"_ errors.

v1.0.0 (2019-09-11)
-------------------

- First release
