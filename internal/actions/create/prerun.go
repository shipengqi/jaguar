package create

import (
	"github.com/charmbracelet/huh"

	"github.com/shipengqi/jaguar/internal/actions/create/config"
	"github.com/shipengqi/jaguar/internal/actions/create/types"
	"github.com/shipengqi/jaguar/internal/actions/create/ui"
)

func prerun(cfg *config.Config) error {
	if err := runBaseForm(cfg); err != nil {
		return err
	}
	if err := runProjectTypeForm(cfg); err != nil {
		return err
	}
	return runToolsForm(cfg)
}

func runBaseForm(cfg *config.Config) error {
	var groups []*huh.Group
	if cfg.ProjectName == "" {
		groups = append(groups, huh.NewGroup(ui.ProjectNameInput(cfg)))
	}
	if cfg.ModuleName == "" {
		groups = append(groups, huh.NewGroup(ui.GoModuleNameInput(cfg)))
	}
	return huh.NewForm(groups...).Run()
}

func runProjectTypeForm(cfg *config.Config) error {
	return huh.NewForm(huh.NewGroup(ui.ProjectTypeSelect(cfg))).Run()
}

func runToolsForm(cfg *config.Config) error {
	var groups []*huh.Group
	if cfg.ProjectType == types.ProjectTypeAPI && cfg.GoFramework == "" {
		groups = append(groups, huh.NewGroup(ui.GoFrameworkSelect(cfg)))
	}
	groups = append(groups,
		huh.NewGroup(ui.IsUseGolangCILintConfirm(cfg)),
		huh.NewGroup(ui.IsUseGoReleaserConfirm(cfg)),
		huh.NewGroup(ui.IsUseGSemverConfirm(cfg)),
	)
	return huh.NewForm(groups...).Run()
}
