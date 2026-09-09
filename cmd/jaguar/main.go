package main

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"charm.land/lipgloss/v2"
	"github.com/spf13/cobra"

	"github.com/shipengqi/jaguar/cmd/jaguar/tool"
)

const rootDesc = "A scaffold that makes it easy to create amazing Go applications."

func main() {
	defer func() {
		// restore cursor visibility only when writing to a real terminal
		if fi, err := os.Stdout.Stat(); err == nil && (fi.Mode()&os.ModeCharDevice) != 0 {
			_, _ = fmt.Fprint(os.Stdout, "\033[?25h")
		}
	}()

	initLogger()

	root := &cobra.Command{
		Use:   "jaguar",
		Short: rootDesc,
		Long:  renderBanner(),
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Help()
		},
		SilenceUsage: true,
	}

	cobra.EnableCommandSorting = false

	root.AddCommand(newCreateCmd(), tool.NewCmd())

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

func renderBanner() string {
	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7571F9"))
	return style.Render("JAGUAR") + "\n" + rootDesc
}

func initLogger() {
	logDir := filepath.Join(os.TempDir(), "jaguar", "logs")
	_ = os.MkdirAll(logDir, 0o700)
	logFile := filepath.Join(logDir,
		fmt.Sprintf("%s.%s.log", filepath.Base(os.Args[0]), time.Now().Format("20060102150405")))
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
	if err != nil {
		return
	}
	slog.SetDefault(slog.New(slog.NewTextHandler(f, &slog.HandlerOptions{Level: slog.LevelDebug})))
	slog.Debug("command started", "args", os.Args)
}
