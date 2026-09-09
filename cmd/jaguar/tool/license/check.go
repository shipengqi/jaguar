package license

import (
	"github.com/spf13/cobra"

	"github.com/shipengqi/jaguar/internal/actions/license"
	"github.com/shipengqi/jaguar/internal/actions/license/config"
	"github.com/shipengqi/jaguar/internal/actions/license/options"
)

const checkCmdDesc = "Checks if the copyright license headers is missing."

func newCheckCmd() *cobra.Command {
	o := options.New()
	cmd := &cobra.Command{
		Use:          license.ActionNameCheck + " [dirs...]",
		Short:        checkCmdDesc,
		SilenceUsage: true,
		RunE: func(_ *cobra.Command, args []string) error {
			if len(args) == 0 {
				return nil
			}
			cfg, err := config.CreateConfigFromOptions(o)
			if err != nil {
				return err
			}
			return license.NewCheckLicenseAction(cfg, args)()
		},
	}
	o.AddFlags(cmd.Flags())
	_ = cmd.Flags().MarkHidden("holder")
	_ = cmd.Flags().MarkHidden("year")
	_ = cmd.Flags().MarkHidden("license")
	_ = cmd.Flags().MarkHidden("license-file")
	return cmd
}
