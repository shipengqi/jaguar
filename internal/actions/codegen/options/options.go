package options

import (
	"errors"

	"github.com/spf13/pflag"
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

func (o *Options) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Types, "types", o.Types, "Comma-separated list of type names.")
	fs.StringVar(&o.Output, "output", o.Output, "Output filename.")
	fs.StringVar(&o.BuildTags, "build-tags", o.BuildTags, "Comma-separated list of build tags to apply.")
	fs.StringVar(&o.TrimPrefix, "trim-prefix", o.TrimPrefix, "Trim the prefix from the generated constant names.")
	fs.BoolVar(&o.Doc, "doc", o.Doc, "Generate error code documentation in markdown format.")
}

func (o *Options) Validate() []error {
	var errs []error
	if o.Types == "" {
		errs = append(errs, errors.New("--types is required"))
	}
	return errs
}
