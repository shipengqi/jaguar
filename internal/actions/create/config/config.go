package config

import (
	"fmt"

	"github.com/shipengqi/component-base/json"

	"github.com/shipengqi/jaguar/internal/actions/create/helpers"
	"github.com/shipengqi/jaguar/internal/actions/create/options"
	"github.com/shipengqi/jaguar/internal/actions/create/types"
)

type Config struct {
	*options.Options
}

func (c *Config) String() string {
	data, _ := json.Marshal(c)

	return string(data)
}

func (c *Config) ExportTemplateData() *types.TemplateData {
	bin := "apiserver"
	switch projectType := c.ProjectType; projectType {
	case types.ProjectTypeCLI:
		bin = "examplecli"
	case types.ProjectTypeGRPC:
		bin = "rpcserver"
	}
	return &types.TemplateData{
		App: types.AppData{
			Name:           c.ProjectName,
			Type:           c.ProjectType,
			ModuleName:     c.ModuleName,
			Logo:           helpers.NormalizeAppLogo(c.ProjectName),
			EnvPrefix:      helpers.NormalizeAppEnv(c.ProjectName),
			NormalizedName: helpers.NormalizeAppName(c.ProjectName),
			DocumentLink:   fmt.Sprintf("https://%s", c.ModuleName),
		},
		Build: types.BuildData{
			Bin:  bin,
			Root: fmt.Sprintf("%s/cmd/%s", c.ModuleName, bin),
		},
	}
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
