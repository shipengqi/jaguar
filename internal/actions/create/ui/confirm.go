package ui

import (
	"github.com/charmbracelet/huh"

	"github.com/shipengqi/jaguar/internal/actions/create/config"
)

// IsUseGolangCILintConfirm runs the confirmation if Golang CI Lint is used.
func IsUseGolangCILintConfirm(cfg *config.Config) *huh.Confirm {
	return huh.NewConfirm().
		Title(FormGolangCILintUsageTitle).
		Description(FormGolangCILintUsageDescription).
		Affirmative("Yes").
		Negative("No").
		Value(&cfg.IsUseGolangCILint)
}

// IsUseGoReleaserConfirm runs the confirmation if GoReleaser is used.
func IsUseGoReleaserConfirm(cfg *config.Config) *huh.Confirm {
	return huh.NewConfirm().
		Title(FormGoReleaserUsageTitle).
		Description(FormGoReleaserUsageDescription).
		Affirmative("Yes").
		Negative("No").
		Value(&cfg.IsUseGoReleaser)
}

// IsUseGSemverConfirm runs the confirmation if GSemver is used.
func IsUseGSemverConfirm(cfg *config.Config) *huh.Confirm {
	return huh.NewConfirm().
		Title(FormGSemverUsageTitle).
		Description(FormGSemverUsageDescription).
		Affirmative("Yes").
		Negative("No").
		Value(&cfg.IsUseGSemver)
}
