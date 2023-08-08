package license

import (
	"github.com/shipengqi/jcli"

	"github.com/shipengqi/jaguar/internal/actions/license"
	"github.com/shipengqi/jaguar/internal/actions/license/config"
	"github.com/shipengqi/jaguar/internal/actions/license/options"
	"github.com/shipengqi/jaguar/internal/pkg/utils/cmdutils"
)

const rmCmdDesc = "Remove copyright license headers contained in source code files."

func newRemoveCmd() *jcli.Command {
	o := options.New()
	c := jcli.NewCommand(
		license.ActionNameRemove,
		rmCmdDesc,
		jcli.WithCommandDesc(cmdutils.SubCmdDesc(rmCmdDesc)),
		jcli.WithCommandCliOptions(o),
		jcli.WithCommandRunFunc(func(cmd *jcli.Command, args []string) error {
			if len(args) == 0 {
				return nil
			}
			cfg, err := config.CreateConfigFromOptions(o)
			if err != nil {
				return err
			}
			return license.NewRemoveLicenseAction(cfg, args).Execute()
		}),
	)

	return c
}
