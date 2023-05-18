package codegen

import (
	"github.com/shipengqi/action"

	"github.com/shipengqi/jaguar/internal/actions/codegen/options"
)

const (
	ActionName      = "codegen"
	ActionNameAlias = "cgen"
)

func NewAction(opts *options.Options) *action.Action {
	act := &action.Action{
		Name: ActionName,
	}

	_ = act.AddAction()

	return act
}
