package config

import (
	"encoding/json"

	"{{ .App.ModuleName }}/internal/actions/hello/options"
)

type Config struct {
	*options.Options
}

func (c *Config) String() string {
	data, _ := json.Marshal(c)
	return string(data)
}

func CreateConfigFromOptions(opts *options.Options, _ []string) (*Config, error) {
	return &Config{Options: opts}, nil
}
