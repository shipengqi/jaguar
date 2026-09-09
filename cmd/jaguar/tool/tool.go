package tool

import (
	"github.com/spf13/cobra"

	"github.com/shipengqi/jaguar/cmd/jaguar/tool/license"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "tool",
		Short:        "Run specified jaguar tool.",
		SilenceUsage: true,
	}
	cmd.AddCommand(
		newCodeGenCmd(),
		license.NewCmd(),
	)
	return cmd
}
