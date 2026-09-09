package xlog

import (
	"github.com/spf13/pflag"
)

// Options holds the configuration for the logger.
type Options struct {
	Level          string `json:"level"          mapstructure:"level"`
	Format         string `json:"format"         mapstructure:"format"`
	OutputPaths    []string `json:"output-paths" mapstructure:"output-paths"`
	EnableColor    bool   `json:"enable-color"   mapstructure:"enable-color"`
	DisableCaller  bool   `json:"disable-caller" mapstructure:"disable-caller"`
}

// NewOptions returns default Options.
func NewOptions() *Options {
	return &Options{
		Level:       "info",
		Format:      "console",
		OutputPaths: []string{"stdout"},
		EnableColor: true,
	}
}

// Validate validates the options.
func (o *Options) Validate() []error {
	return nil
}

// AddFlags adds flags related to logger to the specified FlagSet.
func (o *Options) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Level, "log.level", o.Level, "Minimum log output level (debug, info, warn, error).")
	fs.StringVar(&o.Format, "log.format", o.Format, "Log output format (console or json).")
	fs.StringSliceVar(&o.OutputPaths, "log.output-paths", o.OutputPaths, "Output paths for the log.")
	fs.BoolVar(&o.EnableColor, "log.enable-color", o.EnableColor, "Enable output ansi colors in plain format.")
	fs.BoolVar(&o.DisableCaller, "log.disable-caller", o.DisableCaller, "Disable output of caller information.")
}
