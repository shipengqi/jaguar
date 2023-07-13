package reporter

import (
	"fmt"
	"io"

	"github.com/shipengqi/component-base/term"
	"github.com/shipengqi/jaguar/internal/pkg/spinner"
	"github.com/shipengqi/log"
)

const (
	defaultTermWidth = 64
	middleTermWidth  = 89
	maxTermWidth     = 178
)

var (
	defaultReporter = spinner.New()
	termWidth       = defaultTermWidth
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
	width, _, _ := term.TerminalSize(w)
	if width <= defaultTermWidth {
		termWidth = width
	} else if width > defaultTermWidth && width <= middleTermWidth {
		termWidth = defaultTermWidth
	} else if width > middleTermWidth && width < maxTermWidth {
		half := width / 2
		if half > defaultTermWidth {
			termWidth = half
		} else {
			termWidth = defaultTermWidth
		}
	} else if width >= maxTermWidth {
		termWidth = middleTermWidth
	}
}
