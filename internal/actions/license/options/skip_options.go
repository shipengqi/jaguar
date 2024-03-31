package options

import (
	"github.com/spf13/pflag"
)

type SkipOptions struct {
	SkipDirs  []string
	SkipFiles []string
}

func NewSkipOptions() *SkipOptions {
	return &SkipOptions{}
}

func (o *SkipOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringSliceVar(&o.SkipDirs, "skip-dirs", o.SkipDirs, "Regexps of directories to skip")
	fs.StringSliceVar(&o.SkipFiles, "skip-files", o.SkipFiles, "Regexps of files to skip")
}

// Validate is used to parse and validate the parameters entered by the user at
// the command line when the program starts.
func (o *SkipOptions) Validate() []error {
	return nil
}
