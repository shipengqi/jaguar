package ui

import (
	"github.com/charmbracelet/huh"

	"github.com/shipengqi/jaguar/internal/actions/create/config"
)

// ProjectTypeSelect runs the project type select.
func ProjectTypeSelect(cfg *config.Config) *huh.Select[string] {
	return huh.NewSelect[string]().
		Title(FormProjectTypeTitle).
		Description(FormProjectTypeDescription).
		Options(
			huh.NewOption(SupportedProjectTypes["api"][1], SupportedProjectTypes["api"][0]),
			huh.NewOption(SupportedProjectTypes["cli"][1], SupportedProjectTypes["cli"][0]),
			huh.NewOption(SupportedProjectTypes["grpc"][1], SupportedProjectTypes["grpc"][0]),
		).
		Value(&cfg.ProjectType)
}

// GoFrameworkSelect runs the Go framework select.
func GoFrameworkSelect(cfg *config.Config) *huh.Select[string] {
	return huh.NewSelect[string]().
		Title(FormGoFrameworkTitle).
		Description(FormGoFrameworkDescription).
		Options(
			huh.NewOption(SupportedGoFrameworks["gin"][1], SupportedGoFrameworks["gin"][0]),
			huh.NewOption(SupportedGoFrameworks["fiber"][1], SupportedGoFrameworks["fiber"][0]),
		).
		Value(&cfg.GoFramework)
}
