package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/shipengqi/golib/cliutil"
	"github.com/shipengqi/golib/fsutil"
	"github.com/shipengqi/jcli"
	"github.com/shipengqi/log"
	"github.com/spf13/cobra"
)

const iconFmt = `     __                                   
    |__|____     ____  __ _______ _______ 
    |  \__  \   / ___\|  |  \__  \\_  __ \
    |  |/ __ \_/ /_/  >  |  // __ \|  | \/
/\__|  (____  /\___  /|____/(____  /__|   
\______|    \//_____/            \/    
%s`

const rootDesc = "A scaffold to quickly create a Go project."

func main() {
	defer finally()
	cobra.OnInitialize(logInitializer)

	app := jcli.New(
		"examplecli",
		jcli.WithDesc(jcli.IconBlue(fmt.Sprintf(iconFmt, rootDesc))),
		jcli.DisableConfig(),
		jcli.DisableCommandSorting(),
		jcli.WithOnSignalReceived(func(_ os.Signal) {
			finally()
			os.Exit(0)
		}),
	)

	app.Command().PersistentPreRun = func(cmd *cobra.Command, args []string) {
		log.Info(subdesc(""))
	}

	app.AddCommands(
		newHelloCmd(),
	)

	app.Run()
}

func subdesc(cmdDesc string) string {
	if log.EncodedFilename() != "" {
		desc := fmt.Sprintf("A Log: %s\n", log.EncodedFilename())
		return jcli.IconBlue(fmt.Sprintf(iconFmt, desc))
	}
	if cmdDesc != "" {
		return jcli.IconBlue(fmt.Sprintf(iconFmt, cmdDesc))
	}
	return ""
}

func logInitializer() {
	if isHelpCmd(os.Args) {
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

func isHelpCmd(args []string) bool {
	if len(args) == 1 {
		return true
	}
	if len(args) > 1 && args[1] == "help" {
		return true
	}

	if _, ok := cliutil.RetrieveFlagFromCLI("--help", "-h"); ok {
		return true
	}
	if _, ok := cliutil.RetrieveFlagFromCLI("--version", "-v"); ok {
		return true
	}
	return false
}

func filenameEncoder() string {
	return fmt.Sprintf("%s.%s.log", filepath.Base(os.Args[0]), time.Now().Format("20060102150405"))
}

func finally() {
	_ = log.Close()
	// makes the cursor visible
	_, _ = fmt.Fprint(os.Stdout, "\033[?25h")
}
