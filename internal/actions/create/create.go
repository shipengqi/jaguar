package create

import (
	"github.com/shipengqi/action"
	"github.com/shipengqi/jaguar/internal/actions/create/stages"

	"github.com/shipengqi/jaguar/internal/actions/create/config"
	"github.com/shipengqi/jaguar/internal/actions/create/options"
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
			ss := stages.New(cfg)
			return ss.Run()
		},
	}

	return act, nil
}
