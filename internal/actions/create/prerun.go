package create

import (
	"errors"

	"github.com/shipengqi/jaguar/internal/actions/create/config"
	"github.com/shipengqi/jaguar/internal/actions/create/types"
	"github.com/shipengqi/jaguar/internal/pkg/survey"
)

const (
	DefaultModuleMinLength  = 2
	DefaultModuleMaxLength  = 128
	DefaultProjectMinLength = 2
	DefaultProjectMaxLength = 20
)

func prerun(cfg *config.Config) error {
	var answer string
	var err error
	if cfg.Type == "" {
		var selected string
		answer, err = survey.Select("Select project type:", []string{"CLI", "API", "gRPC"})
		if err != nil {
			return err
		}
		switch answer {
		case "CLI":
			selected = types.ProjectTypeCLI
		case "API":
			selected = types.ProjectTypeAPI
		case "gRPC":
			selected = types.ProjectTypeGRPC
		default:
			return errors.New("unsupported type")
		}
		cfg.Type = selected
	}

	if cfg.ProjectName == "" {
		answer, err = survey.InputString("Please input your project name:",
			DefaultProjectMinLength, DefaultProjectMaxLength)
		if err != nil {
			return err
		}
		cfg.ProjectName = answer
	}

	if cfg.Module == "" {
		answer, err = survey.InputString("Please input your Go module name:",
			DefaultModuleMinLength, DefaultModuleMaxLength)
		if err != nil {
			return err
		}
		cfg.Module = answer
	}

	return nil
}
