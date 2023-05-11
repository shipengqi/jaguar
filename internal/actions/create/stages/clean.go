package stages

import (
	"os"

	"github.com/shipengqi/jaguar/internal/actions/create/reporter"
	"github.com/shipengqi/jaguar/internal/actions/create/types"
)

type cleanStage struct {
	title   string
	project string
}

func NewCleanStage(project string) Interface {
	return &cleanStage{
		title:   titleClean,
		project: project,
	}
}

func (s *cleanStage) Run() error {
	bar := reporter.Start(s.title)

	err := os.RemoveAll(s.project)
	if err != nil {
		bar.StopWithStatus(types.StatusFailed)
		return err
	}
	bar.StopWithStatus(types.StatusOk)
	return nil
}
