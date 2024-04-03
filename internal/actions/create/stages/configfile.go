package stages

// import (
// 	"fmt"
// 	"path"
//
// 	"github.com/shipengqi/log"
//
// 	"github.com/shipengqi/jaguar/internal/actions/create/helpers"
// )
//
// type configFileStage struct {
// 	title    string
// 	filename string
// 	project  string
// }
//
// func NewConfigFileStage(filename, project string) Interface {
// 	return &configFileStage{
// 		title:    fmt.Sprintf("Creating %s", filename),
// 		filename: filename,
// 		project:  project,
// 	}
// }

// func (s *configFileStage) Run() error {
// 	fs, ok := fsmap[s.filename]
// 	if !ok {
// 		log.Debugf("not found %s", s.filename)
// 		return nil
// 	}
// 	return helpers.CopyFile(fs, s.filename, path.Join(s.project, s.filename))
// }
