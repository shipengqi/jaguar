package options

import (
	cliflag "github.com/shipengqi/component-base/cli/flag"
	"github.com/shipengqi/component-base/json"
)

type Options struct {
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
	s := fss.FlagSet("new/create")
	s.BoolVarP(&o.Force, "force", "f", o.Force, "force overwriting of existing files.")
	s.StringVar(&o.ProjectType, "type", o.ProjectType, "the type of your application.")
	s.StringVar(&o.ModuleName, "module", o.ModuleName, "the Go module name in the go.mod file.")
	s.StringVar(&o.SkeletonVersion, "skeleton-version", o.SkeletonVersion, "skeleton version")

	_ = s.MarkHidden("skeleton-version")
	return
}

// Validate is used to parse and validate the parameters entered by the user at
// the command line when the program starts.
func (o *Options) Validate() []error {
	return nil
}

func (o *Options) String() string {
	data, _ := json.Marshal(o)

	return string(data)
}
