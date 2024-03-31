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
	// Only the command itself, should print the help message
	if len(args) == 1 {
		return true
	}
	return cliutil.IsHelpCmd("-h")
}

func IsVersionCmd() bool {
	return cliutil.IsVersionCmd()
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
