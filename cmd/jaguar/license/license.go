package license

import (
	"github.com/shipengqi/jcli"

	"github.com/shipengqi/jaguar/internal/actions/license"
	"github.com/shipengqi/jaguar/internal/pkg/utils/cmdutils"
)

const licenseCmdDesc = "Ensures source code files have copyright license headers by scanning directory patterns recursively."

func NewCmd() *jcli.Command {
	c := jcli.NewCommand(
		license.ActionName,
		licenseCmdDesc,
		jcli.WithCommandDesc(cmdutils.SubCmdDesc(licenseCmdDesc)),
	)

	c.AddCommands(
		newAddCmd(),
		newCheckCmd(),
		newRemoveCmd(),
	)

	return c
}
