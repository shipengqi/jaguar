package create

import (
	"errors"

	"github.com/shipengqi/jaguar/internal/actions/create/config"
	"github.com/shipengqi/jaguar/internal/pkg/survey"
)

const (
	DefaultModuleMinLength  = 2
	DefaultModuleMaxLength  = -1
	DefaultProjectMinLength = 2
	DefaultProjectMaxLength = 20
)

func prerun(cfg *config.Config) error {
	if cfg.Type == "" {
		var selected string
		switch survey.Select("Select project type", []string{"CLI", "API", "gRPC"}) {
		case "CLI":
			selected = ProjectTypeCLI
		case "API":
			selected = ProjectTypeAPI
		case "gRPC":
			selected = ProjectTypeGRPC
		default:
			return errors.New("unsupported type")
		}
		cfg.Type = selected
	}

	if cfg.ProjectName == "" {
		cfg.ProjectName = survey.InputString("Please input your project name",
			DefaultProjectMinLength, DefaultProjectMaxLength)
	}

	if cfg.Module == "" {
		cfg.Module = survey.InputString("Please input your Go module name",
			DefaultModuleMinLength, DefaultModuleMaxLength)
	}

	return nil
}
