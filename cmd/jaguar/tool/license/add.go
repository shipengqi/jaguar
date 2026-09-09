package license

import (
	"github.com/spf13/cobra"

	"github.com/shipengqi/jaguar/internal/actions/license"
	"github.com/shipengqi/jaguar/internal/actions/license/config"
	"github.com/shipengqi/jaguar/internal/actions/license/options"
)

const addCmdDesc = "Add the copyright license headers for source code files."

func newAddCmd() *cobra.Command {
	o := options.New()
	cmd := &cobra.Command{
		Use:          license.ActionNameAdd + " [dirs...]",
		Short:        addCmdDesc,
		SilenceUsage: true,
		RunE: func(_ *cobra.Command, args []string) error {
			if len(args) == 0 {
				return nil
			}
			cfg, err := config.CreateConfigFromOptions(o)
			if err != nil {
				return err
			}
			return license.NewAddLicenseAction(cfg, args)()
		},
	}
	o.AddFlags(cmd.Flags())
	return cmd
}
