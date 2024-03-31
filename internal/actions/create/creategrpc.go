package create

import (
	"github.com/shipengqi/action"
	"github.com/shipengqi/jaguar/internal/actions/create/config"
	"github.com/shipengqi/jaguar/internal/actions/create/stages"
	"github.com/shipengqi/jaguar/internal/actions/create/types"
)

func newCreateGRPCAction(cfg *config.Config) *action.Action {
	var allStages []stages.Interface
	act := &action.Action{
		Name: ActionName,
		Executable: func(_ *action.Action) bool {
			return cfg.ProjectType == types.ProjectTypeGRPC
		},
		PreRun: func(act *action.Action) (err error) {
			stages.Init(cfg.SkeletonVersion)
			if cfg.Force {
				allStages = append(allStages, stages.NewCleanStage(cfg.ProjectName))
			}
			return
		},
		Run: func(act *action.Action) error { return nil },
	}

	return act
}
