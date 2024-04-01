package create

import (
	"github.com/charmbracelet/huh"
	"github.com/shipengqi/golib/strutil"

	"github.com/shipengqi/jaguar/internal/actions/create/config"
	"github.com/shipengqi/jaguar/internal/actions/create/options"
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
	if strutil.IsEmpty(cfg.ProjectName) {
		groups = append(groups, huh.NewGroup(ui.ProjectNameInput(cfg)))
	}
	if strutil.IsEmpty(cfg.ModuleName) {
		groups = append(groups, huh.NewGroup(ui.GoModuleNameInput(cfg)))
	}
	return huh.NewForm(groups...).Run()
}

func runProjectTypeForm(cfg *config.Config) error {
	if _, ok := ui.SupportedProjectTypes[cfg.ProjectType]; !ok || strutil.IsEmpty(cfg.ProjectType) {
		return huh.NewForm(huh.NewGroup(ui.ProjectTypeSelect(cfg))).Run()
	}
	return nil
}

func runToolsForm(cfg *config.Config) error {
	var groups []*huh.Group
	if cfg.ProjectType == types.ProjectTypeAPI {
		if _, ok := ui.SupportedGoFrameworks[cfg.GoFramework]; !ok || strutil.IsEmpty(cfg.GoFramework) {
			groups = append(groups, huh.NewGroup(ui.GoFrameworkSelect(cfg)))
		}
	}

	if !cfg.Changed(options.FlagUseGolangCILint) {
		groups = append(groups, huh.NewGroup(ui.IsUseGolangCILintConfirm(cfg)))
	}
	if !cfg.Changed(options.FlagUseGoReleaser) {
		groups = append(groups, huh.NewGroup(ui.IsUseGoReleaserConfirm(cfg)))
	}
	if !cfg.Changed(options.FlagUseGSemver) {
		groups = append(groups, huh.NewGroup(ui.IsUseGSemverConfirm(cfg)))
	}

	return huh.NewForm(groups...).Run()
}
