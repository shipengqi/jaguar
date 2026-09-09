package options

import (
	"encoding/json"

	"{{ .App.ModuleName }}/pkg/app/flag"
	"{{ .App.ModuleName }}/pkg/cache"
	"{{ .App.ModuleName }}/pkg/db"
	genericoptions "{{ .App.ModuleName }}/pkg/options"
	"{{ .App.ModuleName }}/pkg/xlog"
)

type Options struct {
	GenericServerRunOptions *genericoptions.ServerRunOptions       `json:"server"   mapstructure:"server"`
	SecureServing           *genericoptions.SecureServingOptions   `json:"secure"   mapstructure:"secure"`
	InsecureServing         *genericoptions.InsecureServingOptions `json:"insecure" mapstructure:"insecure"`
	FeatureOptions          *genericoptions.FeatureOptions         `json:"feature"  mapstructure:"feature"`
	MySQLOptions            *db.Options                            `json:"mysql"    mapstructure:"mysql"`
	RedisOptions            *cache.Options                         `json:"redis"    mapstructure:"redis"`
	Log                     *xlog.Options                          `json:"log"      mapstructure:"log"`
	Store                   string                                 `json:"store"    mapstructure:"store"`
}

func New() *Options {
	return &Options{
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		InsecureServing:         genericoptions.NewInsecureServingOptions(),
		SecureServing:           genericoptions.NewSecureServingOptions(),
		FeatureOptions:          genericoptions.NewFeatureOptions(),
		MySQLOptions:            db.NewOptions(),
		RedisOptions:            cache.NewOptions(),
		Log:                     xlog.NewOptions(),
		Store:                   "sqlite",
	}
}

func (o *Options) Flags() flag.NamedFlagSets {
	var fss flag.NamedFlagSets
	o.GenericServerRunOptions.AddFlags(fss.FlagSet("generic"))
	o.InsecureServing.AddFlags(fss.FlagSet("insecure serving"))
	o.SecureServing.AddFlags(fss.FlagSet("secure serving"))
	o.FeatureOptions.AddFlags(fss.FlagSet("features"))
	o.MySQLOptions.AddFlags(fss.FlagSet("mysql"))
	o.RedisOptions.AddFlags(fss.FlagSet("redis"))
	o.Log.AddFlags(fss.FlagSet("logs"))
	fss.FlagSet("storage").StringVar(&o.Store, "store", o.Store, "Storage backend to use: mysql or sqlite.")
	return fss
}

func (o *Options) Validate() []error {
	var errs []error
	errs = append(errs, o.InsecureServing.Validate()...)
	errs = append(errs, o.SecureServing.Validate()...)
	errs = append(errs, o.MySQLOptions.Validate()...)
	errs = append(errs, o.Log.Validate()...)
	errs = append(errs, o.FeatureOptions.Validate()...)
	return errs
}

func (o *Options) Complete() error { return nil }

func (o *Options) String() string {
	data, _ := json.Marshal(o)
	return string(data)
}
