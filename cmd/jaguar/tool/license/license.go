package license

import (
	"github.com/spf13/cobra"

	"github.com/shipengqi/jaguar/internal/actions/license"
)

const licenseCmdDesc = "Ensures source code files have copyright license headers by scanning directory patterns recursively."

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          license.ActionName,
		Short:        licenseCmdDesc,
		SilenceUsage: true,
	}
	cmd.AddCommand(
		newAddCmd(),
		newCheckCmd(),
		newRemoveCmd(),
	)
	return cmd
}
