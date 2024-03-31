package config

import (
	"github.com/shipengqi/component-base/json"

	"github.com/jaguar/apiskeleton/internal/options"
	genericapiserver "github.com/jaguar/apiskeleton/pkg/server"
)

type Config struct {
	*options.Options
}

func (c *Config) String() string {
	data, _ := json.Marshal(c)

	return string(data)
}

func (c *Config) BuildGenericServerConfig() (gsc *genericapiserver.Config, err error) {
	gsc = genericapiserver.NewConfig()

	if err = c.GenericServerRunOptions.ApplyTo(gsc); err != nil {
		return
	}

	if err = c.FeatureOptions.ApplyTo(gsc); err != nil {
		return
	}

	if err = c.SecureServing.ApplyTo(gsc); err != nil {
		return
	}

	if err = c.InsecureServing.ApplyTo(gsc); err != nil {
		return
	}

	return
}

// CreateConfigFromOptions creates a running configuration instance based
// on a given command line.
func CreateConfigFromOptions(opts *options.Options) (*Config, error) {
	cfg := &Config{Options: opts}
	return cfg, nil
}
