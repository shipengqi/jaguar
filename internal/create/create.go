package create

import (
	"context"
	"fmt"
	"os"
	"time"

	"charm.land/lipgloss/v2"
)

const (
	ActionName           = "new"
	ActionNameAlias      = "create"
	ActionNameAliasShort = "n"
)

var spinnerFrames = []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}

// Run is the entry point for the create command.
func Run(cfg *Config) error {
	if err := prerun(cfg); err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	ss := newStages(cfg)
	var stageErr error

	go func() {
		stageErr = ss.run(cancel)
	}()

	runSpinner(ctx)

	if stageErr != nil {
		_, _ = fmt.Fprintf(os.Stderr, "\n✕ Oops... Something went wrong!\n\n%v\n", stageErr)
		return stageErr
	}

	showSummary(cfg)
	return nil
}

func runSpinner(ctx context.Context) {
	dot := lipgloss.NewStyle().Foreground(lipgloss.Color("#F780E2"))
	title := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFDF5"))
	msg := title.Render(" Jaguar CLI is creating your project. Please wait...")

	i := 0
	ticker := time.NewTicker(80 * time.Millisecond)
	defer ticker.Stop()

	_, _ = fmt.Fprint(os.Stderr, "\033[?25l") // hide cursor
	for {
		select {
		case <-ctx.Done():
			_, _ = fmt.Fprint(os.Stderr, "\r\033[K\033[?25h") // clear line, show cursor
			return
		case <-ticker.C:
			frame := dot.Render(spinnerFrames[i%len(spinnerFrames)])
			_, _ = fmt.Fprintf(os.Stderr, "\r%s%s", frame, msg)
			i++
		}
	}
}

