package tool

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/shipengqi/jaguar/internal/actions/codegen"
	"github.com/shipengqi/jaguar/internal/actions/codegen/config"
	"github.com/shipengqi/jaguar/internal/actions/codegen/options"
)

const codeGenCmdDesc = "Automatically generate error codes for API skeleton."

func newCodeGenCmd() *cobra.Command {
	o := options.New()
	cmd := &cobra.Command{
		Use:          codegen.ActionName + " [files...]",
		Short:        codeGenCmdDesc,
		Aliases:      []string{codegen.ActionNameAlias},
		SilenceUsage: true,
		RunE: func(_ *cobra.Command, args []string) error {
			if errs := o.Validate(); len(errs) > 0 {
				return fmt.Errorf("%v", errs[0])
			}
			cfg, err := config.CreateConfigFromOptions(o, args)
			if err != nil {
				return err
			}
			return codegen.NewAction(cfg)()
		},
	}
	o.AddFlags(cmd.Flags())
	return cmd
}
