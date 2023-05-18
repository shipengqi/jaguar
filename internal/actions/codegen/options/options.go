package options

import (
	cliflag "github.com/shipengqi/component-base/cli/flag"
	"github.com/shipengqi/component-base/json"
)

type Options struct{}

func New() *Options {
	return &Options{}
}

func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
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
