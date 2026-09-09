package config

import "{{ .App.ModuleName }}/internal/options"

type Config struct {
	*options.Options
}

// CreateConfigFromOptions creates a running configuration instance based
// on a given command line.
func CreateConfigFromOptions(opts *options.Options) (*Config, error) {
	return &Config{opts}, nil
}
