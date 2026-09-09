package hello

import (
	"{{ .App.ModuleName }}/internal/actions/hello/config"
	"{{ .App.ModuleName }}/pkg/action"
	"{{ .App.ModuleName }}/pkg/xlog"
)

const ActionNameSub = "sub"

func newSubAction(cfg *config.Config) *action.Action {
	return &action.Action{
		Name: ActionNameSub,
		Executable: func(act *action.Action) bool {
			return cfg.Sub
		},
		Run: func(act *action.Action) error {
			xlog.Infof("[Sub Action] Hello, %s.", cfg.Name)
			return nil
		},
	}
}
