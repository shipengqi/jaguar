package options

import (
	"github.com/spf13/pflag"
)

type Options struct {
	HeaderOptions *HeaderOptions `json:"header"   mapstructure:"header"`
	SkipOptions   *SkipOptions   `json:"skip"     mapstructure:"skip"`
}

func New() *Options {
	return &Options{
		HeaderOptions: NewHeaderOptions(),
		SkipOptions:   NewSkipOptions(),
	}
}

func (o *Options) AddFlags(fs *pflag.FlagSet) {
	o.HeaderOptions.AddFlags(fs)
	o.SkipOptions.AddFlags(fs)
}

func (o *Options) Validate() []error {
	var errs []error
	errs = append(errs, o.HeaderOptions.Validate()...)
	errs = append(errs, o.SkipOptions.Validate()...)
	return errs
}
