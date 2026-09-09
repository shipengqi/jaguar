package hello

import (
	"{{ .App.ModuleName }}/internal/actions/hello/config"
	"{{ .App.ModuleName }}/internal/actions/hello/options"
	"{{ .App.ModuleName }}/pkg/action"
	"{{ .App.ModuleName }}/pkg/xlog"
)

const (
	ActionName      = "hello"
	ActionNameAlias = "hi"
)

func NewAction(opts *options.Options, args []string) (*action.Action, error) {
	cfg, _ := config.CreateConfigFromOptions(opts, args)
	act := &action.Action{
		Name: ActionName,
		Executable: func(act *action.Action) bool {
			return !cfg.Sub
		},
		Run: func(act *action.Action) error {
			xlog.Infof("Hello, %s.", cfg.Name)
			return nil
		},
		PersistentPreRun: func(act *action.Action) error {
			xlog.Infof("[%s] PersistentPreRun.", act.Name)
			return nil
		},
		PersistentPostRun: func(act *action.Action) error {
			xlog.Infof("[%s] PersistentPostRun.", act.Name)
			return nil
		},
	}
	_ = act.AddAction(newSubAction(cfg))
	return act, nil
}
