package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/shipengqi/component-base/version"
	"github.com/shipengqi/golib/fsutil"
	"github.com/shipengqi/jcli"
	"github.com/shipengqi/log"
	"github.com/spf13/cobra"

	"github.com/shipengqi/jaguar/cmd/jaguar/tool"
	"github.com/shipengqi/jaguar/internal/pkg/utils/cmdutils"
)

const rootDesc = "A scaffold that makes it easy to create amazing Go applications."

func main() {
	defer finally()
	cobra.OnInitialize(logInitializer)

	app := jcli.NewCommand(
		"jaguar",
		rootDesc,
		jcli.WithCommandDesc(cmdutils.RootCmdDesc(rootDesc)),
		jcli.EnableCommandVersion(),
		jcli.WithCommandRunFunc(func(cmd *jcli.Command, _ []string) error {
			log.Infof("%s Version: \n%s", "=======>", version.Get().String())
			return cmd.Help()
		}),
	)

	app.CobraCommand().PersistentPreRun = func(_ *cobra.Command, _ []string) {
		// The log content should not contain the logo string,
		// so use 'fmt' instead of 'log'
		desc := ""
		if cmdutils.IsVersionCmd() {
			desc = cmdutils.RootCmdDesc(rootDesc) + "\n"
		} else {
			desc = cmdutils.SubCmdDesc("")
		}
		if desc == "" {
			return
		}
		fmt.Println(desc)
	}
	app.AddCommands(
		newCreateCmd(),
		tool.NewCmd(),
	)
	cobra.EnableCommandSorting = false

	app.Run()
}

func logInitializer() {
	if cmdutils.IsHelpOrVersionCmd() {
		return
	}

	opts := &log.Options{
		DisableRotate:        true,
		DisableFileCaller:    true,
		DisableConsoleCaller: true,
		DisableConsoleLevel:  true,
		DisableConsoleTime:   true,
		Output:               fmt.Sprintf("%s/jaguar/logs", os.TempDir()),
		ConsoleLevel:         log.InfoLevel.String(),
		FileLevel:            log.DebugLevel.String(),
		FilenameEncoder:      filenameEncoder,
	}

	err := fsutil.MkDirAll(opts.Output)
	if err != nil {
		panic(err)
	}
	log.Configure(opts)

	log.Debugf("command: %s", strings.Join(os.Args, " "))
}

func filenameEncoder() string {
	return fmt.Sprintf("%s.%s.log", filepath.Base(os.Args[0]), time.Now().Format("20060102150405"))
}

func finally() {
	_ = log.Close()
	// makes the cursor visible
	_, _ = fmt.Fprint(os.Stdout, "\033[?25h")
}
