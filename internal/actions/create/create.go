package create

import (
	"context"

	"github.com/shipengqi/action"

	"github.com/shipengqi/jaguar/internal/actions/create/config"
	"github.com/shipengqi/jaguar/internal/actions/create/options"
	"github.com/shipengqi/jaguar/internal/actions/create/stages"
	"github.com/shipengqi/jaguar/internal/actions/create/ui"
	"github.com/shipengqi/jaguar/internal/pkg/spinner"
)

const (
	ActionName           = "new"
	ActionNameAlias      = "create"
	ActionNameAliasShort = "n"
)

func NewAction(opts *options.Options, args []string) (*action.Action, error) {
	cfg, _ := config.CreateConfigFromOptions(opts, args)
	err := prerun(cfg)
	if err != nil {
		return nil, err
	}
	act := &action.Action{
		Name: ActionName,
		Run: func(act *action.Action) error {
			ctx, cancel := context.WithCancel(context.Background())
			ss := stages.New(cfg)
			var result error
			go func() {
				result = ss.Run(cancel)
			}()
			if err = spinner.New().Type(spinner.Dots).
				Title(" Jaguar CLI is creating your project. Please wait ...").
				Context(ctx).Run(); err != nil {
				return err
			}
			if result != nil {
				return result
			}

			ui.ShowSummary(cfg)
			return nil
		},
	}

	return act, nil
}
