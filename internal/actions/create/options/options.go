package options

import (
	cliflag "github.com/shipengqi/component-base/cli/flag"
	"github.com/shipengqi/component-base/json"
)

type Options struct {
	Force  bool
	Type   string
	Module string
}

func New() *Options {
	o := Options{
		Force: false,
	}

	return &o
}

func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	s := fss.FlagSet("new")
	s.BoolVarP(&o.Force, "force", "f", o.Force, "Force overwriting of existing files.")
	s.StringVar(&o.Type, "type", o.Type, "Project type")
	s.StringVar(&o.Module, "module", o.Module, "Go module name")

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
