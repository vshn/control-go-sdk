# Contributing

Meaningful pull requests to this repository are welcome! However, it's alway a good idea to open an issue first to discuss your plan with us first.

## Code guidelines

Ensure the following points are okay before submitting the PR:

- Your code adheres to `go fmt`
  - As a bonus, also passes `golangci-lint run ./...`
- All exported funcs, types, vars and constants are documented
- You ran `go mod tidy`
- All tests pass
- The changelog is updated (under the "Unreleased" section)

## Getting started

This projects uses `go mod`, so simply fork & clone this project and start hacking away!

## Tests

    go test ./...

## Releasing

- update changelog
- update version in `version.go`
- Draft new release https://github.com/vshn/control-go-sdk/releases/new
- Tag as `vX.X.X`
- changelog entries
