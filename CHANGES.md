CHANGES
=======

- `AuthorizationFailReason` type
    * Added `AuthorizationFailReason_INVALID_TARGET`.

- `AuthorizationIssueRequest` struct
    * Added `IdtHeaderParams` field.

- `AuthorizationResponse` struct
    * Added `Resources` field.
    * Added `Purpose` field.

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
