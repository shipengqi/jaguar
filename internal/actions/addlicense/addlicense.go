package addlicense

import (
	"github.com/shipengqi/action"

	"github.com/shipengqi/jaguar/internal/actions/addlicense/options"
)

const (
	ActionName      = "addlicense"
	ActionNameAlias = "addl"
)

func NewAction(opts *options.Options) *action.Action {
	act := &action.Action{
		Name: ActionName,
	}

	_ = act.AddAction()

	return act
}
