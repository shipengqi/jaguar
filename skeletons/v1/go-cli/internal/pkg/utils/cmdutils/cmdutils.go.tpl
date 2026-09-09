package cmdutils

import (
	"fmt"
	"os"
	"strings"
)

const iconFmt = `{{ .App.Logo }}
%s`

func IsHelpCmd() bool {
	args := os.Args
	if len(args) == 1 {
		return true
	}
	for _, a := range args[1:] {
		if a == "-h" || a == "--help" || a == "help" {
			return true
		}
	}
	return false
}

func IsVersionCmd() bool {
	for _, a := range os.Args[1:] {
		if a == "-v" || a == "--version" || a == "version" {
			return true
		}
	}
	return false
}

func IsHelpOrVersionCmd() bool {
	return IsHelpCmd() || IsVersionCmd()
}

func RootCmdDesc(rootDesc string) string {
	return fmt.Sprintf(iconFmt, rootDesc)
}

func SubCmdDesc(cmdDesc string) string {
	if cmdDesc != "" {
		return fmt.Sprintf(iconFmt, cmdDesc)
	}
	return ""
}

// LogFilename returns a timestamped log filename for the current binary.
func LogFilename(dir string) string {
	return strings.Join([]string{dir, fmt.Sprintf("%s.log", os.Args[0])}, string(os.PathSeparator))
}
