package config

import (
	"github.com/shipengqi/component-base/json"
	"github.com/shipengqi/jaguar/internal/actions/create/options"
)

type Config struct {
	*options.Options

	Type string
}

func (o *Config) String() string {
	data, _ := json.Marshal(o)

	return string(data)
}

// CreateConfigFromOptions creates a running configuration instance based
// on a given command line.
func CreateConfigFromOptions(opts *options.Options) (*Config, error) {
	cfg := &Config{Options: opts}

	return cfg, nil
}
