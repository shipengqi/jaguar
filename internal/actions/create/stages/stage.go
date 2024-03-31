package stages

import (
	"embed"
	// skeletonsv1 "github.com/shipengqi/jaguar/skeletons/v1"
)

const (
	FilenameGitIgnore = ".gitignore"
	FilenameGoCI      = ".golangci.yaml"
	FilenameReleaser  = ".goreleaser.yaml"
	FilenameSemver    = ".gsemver.yaml"
)

const (
	titleClean     = "Cleaning project"
	titleProject   = "Creating %s project files"
	titleMakeRules = "Creating make rules"
	titleSwagger   = "Configuring swagger"
	titleInitMod   = "Initializing Go module"
	titleModTidy   = "Installing dependencies"
)

var fsmap = make(map[string]embed.FS)

type Interface interface {
	Run() error
}

func Init(_ string) {
	// switch version {
	// case "v1":
	// 	fsmap[FilenameGitIgnore] = skeletonsv1.GitIgnore
	// 	fsmap[FilenameGoCI] = skeletonsv1.GoCI
	// 	fsmap[FilenameReleaser] = skeletonsv1.Releaser
	// 	fsmap[FilenameSemver] = skeletonsv1.Semver
	// 	fsmap[types.ProjectTypeAPI] = skeletonsv1.API
	// 	fsmap[types.ProjectTypeCLI] = skeletonsv1.CLI
	// 	fsmap[types.ProjectTypeGRPC] = skeletonsv1.GRPC
	// default:
	// 	fsmap[FilenameGitIgnore] = skeletonsv1.GitIgnore
	// 	fsmap[FilenameGoCI] = skeletonsv1.GoCI
	// 	fsmap[FilenameReleaser] = skeletonsv1.Releaser
	// 	fsmap[FilenameSemver] = skeletonsv1.Semver
	// 	fsmap[types.ProjectTypeAPI] = skeletonsv1.API
	// 	fsmap[types.ProjectTypeCLI] = skeletonsv1.CLI
	// 	fsmap[types.ProjectTypeGRPC] = skeletonsv1.GRPC
	// }
}
