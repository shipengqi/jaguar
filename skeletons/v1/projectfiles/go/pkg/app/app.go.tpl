package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"{{ .App.ModuleName }}/pkg/app/flag"
	"{{ .App.ModuleName }}/pkg/version"
	"{{ .App.ModuleName }}/pkg/xlog"
)

var progressMessage = color.GreenString("==>")

// RunFunc is the application's run callback.
type RunFunc func() error

// App is the main structure of a CLI application.
type App struct {
	name        string
	basename    string
	description string
	runfunc     RunFunc
	opts        CliOptions
	cmd         *cobra.Command
	silence     bool
	noVersion   bool
	noConfig    bool
}

// New creates a new application.
func New(name string, opts ...Option) *App {
	a := &App{name: name}
	for _, o := range opts {
		o(a)
	}
	a.cmd = a.buildCommand()
	return a
}

// Run launches the application.
func (a *App) Run() {
	if err := a.cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s %v\n", color.RedString("Error:"), err)
		os.Exit(1)
	}
}

// Command returns the underlying cobra.Command.
func (a *App) Command() *cobra.Command { return a.cmd }

func (a *App) buildCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:           normalizeName(a.basename),
		Short:         a.name,
		Long:          a.description,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	cmd.Flags().SortFlags = false

	var nfs flag.NamedFlagSets
	if a.opts != nil {
		nfs = a.opts.Flags()
		for _, fs := range nfs.FlagSets {
			cmd.Flags().AddFlagSet(fs)
		}
	}

	globalFS := nfs.FlagSet("global")
	if !a.noVersion {
		globalFS.BoolP("version", "v", false, "Print version information and quit.")
	}
	if !a.noConfig {
		globalFS.StringP("config", "c", "", "Path to configuration file.")
	}
	cmd.Flags().AddFlagSet(globalFS)

	cmd.RunE = a.run
	return cmd
}

func (a *App) run(cmd *cobra.Command, _ []string) error {
	if !a.noVersion {
		if v, _ := cmd.Flags().GetBool("version"); v {
			fmt.Println(version.Get().String())
			os.Exit(0)
		}
	}

	if !a.noConfig {
		if cfgFile, _ := cmd.Flags().GetString("config"); cfgFile != "" {
			viper.SetConfigFile(cfgFile)
		} else {
			viper.SetConfigName("apiserver")
			viper.SetConfigType("yaml")
			viper.AddConfigPath("configs")
			viper.AddConfigPath(".")
		}
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				return err
			}
		}
	}

	if !a.silence {
		xlog.Infof("%s Starting %s ...", progressMessage, a.name)
		if !a.noVersion {
			xlog.Infof("%s Version: %s", progressMessage, version.Get().Version)
		}
		if !a.noConfig {
			if cfg := viper.ConfigFileUsed(); cfg != "" {
				xlog.Infof("%s Config file: %s", progressMessage, cfg)
			}
		}
	}

	if a.opts != nil {
		if !a.noConfig {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			if err := viper.Unmarshal(a.opts); err != nil {
				return err
			}
		}
		if err := applyOptions(a.opts); err != nil {
			return err
		}
		if !a.silence {
			if p, ok := a.opts.(PrintableOptions); ok {
				xlog.Infof("%s Options: %s", progressMessage, p.String())
			}
		}
	}

	if a.runfunc != nil {
		return a.runfunc()
	}
	return nil
}

func applyOptions(opts CliOptions) error {
	if c, ok := opts.(CompletableOptions); ok {
		if err := c.Complete(); err != nil {
			return err
		}
	}
	if errs := opts.Validate(); len(errs) > 0 {
		msgs := make([]string, 0, len(errs))
		for _, e := range errs {
			msgs = append(msgs, e.Error())
		}
		return fmt.Errorf("validation failed: %s", strings.Join(msgs, "; "))
	}
	return nil
}

func normalizeName(name string) string {
	if name == "" {
		return strings.TrimSuffix(os.Args[0], ".exe")
	}
	return strings.ToLower(strings.TrimSuffix(name, ".exe"))
}
