package config

import (
	"strings"

	"github.com/shipengqi/component-base/json"

	"github.com/shipengqi/jaguar/internal/actions/codegen/options"
)

type Config struct {
	*options.Options

	BuildTagSlice []string
}

func (o *Config) String() string {
	data, _ := json.Marshal(o)

	return string(data)
}

// CreateConfigFromOptions creates a running configuration instance based
// on a given command line.
func CreateConfigFromOptions(opts *options.Options) (*Config, error) {
	cfg := &Config{}
	if opts.BuildTags != "" {
		tags := strings.Split(opts.BuildTags, ",")
		for _, v := range tags {
			cfg.BuildTagSlice = append(cfg.BuildTagSlice, strings.TrimSpace(v))
		}
	}

	return cfg, nil
}

