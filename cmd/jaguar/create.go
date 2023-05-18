package main

import (
	"github.com/shipengqi/jcli"

	"github.com/shipengqi/jaguar/internal/actions/create"
	"github.com/shipengqi/jaguar/internal/actions/create/options"
)

func newCreateCmd() *jcli.Command {
	o := options.New()

	c := jcli.NewCommand(
		create.ActionName,
		"Creates a new Go project.",
		jcli.WithCommandDesc(subdesc()),
		jcli.WithCommandAliases(create.ActionNameAlias, create.ActionNameAliasShort),
		jcli.WithCommandCliOptions(o),
		jcli.WithCommandRunFunc(func(args []string) error {
			a := create.NewAction(o, args)
			return a.Execute()
		}),
	)

	return c
}
