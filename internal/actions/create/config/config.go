package config

import (
	"github.com/shipengqi/component-base/json"

	"github.com/shipengqi/jaguar/internal/actions/create/options"
)

type Config struct {
	*options.Options

	Type        string
	ProjectName string
}

func (o *Config) String() string {
	data, _ := json.Marshal(o)

	return string(data)
}

// CreateConfigFromOptions creates a running configuration instance based
// on a given command line.
func CreateConfigFromOptions(opts *options.Options, args []string) (*Config, error) {
	cfg := &Config{Options: opts}
	// use the first arg as the project name, args slice doesn't contain flags
	if len(args) > 0 {
		cfg.ProjectName = args[0]
	}
	return cfg, nil
}
