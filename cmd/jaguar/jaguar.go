package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/shipengqi/golib/fsutil"
	"github.com/shipengqi/jcli"
	"github.com/shipengqi/log"
	"github.com/spf13/cobra"

	"github.com/shipengqi/jaguar/cmd/jaguar/license"
	"github.com/shipengqi/jaguar/internal/pkg/utils/cmdutils"
)

const rootDesc = "A scaffold to quickly create a Go project."

func main() {
	defer finally()
	cobra.OnInitialize(logInitializer)

	app := jcli.NewCommand(
		"jaguar",
		rootDesc,
		jcli.WithCommandDesc(cmdutils.RootCmdDesc(rootDesc)),
	)

	app.CobraCommand().PersistentPreRun = func(cmd *cobra.Command, args []string) {
		log.Info(cmdutils.SubCmdDesc(""))
	}

	app.AddCommands(
		newCreateCmd(),
		newCodeGenCmd(),
		license.NewCmd(),
	)
	cobra.EnableCommandSorting = false

	app.Run()
}

func logInitializer() {
	if cmdutils.IsHelpCmd(os.Args) {
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
	}

	err := fsutil.MkDirAll(opts.Output)
	if err != nil {
		panic(err)
	}
	log.Configure(
		opts,
		log.WithFilenameEncoder(filenameEncoder),
	)

	log.Debugf("track: %s", strings.Join(os.Args, " "))
}

func filenameEncoder() string {
	return fmt.Sprintf("%s.%s.log", filepath.Base(os.Args[0]), time.Now().Format("20060102150405"))
}

func finally() {
	_ = log.Close()
	// makes the cursor visible
	_, _ = fmt.Fprint(os.Stdout, "\033[?25h")
}
