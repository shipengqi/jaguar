package main

import (
	"github.com/shipengqi/jcli"

	"github.com/shipengqi/jaguar/internal/actions/addlicense"
	"github.com/shipengqi/jaguar/internal/actions/addlicense/options"
)

const addLCmdDesc = "The program ensures source code files have copyright license headers by scanning directory patterns recursively."

func newAddLicenseCmd() *jcli.Command {
	o := options.New()
	c := jcli.NewCommand(
		addlicense.ActionName,
		addLCmdDesc,
		jcli.WithCommandDesc(subdesc(addLCmdDesc)),
		jcli.WithCommandAliases(addlicense.ActionNameAlias),
		jcli.WithCommandCliOptions(o),
		jcli.WithCommandRunFunc(func(args []string) error {
			a := addlicense.NewAction(o)
			return a.Execute()
		}),
	)

	return c
}
