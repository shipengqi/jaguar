package main

import (
	"github.com/shipengqi/jcli"

	"github.com/jaguar/cliskeleton/internal/actions/hello"
	"github.com/jaguar/cliskeleton/internal/actions/hello/options"
)

const helloCmdDesc = "Example: Say Hello."

func newHelloCmd() *jcli.Command {
	o := options.New()
	c := jcli.NewCommand(
		hello.ActionName,
		helloCmdDesc,
		jcli.WithCommandDesc(subdesc(helloCmdDesc)),
		jcli.WithCommandAliases(hello.ActionNameAlias),
		jcli.WithCommandCliOptions(o),
		jcli.WithCommandRunFunc(func(cmd *jcli.Command, args []string) error {
			a, err := hello.NewAction(o, args)
			if err != nil {
				return err
			}
			return a.Execute()
		}),
	)

	return c
}
