package main

import (
	"github.com/shipengqi/jaguar/internal/actions/create"
	"github.com/shipengqi/jaguar/internal/actions/create/options"
	"github.com/shipengqi/jcli"
)

func newCmd() *jcli.Command {
	o := options.New()

	c := jcli.NewCommand(
		create.ActionName,
		"Creates a new Go project.",
		jcli.WithCommandAliases(create.ActionNameAlias),
		jcli.WithCommandCliOptions(o),
		jcli.WithCommandRunFunc(func(args []string) error {
			a := create.NewAction(o)
			return a.Execute()
		}),
	)

	return c
}
