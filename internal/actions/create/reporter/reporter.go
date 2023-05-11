package reporter

import (
	"fmt"
	"io"

	"github.com/shipengqi/component-base/term"
	"github.com/shipengqi/jaguar/internal/pkg/spinner"
	"github.com/shipengqi/log"
)

var (
	defaultReporter      = spinner.New()
	defaultTermWidth int = 64
)

func Start(msg string) *spinner.Spinner {
	prefix := PrettyPrefix(msg)
	defaultReporter.Reset().WithPrefix(prefix).WithSuffix(" ]").Start()
	return defaultReporter
}

func Startf(template string, args ...interface{}) *spinner.Spinner {
	prefix := PrettyPrefix(fmt.Sprintf(template, args...))
	defaultReporter.Reset().WithPrefix(prefix).WithSuffix(" ]").Start()
	return defaultReporter
}

func End(status string) {
	endStr := defaultReporter.StopWithStatus(ColorizeStatus(status))
	log.Debug(endStr)
}

func Init(w io.Writer) {
	defaultTermWidth, _, _ = term.TerminalSize(w)
}
