package v1

import "embed"

//go:embed api
var API embed.FS

//go:embed cli
var CLI embed.FS

//go:embed grpc
var GRPC embed.FS

//go:embed .gitignore
var GitIgnore embed.FS

//go:embed .golangci.yaml
var GoCI embed.FS

//go:embed .goreleaser.yaml
var Releaser embed.FS

//go:embed .gsemver.yaml
var Semver embed.FS
