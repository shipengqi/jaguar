package skeletons

import "embed"

//go:embed v1/api
var API embed.FS

//go:embed v1/cli
var CLI embed.FS

//go:embed v1/grpc
var GRPC embed.FS

//go:embed v1/.gitignore
var GitIgnore embed.FS
