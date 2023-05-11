package create

import (
	"fmt"
	"os"

	"github.com/shipengqi/action"
	"github.com/shipengqi/component-base/term"
	"github.com/shipengqi/jaguar/internal/actions/create/config"
	"github.com/shipengqi/jaguar/internal/actions/create/options"
	"github.com/shipengqi/jaguar/internal/actions/create/reporter"
)

const (
	ActionName       = "new"
	ActionNameAlias  = "create"
	ActionNameAlias2 = "n"
)

func NewAction(opts *options.Options, args []string) *action.Action {
	width, _, err := term.TerminalSize(os.Stdout)
	if err != nil {
		panic(err)
	}
	reporter.Init(os.Stdout)
	fmt.Println(width)

	cfg, _ := config.CreateConfigFromOptions(opts, args)
	// Todo move prerun to options.Complete
	err = prerun(cfg)
	if err != nil {
		panic(err)
	}
	act := &action.Action{
		Name: ActionName,
	}

	_ = act.AddAction(
		newCreateAPIAction(cfg),
		newCreateCLIAction(cfg),
		newCreateGRPCAction(cfg),
	)

	return act
}
