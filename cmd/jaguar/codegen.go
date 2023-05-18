package main

import (
	"github.com/shipengqi/jcli"

	"github.com/shipengqi/jaguar/internal/actions/codegen"
	"github.com/shipengqi/jaguar/internal/actions/codegen/options"
)

func newCodeGenCmd() *jcli.Command {
	o := options.New()

	c := jcli.NewCommand(
		codegen.ActionName,
		"Automatically generate error codes.",
		jcli.WithCommandDesc(subdesc()),
		jcli.WithCommandAliases(codegen.ActionNameAlias),
		jcli.WithCommandCliOptions(o),
		jcli.WithCommandRunFunc(func(args []string) error {
			a := codegen.NewAction(o)
			return a.Execute()
		}),
	)

	return c
}
