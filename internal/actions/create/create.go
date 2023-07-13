package create

import (
	"os"

	"github.com/shipengqi/action"

	"github.com/shipengqi/jaguar/internal/actions/create/config"
	"github.com/shipengqi/jaguar/internal/actions/create/options"
	"github.com/shipengqi/jaguar/internal/actions/create/reporter"
)

const (
	ActionName           = "new"
	ActionNameAlias      = "create"
	ActionNameAliasShort = "n"
)

func NewAction(opts *options.Options, args []string) (*action.Action, error) {
	reporter.Init(os.Stdout)

	cfg, _ := config.CreateConfigFromOptions(opts, args)
	err := prerun(cfg)
	if err != nil {
		return nil, err
	}
	act := &action.Action{
		Name: ActionName,
	}

	_ = act.AddAction(
		newCreateAPIAction(cfg),
		newCreateCLIAction(cfg),
		newCreateGRPCAction(cfg),
	)

	return act, nil
}
