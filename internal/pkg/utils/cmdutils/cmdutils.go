package cmdutils

import (
	"fmt"
	"os"

	"github.com/shipengqi/golib/cliutil"
	"github.com/shipengqi/jcli"
	"github.com/shipengqi/log"
)

const iconFmt = `     __                                   
    |__|____     ____  __ _______ _______ 
    |  \__  \   / ___\|  |  \__  \\_  __ \
    |  |/ __ \_/ /_/  >  |  // __ \|  | \/
/\__|  (____  /\___  /|____/(____  /__|   
\______|    \//_____/            \/    
%s`

func IsHelpCmd() bool {
	args := os.Args
	if len(args) == 1 {
		return true
	}
	if len(args) > 1 && args[1] == "help" {
		return true
	}

	if _, ok := cliutil.RetrieveFlagFromCLI("--help", "-h"); ok {
		return true
	}
	return false
}

func IsVersionCmd() bool {
	if _, ok := cliutil.RetrieveFlagFromCLI("--version", "-v"); ok {
		return true
	}
	return false
}

func IsHelpOrVersionCmd() bool {
	return IsHelpCmd() || IsVersionCmd()
}

func RootCmdDesc(rootDesc string) string {
	return jcli.IconBlue(fmt.Sprintf(iconFmt, rootDesc))
}

func SubCmdDesc(cmdDesc string) string {
	if log.EncodedFilename() != "" {
		desc := fmt.Sprintf("A Log: %s\n", log.EncodedFilename())
		return jcli.IconBlue(fmt.Sprintf(iconFmt, desc))
	}
	if cmdDesc != "" {
		return jcli.IconBlue(fmt.Sprintf(iconFmt, cmdDesc))
	}
	return ""
}
