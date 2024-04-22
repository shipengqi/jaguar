package config

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/shipengqi/component-base/json"
	"github.com/shipengqi/golib/convutil"
	"github.com/shipengqi/golib/fsutil"

	"github.com/shipengqi/jaguar/internal/actions/codegen/options"
)

type Config struct {
	*options.Options

	BuildTagSlice []string
	TypeSlice     []string
	OriginArgs    []string
	TargetDir     string
}

func (c *Config) String() string {
	data, _ := json.Marshal(c)

	return convutil.B2S(data)
}

// CreateConfigFromOptions creates a running configuration instance based
// on a given command line.
func CreateConfigFromOptions(opts *options.Options, args []string) (*Config, error) {
	cfg := &Config{
		Options: opts,
	}
	if opts.BuildTags != "" {
		tags := strings.Split(opts.BuildTags, ",")
		for _, v := range tags {
			cfg.BuildTagSlice = append(cfg.BuildTagSlice, strings.TrimSpace(v))
		}
	}
	cfg.TypeSlice = strings.Split(opts.Types, ",")

	if len(args) == 0 {
		// Default: process whole package in current directory.
		args = []string{"."}
	}
	cfg.OriginArgs = args

	// TODO: accept other patterns for packages (directories, list of files, import paths, etc).
	if len(args) == 1 && fsutil.IsDir(args[0]) {
		cfg.TargetDir = args[0]
	} else {
		if len(opts.BuildTags) != 0 {
			return nil, errors.New("--build-tags option applies only to directories, not when files are specified")
		}

		cfg.TargetDir = filepath.Dir(args[0])
	}

	return cfg, nil
}
