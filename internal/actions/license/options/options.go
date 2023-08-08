package options

import (
	cliflag "github.com/shipengqi/component-base/cli/flag"
	"github.com/shipengqi/component-base/json"
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

func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	s := fss.FlagSet("license")
	o.HeaderOptions.AddFlags(s)
	o.SkipOptions.AddFlags(s)

	return
}

// Validate is used to parse and validate the parameters entered by the user at
// the command line when the program starts.
func (o *Options) Validate() []error {
	var errs []error

	errs = append(errs, o.HeaderOptions.Validate()...)
	errs = append(errs, o.SkipOptions.Validate()...)

	return errs
}

func (o *Options) String() string {
	data, _ := json.Marshal(o)

	return string(data)
}
