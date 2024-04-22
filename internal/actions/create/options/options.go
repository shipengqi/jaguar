package options

import (
	cliflag "github.com/shipengqi/component-base/cli/flag"
	"github.com/shipengqi/component-base/json"
	"github.com/spf13/pflag"
)

const (
	FlagProjectType      = "type"
	FlagModuleName       = "module"
	FlagProjectName      = "project"
	FlagWebFramework     = "web-framework"
	FlagUseGolangCILint  = "use-golangci-lint"
	FlagUseGoReleaser    = "use-goreleaser"
	FlagUseGSemver       = "use-gsemver"
	FlagUseGithubActions = "use-github-actions"
	FlagSkeletonVersion  = "skeleton-version"
)

const (
	SkeletonVersion1 = "v1"
)

type Options struct {
	fs *pflag.FlagSet

	IsUseGSemver       bool
	IsUseGoReleaser    bool
	IsUseGolangCILint  bool
	IsUseGithubActions bool
	ProjectType        string
	ProjectName        string
	ModuleName         string
	GoFramework        string
	SkeletonVersion    string
}

func New() *Options {
	o := Options{
		IsUseGSemver:       true,
		IsUseGoReleaser:    true,
		IsUseGolangCILint:  true,
		IsUseGithubActions: true,
		SkeletonVersion:    SkeletonVersion1,
		GoFramework:        "gin",
	}

	return &o
}

func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.fs = fss.FlagSet("New/Create")
	o.fs.StringVarP(&o.ProjectType, FlagProjectType, "t", o.ProjectType, "the type of your application.")
	o.fs.StringVarP(&o.ProjectName, FlagProjectName, "n", o.ProjectName, "the name of your application.")
	o.fs.StringVarP(&o.ModuleName, FlagModuleName, "m", o.ModuleName, "the Go module name in the go.mod file.")
	o.fs.StringVar(&o.GoFramework, FlagWebFramework, o.GoFramework, "the web framework will be used to build the backend part of your application.\nMust be one of: gin, fiber.")
	o.fs.BoolVar(&o.IsUseGolangCILint, FlagUseGolangCILint, o.IsUseGolangCILint, "use the Golang CI Lint to lint your Go code.")
	o.fs.BoolVar(&o.IsUseGoReleaser, FlagUseGoReleaser, o.IsUseGoReleaser, "use the GoReleaser to deliver your Go binaries.")
	o.fs.BoolVar(&o.IsUseGSemver, FlagUseGSemver, o.IsUseGSemver, "use the GSemver to generate your next semver version.")
	o.fs.BoolVar(&o.IsUseGithubActions, FlagUseGithubActions, o.IsUseGithubActions, "add common Github actions.")
	o.fs.StringVar(&o.SkeletonVersion, FlagSkeletonVersion, o.SkeletonVersion, "skeleton version")

	// Todo The fiber is still under development and needs to be hidden.
	_ = o.fs.MarkHidden(FlagWebFramework)
	_ = o.fs.MarkHidden(FlagSkeletonVersion)
	return
}

func (o *Options) Changed(name string) bool {
	return o.fs.Changed(name)
}

// Validate is used to parse and validate the parameters entered by the user at
// the command line when the program starts.
func (o *Options) Validate() []error {
	// Todo add config validators here
	return nil
}

func (o *Options) String() string {
	data, _ := json.Marshal(o)

	return string(data)
}
