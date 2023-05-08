package skeletons

import "embed"

//go:embed api
var API embed.FS

//go:embed cli
var CLI embed.FS

//go:embed grpc
var GRPC embed.FS

//go:embed .gitignore
var GitIgnore embed.FS
