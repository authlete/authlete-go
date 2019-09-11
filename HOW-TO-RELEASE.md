HOW TO RELEASE
==============

    $ go mod tidy
    $ go clean ./...
    $ go build ./...
    $ go test ./...
    $ git tag vX.Y.Z
    $ git push origin vX.Y.Z

- Note that tag names must start with `v`.
- `go.sum` should be put into the repository with `go.mod`, too.
- See [Modules](https://github.com/golang/go/wiki/Modules) for details.
