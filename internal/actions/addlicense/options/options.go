package options

import (
	"fmt"

	cliflag "github.com/shipengqi/component-base/cli/flag"
	"github.com/shipengqi/component-base/json"
)

type Options struct {
	Holder      string
	Year        string
	License     string
	LicenseFile string
	Check       bool
	SkipDirs    []string
	SkipFiles   []string
}

func New() *Options {
	return &Options{
		License: "mit",
	}
}

func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	s := fss.FlagSet("checklicense")
	s.BoolVar(&o.Check, "check", o.Check, "Check only mode: verify if the license header is missing.")
	s.StringVar(&o.Holder, "holder", o.Holder, "Copyright holder")
	s.StringVar(&o.Year, "year", o.Year, "Copyright year(s)")
	s.StringVarP(&o.License, "license", "l", o.License, "Supported license type: apache, bsd, mit, mpl")
	s.StringSliceVar(&o.SkipDirs, "skip-dirs", o.SkipDirs, "Regexps of directories to skip")
	s.StringSliceVar(&o.SkipFiles, "skip-files", o.SkipFiles, "Regexps of files to skip")
	return
}

// Validate is used to parse and validate the parameters entered by the user at
// the command line when the program starts.
func (o *Options) Validate() []error {
	var errs []error

	switch o.License {
	case "apache", "mit", "bsd", "mpl":
		// do nothing
		break
	default:
		errs = append(errs, fmt.Errorf("unsupported license: %s", o.License))
	}
	return errs
}

func (o *Options) String() string {
	data, _ := json.Marshal(o)

	return string(data)
}
