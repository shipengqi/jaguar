package options

import (
	"encoding/json"

	"{{ .App.ModuleName }}/pkg/app/flag"
)

type Options struct {
	Name string
	Sub  bool
}

func New() *Options {
	return &Options{Name: "World"}
}

func (o *Options) Flags() flag.NamedFlagSets {
	var fss flag.NamedFlagSets
	s := fss.FlagSet("hello")
	s.StringVarP(&o.Name, "name", "n", o.Name, "example name")
	s.BoolVar(&o.Sub, "sub", o.Sub, "sub action example")
	return fss
}

func (o *Options) Validate() []error { return nil }

func (o *Options) Complete() error { return nil }

func (o *Options) String() string {
	data, _ := json.Marshal(o)
	return string(data)
}
