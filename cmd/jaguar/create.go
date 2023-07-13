package main

import (
	"github.com/shipengqi/jcli"

	"github.com/shipengqi/jaguar/internal/actions/create"
	"github.com/shipengqi/jaguar/internal/actions/create/options"
)

const createCmdDesc = "Creates a new Go project."

func newCreateCmd() *jcli.Command {
	o := options.New()
	c := jcli.NewCommand(
		create.ActionName,
		createCmdDesc,
		jcli.WithCommandDesc(subdesc(createCmdDesc)),
		jcli.WithCommandAliases(create.ActionNameAlias, create.ActionNameAliasShort),
		jcli.WithCommandCliOptions(o),
		jcli.WithCommandRunFunc(func(args []string) error {
			a, err := create.NewAction(o, args)
			if err != nil {
				return err
			}
			return a.Execute()
		}),
	)

	return c
}
