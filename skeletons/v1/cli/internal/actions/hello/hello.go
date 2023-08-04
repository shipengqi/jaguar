package hello

import (
	"github.com/shipengqi/action"
	"github.com/shipengqi/log"

	"github.com/jaguar/cliskeleton/internal/actions/hello/config"
	"github.com/jaguar/cliskeleton/internal/actions/hello/options"
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
			log.Infof("Hello, %s.", cfg.Name)
			return nil
		},
		PersistentPreRun: func(act *action.Action) error {
			log.Infof("[%s] PersistentPreRun.", act.Name)
			return nil
		},
		PersistentPostRun: func(act *action.Action) error {
			log.Infof("[%s] PersistentPostRun.", act.Name)
			return nil
		},
	}
	_ = act.AddAction(
		newSubAction(cfg),
	)

	return act, nil
}
