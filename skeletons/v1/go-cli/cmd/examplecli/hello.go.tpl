package main

import (
	"github.com/spf13/cobra"

	"{{ .App.ModuleName }}/internal/actions/hello"
	"{{ .App.ModuleName }}/internal/actions/hello/options"
	"{{ .App.ModuleName }}/internal/pkg/utils/cmdutils"
)

const helloCmdDesc = "Example: Say Hello."

func newHelloCmd() *cobra.Command {
	o := options.New()
	cmd := &cobra.Command{
		Use:     hello.ActionName,
		Short:   helloCmdDesc,
		Long:    cmdutils.SubCmdDesc(helloCmdDesc),
		Aliases: []string{hello.ActionNameAlias},
		RunE: func(cmd *cobra.Command, args []string) error {
			a, err := hello.NewAction(o, args)
			if err != nil {
				return err
			}
			return a.Execute()
		},
	}

	fss := o.Flags()
	for _, fs := range fss.FlagSets {
		cmd.Flags().AddFlagSet(fs)
	}
	cmd.Flags().SortFlags = false

	return cmd
}
