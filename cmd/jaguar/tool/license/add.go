package license

import (
	"github.com/shipengqi/jcli"

	"github.com/shipengqi/jaguar/internal/actions/license"
	"github.com/shipengqi/jaguar/internal/actions/license/config"
	"github.com/shipengqi/jaguar/internal/actions/license/options"
	"github.com/shipengqi/jaguar/internal/pkg/utils/cmdutils"
)

const addCmdDesc = "Add the copyright license headers for source code files."

func newAddCmd() *jcli.Command {
	o := options.New()
	c := jcli.NewCommand(
		license.ActionNameAdd,
		addCmdDesc,
		jcli.WithCommandDesc(cmdutils.SubCmdDesc(addCmdDesc)),
		jcli.WithCommandCliOptions(o),
		jcli.WithCommandRunFunc(func(_ *jcli.Command, args []string) error {
			if len(args) == 0 {
				return nil
			}
			cfg, err := config.CreateConfigFromOptions(o)
			if err != nil {
				return err
			}
			return license.NewAddLicenseAction(cfg, args).Execute()
		}),
	)

	return c
}
