package license

import (
	"github.com/shipengqi/jcli"

	"github.com/shipengqi/jaguar/internal/actions/license"
	"github.com/shipengqi/jaguar/internal/actions/license/config"
	"github.com/shipengqi/jaguar/internal/actions/license/options"
	"github.com/shipengqi/jaguar/internal/pkg/utils/cmdutils"
)

const checkCmdDesc = "Checks if the copyright license headers is missing."

func newCheckCmd() *jcli.Command {
	o := options.New()
	c := jcli.NewCommand(
		license.ActionNameCheck,
		checkCmdDesc,
		jcli.WithCommandDesc(cmdutils.SubCmdDesc(checkCmdDesc)),
		jcli.WithCommandCliOptions(o),
		jcli.WithCommandRunFunc(func(_ *jcli.Command, args []string) error {
			if len(args) == 0 {
				return nil
			}
			cfg, err := config.CreateConfigFromOptions(o)
			if err != nil {
				return err
			}
			return license.NewCheckLicenseAction(cfg, args).Execute()
		}),
	)

	// unused flags
	c.MarkHidden("holder", "year", "license", "license-file")

	return c
}
