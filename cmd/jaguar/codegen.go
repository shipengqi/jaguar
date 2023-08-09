package main

import (
	"github.com/shipengqi/jcli"

	"github.com/shipengqi/jaguar/internal/actions/codegen"
	"github.com/shipengqi/jaguar/internal/actions/codegen/config"
	"github.com/shipengqi/jaguar/internal/actions/codegen/options"
	"github.com/shipengqi/jaguar/internal/pkg/utils/cmdutils"
)

const codeGenCmdDesc = "Automatically generate error codes for API skeleton."

func newCodeGenCmd() *jcli.Command {
	o := options.New()
	c := jcli.NewCommand(
		codegen.ActionName,
		codeGenCmdDesc,
		jcli.WithCommandDesc(cmdutils.SubCmdDesc(codeGenCmdDesc)),
		jcli.WithCommandAliases(codegen.ActionNameAlias),
		jcli.WithCommandCliOptions(o),
		jcli.WithCommandRunFunc(func(_ *jcli.Command, args []string) error {
			cfg, err := config.CreateConfigFromOptions(o, args)
			if err != nil {
				return err
			}
			a := codegen.NewAction(cfg)
			return a.Execute()
		}),
	)

	return c
}
