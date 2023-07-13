package addlicense

import (
	"regexp"

	"github.com/shipengqi/action"

	"github.com/shipengqi/jaguar/internal/actions/addlicense/config"
	"github.com/shipengqi/jaguar/internal/actions/addlicense/options"
)

const (
	ActionName         = "addlicense"
	ActionNameAlias    = "addl"
	SubActionNameCheck = "check"
	SubActionNameAdd   = "add"
)

var patterns = struct {
	dirs  []*regexp.Regexp
	files []*regexp.Regexp
}{}

func NewAction(opts *options.Options) *action.Action {
	act := &action.Action{
		Name: ActionName,
	}
	cfg, _ := config.CreateConfigFromOptions(opts)

	_ = act.AddAction(
		newCheckAction(cfg),
		newAddAction(cfg),
	)

	return act
}
