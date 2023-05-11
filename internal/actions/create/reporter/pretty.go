package reporter

import (
	"strings"

	"github.com/fatih/color"
)

func PrettyPrefix(msg string) string {
	return prettyPrefix(msg).String()
}

func Pretty(msg, status string) string {
	b := prettyPrefix(msg)
	b.WriteString(status)
	b.WriteString(" ]")
	return b.String()
}

func prettyPrefix(msg string) *strings.Builder {
	var b strings.Builder
	var padding string
	ml := len(msg)
	pl := defaultTermWidth - ml - 18
	if pl > 0 {
		padding = strings.Repeat(".", pl)
	}
	b.WriteString(msg)
	b.WriteString(" ")
	b.WriteString(padding)
	b.WriteString(" ")
	b.WriteString("[ ")
	return &b
}

func PrettyWithColor(msg, status string, c color.Attribute) string {
	b := prettyPrefix(msg)
	b.WriteString(Colorize(status, c))
	b.WriteString(" ]")
	return b.String()
}

// Colorize a string based on given color.
func Colorize(s string, attrs ...color.Attribute) string {
	co := color.New(attrs...)
	// fmt.Sprintf("\033[1;%s;40m%s\033[0m", strconv.Itoa(int(c)), s)
	return co.Sprint(s)
}

func ColorizeStatus(s string) string {
	switch strings.ToLower(s) {
	case "ok", "pass", "success":
		return color.GreenString(s)
	case "failed", "error":
		return color.RedString(s)
	case "warn", "warning", "skip":
		return color.YellowString(s)
	default:
		return color.WhiteString(s)
	}
}
