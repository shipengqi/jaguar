package options

import (
	"fmt"

	"github.com/spf13/pflag"
)

type HeaderOptions struct {
	Holder      string
	Year        string
	License     string
	LicenseFile string
}

func NewHeaderOptions() *HeaderOptions {
	return &HeaderOptions{
		License: "apache",
	}
}

func (o *HeaderOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Holder, "holder", o.Holder, "Copyright holder")
	fs.StringVar(&o.Year, "year", o.Year, "Copyright year(s)")
	fs.StringVarP(&o.License, "license", "l", o.License, "License type: apache, bsd, mit, mpl")
	fs.StringVarP(&o.LicenseFile, "license-file", "f", o.LicenseFile, "License file")
}

// Validate is used to parse and validate the parameters entered by the user at
// the command line when the program starts.
func (o *HeaderOptions) Validate() []error {
	var errs []error

	switch o.License {
	case "apache", "mit", "bsd", "mpl":
		// do nothing
		break
	default:
		errs = append(errs, fmt.Errorf("unsupported license type: %s", o.License))
	}
	return errs
}
