package config

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/shipengqi/jaguar/internal/actions/codegen/options"
)

type Config struct {
	*options.Options

	BuildTagSlice []string
	TypeSlice     []string
	OriginArgs    []string
	TargetDir     string
}

// CreateConfigFromOptions creates a running configuration instance based
// on a given command line.
func CreateConfigFromOptions(opts *options.Options, args []string) (*Config, error) {
	cfg := &Config{Options: opts}
	if opts.BuildTags != "" {
		tags := strings.Split(opts.BuildTags, ",")
		for _, v := range tags {
			cfg.BuildTagSlice = append(cfg.BuildTagSlice, strings.TrimSpace(v))
		}
	}
	cfg.TypeSlice = strings.Split(opts.Types, ",")

	if len(args) == 0 {
		args = []string{"."}
	}
	cfg.OriginArgs = args

	if len(args) == 1 && isDir(args[0]) {
		cfg.TargetDir = args[0]
	} else {
		if len(opts.BuildTags) != 0 {
			return nil, errors.New("--build-tags option applies only to directories, not when files are specified")
		}
		cfg.TargetDir = filepath.Dir(args[0])
	}

	return cfg, nil
}

func isDir(p string) bool {
	fi, err := os.Stat(p)
	return err == nil && fi.IsDir()
}
