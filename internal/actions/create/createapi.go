package create

import (
	"github.com/shipengqi/action"
	"github.com/shipengqi/jaguar/internal/actions/create/config"
	"github.com/shipengqi/jaguar/internal/actions/create/stages"
	"github.com/shipengqi/jaguar/internal/actions/create/types"
)

func newCreateAPIAction(cfg *config.Config) *action.Action {
	var allStages []stages.Interface
	act := &action.Action{
		Name: ActionName,
		Executable: func(_ *action.Action) bool {
			return cfg.ProjectType == types.ProjectTypeAPI
		},
		PreRun: func(act *action.Action) (err error) {
			stages.Init(cfg.SkeletonVersion)
			if cfg.Force {
				allStages = append(allStages, stages.NewCleanStage(cfg.ProjectName))
			}
			allStages = append(allStages,
				stages.NewProjectStage(types.ProjectTypeAPI, cfg.ProjectName),
				stages.NewConfigFileStage(stages.FilenameGitIgnore, cfg.ProjectName),
				stages.NewConfigFileStage(stages.FilenameGoCI, cfg.ProjectName),
				stages.NewConfigFileStage(stages.FilenameReleaser, cfg.ProjectName),
				stages.NewConfigFileStage(stages.FilenameSemver, cfg.ProjectName),
				stages.NewInitModStage(cfg.ProjectName),
			)
			return
		},
		Run: func(act *action.Action) (err error) {
			for _, v := range allStages {
				err = v.Run()
				if err != nil {
					return
				}
			}
			return
		},
	}

	return act
}
