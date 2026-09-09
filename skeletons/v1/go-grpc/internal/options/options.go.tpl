package options

import (
	"encoding/json"

	"{{ .App.ModuleName }}/pkg/app/flag"
	"{{ .App.ModuleName }}/pkg/rpcsrv"
	"{{ .App.ModuleName }}/pkg/xlog"
)

type Options struct {
	GRPCOptions *rpcsrv.Options `json:"grpc"    mapstructure:"grpc"`
	Log         *xlog.Options   `json:"log"      mapstructure:"log"`
}

func New() *Options {
	o := Options{
		GRPCOptions: rpcsrv.NewOptions(),
		Log:         xlog.NewOptions(),
	}

	return &o
}

func (o *Options) Flags() (fss flag.NamedFlagSets) {
	o.GRPCOptions.AddFlags(fss.FlagSet("grpc"))
	return
}

// Validate is used to parse and validate the parameters entered by the user at
// the command line when the program starts.
func (o *Options) Validate() []error {
	var errs []error

	errs = append(errs, o.GRPCOptions.Validate()...)
	errs = append(errs, o.Log.Validate()...)

	return errs
}

func (o *Options) String() string {
	data, _ := json.Marshal(o)

	return string(data)
}
