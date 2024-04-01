package ui

import (
	"github.com/charmbracelet/huh"

	"github.com/shipengqi/jaguar/internal/actions/create/config"
)

const (
	DefaultGoModuleMinLength    = 2
	DefaultGoModuleMaxLength    = 128
	DefaultProjectNameMinLength = 2
	DefaultProjectNameMaxLength = 20
)

// GoModuleNameInput runs the go module name input.
func GoModuleNameInput(cfg *config.Config) *huh.Input {
	return huh.NewInput().
		Title(FormGoModuleNameTitle).
		Description(FormGoModuleNameDescription).
		Prompt(FormPromptSignature).
		Validate(StringValidator("Go module name", DefaultGoModuleMinLength, DefaultGoModuleMaxLength)).
		Value(&cfg.ModuleName)
}

// ProjectNameInput runs the project name input.
func ProjectNameInput(cfg *config.Config) *huh.Input {
	return huh.NewInput().
		Title(FormProjectNameTitle).
		Description(FormProjectNameDescription).
		Prompt(FormPromptSignature).
		Validate(StringValidator("project name", DefaultProjectNameMinLength, DefaultProjectNameMaxLength)).
		Value(&cfg.ProjectName)
}
