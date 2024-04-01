package options

import (
	cliflag "github.com/shipengqi/component-base/cli/flag"
	"github.com/shipengqi/component-base/json"
	"github.com/spf13/pflag"
)

const (
	FlagForce           = "force"
	FlagProjectType     = "type"
	FlagModuleName      = "module"
	FlagWebFramework    = "web-framework"
	FlagUseGolangCILint = "use-golangci-lint"
	FlagUseGoReleaser   = "use-goreleaser"
	FlagUseGSemver      = "use-gsemver"
	FlagSkeletonVersion = "skeleton-version"
)

type Options struct {
	fs *pflag.FlagSet

	Force             bool
	IsUseGSemver      bool
	IsUseGoReleaser   bool
	IsUseGolangCILint bool
	ProjectType       string
	ProjectName       string
	ModuleName        string
	GoFramework       string
	SkeletonVersion   string
}

func New() *Options {
	o := Options{
		Force:             false,
		IsUseGSemver:      true,
		IsUseGoReleaser:   true,
		IsUseGolangCILint: true,
		SkeletonVersion:   "v1",
	}

	return &o
}

func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.fs = fss.FlagSet("new/create")
	o.fs.BoolVarP(&o.Force, FlagForce, "f", o.Force, "force overwriting of existing files.")
	o.fs.StringVar(&o.ProjectType, FlagProjectType, o.ProjectType, "the type of your application.")
	o.fs.StringVar(&o.ModuleName, FlagModuleName, o.ModuleName, "the Go module name in the go.mod file.")
	o.fs.StringVar(&o.GoFramework, FlagWebFramework, o.GoFramework, "the web framework will be used to build the backend part of your application.")
	o.fs.BoolVar(&o.IsUseGolangCILint, FlagUseGolangCILint, o.IsUseGolangCILint, "use the Golang CI Lint to lint your Go code.")
	o.fs.BoolVar(&o.IsUseGoReleaser, FlagUseGoReleaser, o.IsUseGoReleaser, "use the GoReleaser to deliver your Go binaries.")
	o.fs.BoolVar(&o.IsUseGSemver, FlagUseGSemver, o.IsUseGSemver, "use the GSemver to generate your next semver version.")
	o.fs.StringVar(&o.SkeletonVersion, FlagSkeletonVersion, o.SkeletonVersion, "skeleton version")

	_ = o.fs.MarkHidden("skeleton-version")
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
