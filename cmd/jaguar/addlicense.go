package main

import (
	"github.com/shipengqi/jcli"

	"github.com/shipengqi/jaguar/internal/actions/addlicense"
	"github.com/shipengqi/jaguar/internal/actions/addlicense/options"
)

func newAddLicenseCmd() *jcli.Command {
	o := options.New()

	c := jcli.NewCommand(
		addlicense.ActionName,
		"The program ensures source code files have copyright license headers by scanning directory patterns recursively.",
		jcli.WithCommandDesc(subdesc()),
		jcli.WithCommandAliases(addlicense.ActionNameAlias),
		jcli.WithCommandCliOptions(o),
		jcli.WithCommandRunFunc(func(args []string) error {
			a := addlicense.NewAction(o)
			return a.Execute()
		}),
	)

	return c
}
