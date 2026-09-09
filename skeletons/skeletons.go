package skeletons

import "embed"

var (
	//go:embed all:v1/go-api
	V1GoAPI embed.FS

	//go:embed all:v1/go-embed
	V1GoEmbed embed.FS

	//go:embed all:v1/go-cli
	V1GoCLI embed.FS

	//go:embed all:v1/go-grpc
	V1GoGRPC embed.FS

	//go:embed all:v1/nodejs
	V1NodeJS embed.FS

	//go:embed all:v1/python
	V1Python embed.FS

	//go:embed all:v1/projectfiles
	V1ProjectFiles embed.FS
)

type Skeleton struct {
	GoAPI, GoEmbed, GoCLI, GoGRPC, NodeJS, Python, ProjectFiles embed.FS
}

type Skeletons struct {
	V1 Skeleton
}

func New() *Skeletons {
	return &Skeletons{
		V1: Skeleton{
			GoAPI:        V1GoAPI,
			GoEmbed:      V1GoEmbed,
			GoCLI:        V1GoCLI,
			GoGRPC:       V1GoGRPC,
			NodeJS:       V1NodeJS,
			Python:       V1Python,
			ProjectFiles: V1ProjectFiles,
		},
	}
}
