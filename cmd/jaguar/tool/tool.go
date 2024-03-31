package tool

import (
	"github.com/shipengqi/jcli"

	"github.com/shipengqi/jaguar/cmd/jaguar/tool/license"
	"github.com/shipengqi/jaguar/internal/pkg/utils/cmdutils"
)

func NewCmd() *jcli.Command {
	c := jcli.NewCommand(
		"tool",
		"run specified jaguar tool.",
		jcli.WithCommandDesc(cmdutils.SubCmdDesc("Tool runs the jaguar tool command identified by the arguments.")),
	)

	c.AddCommands(
		newCodeGenCmd(),
		license.NewCmd(),
	)

	return c
}
