package stages

import (
	"fmt"
	"strings"

	"github.com/shipengqi/jaguar/internal/actions/create/helpers"
	"github.com/shipengqi/jaguar/internal/actions/create/reporter"
	"github.com/shipengqi/jaguar/internal/actions/create/types"
)

type projectStage struct {
	title    string
	skeleton string
	project  string
}

func NewProjectStage(skeleton, project string) Interface {
	return &projectStage{
		title:    fmt.Sprintf(titleProject, strings.ToUpper(skeleton)),
		skeleton: skeleton,
		project:  project,
	}
}

func (s *projectStage) Run() error {
	var count int
	fs := fsmap[s.skeleton]
	err := helpers.CalculateFilesFromEmbedFS(fs, s.skeleton, &count)
	if err != nil {
		return err
	}
	bar := reporter.Start(s.title)

	err = helpers.Copy(fs, s.skeleton, s.project)
	if err != nil {
		bar.StopWithStatus(types.StatusFailed)
		return err
	}
	bar.StopWithStatus(types.StatusOk)
	return nil
}
