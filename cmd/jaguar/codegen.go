package main

import (
	"github.com/shipengqi/jcli"

	"github.com/shipengqi/jaguar/internal/actions/codegen"
	"github.com/shipengqi/jaguar/internal/actions/codegen/options"
	"github.com/shipengqi/jaguar/internal/pkg/utils/cmdutils"
)

const codeGenCmdDesc = "Automatically generate error codes."

func newCodeGenCmd() *jcli.Command {
	o := options.New()
	c := jcli.NewCommand(
		codegen.ActionName,
		codeGenCmdDesc,
		jcli.WithCommandDesc(cmdutils.SubCmdDesc(codeGenCmdDesc)),
		jcli.WithCommandAliases(codegen.ActionNameAlias),
		jcli.WithCommandCliOptions(o),
		jcli.WithCommandRunFunc(func(_ *jcli.Command, args []string) error {
			a := codegen.NewAction(o)
			return a.Execute()
		}),
	)

	return c
}
