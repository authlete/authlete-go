CHANGES
=======

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
