package create

import (
	"charm.land/huh/v2"

	"github.com/shipengqi/jaguar/internal/create/ui"
)

func prerun(cfg *Config) error {
	// Step 1: project name (if not provided via args)
	if cfg.ProjectName == "" {
		if err := huh.NewForm(huh.NewGroup(ui.ProjectNameInput(&cfg.ProjectName))).Run(); err != nil {
			return err
		}
	}

	// Step 2: output directory (if not provided via flag)
	if !cfg.Changed(FlagOutputDir) {
		if err := huh.NewForm(huh.NewGroup(ui.OutputDirInput(&cfg.OutputDir))).Run(); err != nil {
			return err
		}
		if cfg.OutputDir == "" {
			cfg.OutputDir = "output"
		}
	}

	// Step 3: language / type selection
	if err := runTypeSelection(cfg); err != nil {
		return err
	}

	// Step 4: framework (go-api, go-embed, nodejs)
	if err := runFrameworkSelection(cfg); err != nil {
		return err
	}

	// Step 5: frontend framework (go-embed only)
	if cfg.ProjectType == ProjectTypeGoEmbed && !cfg.Changed(FlagFrontendFramework) {
		if err := huh.NewForm(huh.NewGroup(ui.FrontendFrameworkSelect(&cfg.FrontendFramework))).Run(); err != nil {
			return err
		}
	}

	// Step 6: Go module name (Go types only)
	if cfg.Language == LanguageGo && cfg.ModuleName == "" {
		if err := huh.NewForm(huh.NewGroup(ui.GoModuleInput(&cfg.ModuleName))).Run(); err != nil {
			return err
		}
	}

	// Step 7: toolchain options
	return runToolchainSelection(cfg)
}

func runTypeSelection(cfg *Config) error {
	// If type already fully specified via flag, infer language and return.
	if cfg.ProjectType != "" && isKnownType(cfg.ProjectType) {
		cfg.Language = languageForType(cfg.ProjectType)
		return nil
	}

	// Ask language first.
	if cfg.Language == "" {
		if err := huh.NewForm(huh.NewGroup(ui.LanguageSelect(&cfg.Language))).Run(); err != nil {
			return err
		}
	}

	// Then ask type within language.
	switch cfg.Language {
	case LanguageGo:
		if cfg.ProjectType == "" || !isGoType(cfg.ProjectType) {
			if err := huh.NewForm(huh.NewGroup(ui.GoTypeSelect(&cfg.ProjectType))).Run(); err != nil {
				return err
			}
		}
	case LanguageNodeJS:
		if cfg.ProjectType == "" {
			cfg.ProjectType = ProjectTypeNodeJS
		}
	case LanguagePython:
		cfg.ProjectType = ProjectTypePython
	case LanguageFrontend:
		if cfg.ProjectType == "" || !isFrontendType(cfg.ProjectType) {
			if err := huh.NewForm(huh.NewGroup(ui.FrontendTypeSelect(&cfg.ProjectType))).Run(); err != nil {
				return err
			}
		}
	}
	return nil
}

func runFrameworkSelection(cfg *Config) error {
	if cfg.Changed(FlagFramework) {
		return nil
	}
	switch cfg.ProjectType {
	case ProjectTypeGoAPI, ProjectTypeGoEmbed:
		if cfg.Framework == "" || (cfg.Framework != FrameworkGin && cfg.Framework != FrameworkFiber) {
			return huh.NewForm(huh.NewGroup(ui.GoFrameworkSelect(&cfg.Framework))).Run()
		}
	case ProjectTypeGoCLI, ProjectTypeGoGRPC:
		cfg.Framework = ""
	case ProjectTypeNodeJS:
		if cfg.Framework == "" || (cfg.Framework != FrameworkKoa && cfg.Framework != FrameworkNestJS) {
			return huh.NewForm(huh.NewGroup(ui.NodeFrameworkSelect(&cfg.Framework))).Run()
		}
	case ProjectTypePython:
		cfg.Framework = FrameworkFastAPI
	case ProjectTypeFrontendReact, ProjectTypeFrontendVue, ProjectTypeFrontendAngular:
		cfg.Framework = ""
	}
	return nil
}

func runToolchainSelection(cfg *Config) error {
	var groups []*huh.Group
	switch cfg.Language {
	case LanguageGo:
		if !cfg.Changed(FlagUseGolangCILint) {
			groups = append(groups, huh.NewGroup(ui.GolangCILintConfirm(&cfg.UseGolangCILint)))
		}
		if !cfg.Changed(FlagUseGoReleaser) {
			groups = append(groups, huh.NewGroup(ui.GoReleaserConfirm(&cfg.UseGoReleaser)))
		}
		if !cfg.Changed(FlagUseGSemver) {
			groups = append(groups, huh.NewGroup(ui.GSemverConfirm(&cfg.UseGSemver)))
		}
		if !cfg.Changed(FlagUseGithubActions) {
			groups = append(groups, huh.NewGroup(ui.GithubActionsConfirm(&cfg.UseGithubActions)))
		}
	case LanguageNodeJS:
		if !cfg.Changed(FlagUseGithubActions) {
			groups = append(groups, huh.NewGroup(ui.GithubActionsConfirm(&cfg.UseGithubActions)))
		}
	case LanguagePython:
		if !cfg.Changed(FlagUseGithubActions) {
			groups = append(groups, huh.NewGroup(ui.GithubActionsConfirm(&cfg.UseGithubActions)))
		}
	case LanguageFrontend:
		if !cfg.Changed(FlagUseGithubActions) {
			groups = append(groups, huh.NewGroup(ui.GithubActionsConfirm(&cfg.UseGithubActions)))
		}
	}
	if len(groups) == 0 {
		return nil
	}
	return huh.NewForm(groups...).Run()
}

func isKnownType(t string) bool {
	switch t {
	case ProjectTypeGoAPI, ProjectTypeGoEmbed, ProjectTypeGoCLI, ProjectTypeGoGRPC,
		ProjectTypeNodeJS, ProjectTypePython,
		ProjectTypeFrontendReact, ProjectTypeFrontendVue, ProjectTypeFrontendAngular:
		return true
	}
	return false
}

func isGoType(t string) bool {
	switch t {
	case ProjectTypeGoAPI, ProjectTypeGoEmbed, ProjectTypeGoCLI, ProjectTypeGoGRPC:
		return true
	}
	return false
}

func languageForType(t string) string {
	switch t {
	case ProjectTypeGoAPI, ProjectTypeGoEmbed, ProjectTypeGoCLI, ProjectTypeGoGRPC:
		return LanguageGo
	case ProjectTypeNodeJS:
		return LanguageNodeJS
	case ProjectTypePython:
		return LanguagePython
	case ProjectTypeFrontendReact, ProjectTypeFrontendVue, ProjectTypeFrontendAngular:
		return LanguageFrontend
	}
	return ""
}

func isFrontendType(t string) bool {
	switch t {
	case ProjectTypeFrontendReact, ProjectTypeFrontendVue, ProjectTypeFrontendAngular:
		return true
	}
	return false
}
