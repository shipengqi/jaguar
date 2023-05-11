package stages

import (
	"time"

	"github.com/shipengqi/jaguar/internal/actions/create/reporter"
	"github.com/shipengqi/jaguar/internal/actions/create/types"
)

type initModStage struct {
	title   string
	project string
}

func NewInitModStage(project string) Interface {
	return &initModStage{
		title:   titleInitMod,
		project: project,
	}
}

func (s *initModStage) Run() error {
	bar := reporter.Start(s.title)

	time.Sleep(5 * time.Second)
	bar.StopWithStatus(types.StatusOk)
	return nil
}
