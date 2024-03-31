package ui

import (
	"errors"
	"github.com/shipengqi/jaguar/internal/actions/create/config"

	"github.com/charmbracelet/huh"
)

// GoModuleNameInput runs the go module name input.
func GoModuleNameInput(cfg *config.Config) *huh.Input {
	return huh.NewInput().
		Title(FormGoModuleNameTitle).
		Description(FormGoModuleNameDescription).
		Prompt(FormPromptSignature).
		Validate(func(s string) error {
			if s == "" {
				return errors.New("enter correct Go module name")
			}
			return nil
		}).
		Value(&cfg.ModuleName)
}

// ProjectNameInput runs the project name input.
func ProjectNameInput(cfg *config.Config) *huh.Input {
	return huh.NewInput().
		Title(FormProjectNameTitle).
		Description(FormProjectNameDescription).
		Prompt(FormPromptSignature).
		Validate(func(s string) error {
			if s == "" {
				return errors.New("enter correct project name")
			}
			return nil
		}).
		Value(&cfg.ProjectName)
}
