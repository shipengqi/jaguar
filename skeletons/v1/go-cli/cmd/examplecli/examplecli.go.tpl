package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"{{ .App.ModuleName }}/internal/pkg/utils/cmdutils"
	"{{ .App.ModuleName }}/pkg/xlog"
)

const rootDesc = "An example of a CLI application created by the Jaguar CLI."

func main() {
	defer finally()

	root := &cobra.Command{
		Use:           "examplecli",
		Short:         rootDesc,
		Long:          cmdutils.RootCmdDesc(rootDesc),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Help()
		},
	}
	root.Flags().SortFlags = false
	cobra.EnableCommandSorting = false

	root.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		logInitializer()
	}

	root.AddCommand(newHelloCmd())

	if err := root.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func logInitializer() {
	if cmdutils.IsHelpOrVersionCmd() {
		return
	}

	logDir := fmt.Sprintf("%s/{{ .App.NormalizedName }}/logs", os.TempDir())
	if err := os.MkdirAll(logDir, 0o755); err != nil {
		panic(err)
	}

	xlog.Init(&xlog.Options{
		Level:       "debug",
		OutputPaths: []string{filepath.Join(logDir, filenameEncoder())},
	})

	xlog.Debugf("command: %s", strings.Join(os.Args, " "))
}

func filenameEncoder() string {
	return fmt.Sprintf("%s.%s.log", filepath.Base(os.Args[0]), time.Now().Format("20060102150405"))
}

func finally() {
	_ = xlog.Sync()
	// makes the cursor visible
	_, _ = fmt.Fprint(os.Stdout, "\033[?25h")
}
