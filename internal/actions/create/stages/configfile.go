package stages

import (
	"fmt"
	"path"

	"github.com/shipengqi/jaguar/internal/actions/create/helpers"
	"github.com/shipengqi/jaguar/internal/actions/create/reporter"
	"github.com/shipengqi/jaguar/internal/actions/create/types"
	"github.com/shipengqi/log"
)

type configFileStage struct {
	title    string
	filename string
	project  string
}

func NewConfigFileStage(filename, project string) Interface {
	return &configFileStage{
		title:    fmt.Sprintf("Creating %s", filename),
		filename: filename,
		project:  project,
	}
}

func (s *configFileStage) Run() error {
	bar := reporter.Start(s.title)

	fs, ok := fsmap[s.filename]
	if !ok {
		log.Debugf("not found %s", s.filename)
		return nil
	}
	err := helpers.CopyFile(fs, s.filename, path.Join(s.project, s.filename))
	if err != nil {
		bar.StopWithStatus(types.StatusFailed)
		return err
	}
	bar.StopWithStatus(types.StatusOk)
	return nil
}
