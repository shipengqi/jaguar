package license

import (
	"github.com/spf13/cobra"

	"github.com/shipengqi/jaguar/internal/actions/license"
	"github.com/shipengqi/jaguar/internal/actions/license/config"
	"github.com/shipengqi/jaguar/internal/actions/license/options"
)

const rmCmdDesc = "Remove copyright license headers contained in source code files."

func newRemoveCmd() *cobra.Command {
	o := options.New()
	cmd := &cobra.Command{
		Use:          license.ActionNameRemove + " [dirs...]",
		Short:        rmCmdDesc,
		SilenceUsage: true,
		RunE: func(_ *cobra.Command, args []string) error {
			if len(args) == 0 {
				return nil
			}
			cfg, err := config.CreateConfigFromOptions(o)
			if err != nil {
				return err
			}
			return license.NewRemoveLicenseAction(cfg, args)()
		},
	}
	o.AddFlags(cmd.Flags())
	return cmd
}
