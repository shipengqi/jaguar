package main

import (
	"fmt"

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

const rootDesc = "A scaffold to quickly create a Go project."

func main() {
	app := jcli.New(
		"jaguar",
		jcli.WithDesc(jcli.IconBlue(fmt.Sprintf(iconFmt, rootDesc))),
		jcli.DisableConfig(),
		jcli.WithRunFunc(func() error {
			return nil
		}),
	)

	app.AddCommands(newCmd())

	app.Run()
}

func subdesc() string {
	if log.EncodedFilename() != "" {
		desc := fmt.Sprintf("A Log: %s\n", log.EncodedFilename())
		return jcli.IconBlue(fmt.Sprintf(iconFmt, desc))
	}
	return ""
}
