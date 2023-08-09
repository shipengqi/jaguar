package options

import (
	"errors"

	cliflag "github.com/shipengqi/component-base/cli/flag"
	"github.com/shipengqi/component-base/json"
)

type Options struct {
	Types      string
	BuildTags  string
	TrimPrefix string
	Output     string
	Doc        bool
}

func New() *Options {
	return &Options{}
}

func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	s := fss.FlagSet("codegen")
	s.StringVar(&o.Types, "types", o.Types, "Comma-separated list of type names.")
	s.StringVar(&o.Output, "output", o.Output, "Output filename, format: <src dir>/<type>_string.go.")
	s.StringVar(&o.BuildTags, "build-tags", o.BuildTags, "Comma-separated list of build tags to apply.")
	s.StringVar(&o.TrimPrefix, "trim-prefix", o.TrimPrefix, "Trim the `prefix` from the generated constant names.")
	s.BoolVar(&o.Doc, "doc", o.Doc, "Generate error code documentation in markdown format.")
	return
}

// Validate is used to parse and validate the parameters entered by the user at
// the command line when the program starts.
func (o *Options) Validate() []error {
	var errs []error
	if o.Types == "" {
		errs = append(errs, errors.New("--types is required"))
	}
	return errs
}

func (o *Options) String() string {
	data, _ := json.Marshal(o)

	return string(data)
}
