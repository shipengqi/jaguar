package main

import (
	"github.com/spf13/cobra"

	"github.com/shipengqi/jaguar/internal/create"
)

const createCmdDesc = "Creates a new application project from a scaffold template."

func newCreateCmd() *cobra.Command {
	cfg := create.NewConfig()

	cmd := &cobra.Command{
		Use:     create.ActionName + " [project-name]",
		Short:   createCmdDesc,
		Aliases: []string{create.ActionNameAlias, create.ActionNameAliasShort},
		Args:    cobra.MaximumNArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			if len(args) > 0 {
				cfg.ProjectName = args[0]
			}
			return create.Run(cfg)
		},
		SilenceUsage: true,
	}

	cfg.AddFlags(cmd.Flags())
	return cmd
}
