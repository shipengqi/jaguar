package ui

import (
	"fmt"

	"charm.land/huh/v2"
)

const (
	FormPromptSignature = "> "

	FormProjectNameTitle       = "What's your project name?\n"
	FormProjectNameDescription = "The name of the new project (e.g., `my-app`).\n"

	FormLanguageTitle       = "Select the language\n"
	FormLanguageDescription = "The programming language for your project.\n"

	FormGoTypeTitle       = "Select the Go project type\n"
	FormGoTypeDescription = "The type of your Go application.\n"

	FormNodeTypeTitle       = "Select the Node.js project type\n"
	FormNodeTypeDescription = "The type of your Node.js application.\n"

	FormFrameworkTitle       = "Select the web framework\n"
	FormFrameworkDescription = "The web framework for your application.\n"

	FormFrontendTitle       = "Select the frontend framework\n"
	FormFrontendDescription = "The frontend framework to embed into your binary.\n"

	FormFrontendTypeTitle       = "Select the frontend framework\n"
	FormFrontendTypeDescription = "The standalone frontend framework to scaffold.\n"

	FormGoModuleTitle       = "What's your Go module name?\n"
	FormGoModuleDescription = "Option can be any name of your Go module (e.g., `github.com/user/project`).\n"

	FormOutputDirTitle       = "Where should the project be generated?\n"
	FormOutputDirDescription = "Directory to create the project in (default: `output`).\n"

	FormGolangCILintTitle       = "Use golangci-lint to lint your Go code?\n"
	FormGolangCILintDescription = "For more info → https://github.com/golangci/golangci-lint"

	FormGoReleaserTitle       = "Use GoReleaser to deliver your Go binaries?\n"
	FormGoReleaserDescription = "For more info → https://github.com/goreleaser/goreleaser"

	FormGSemverTitle       = "Use GSemver to generate your next semver version?\n"
	FormGSemverDescription = "For more info → https://github.com/arnaud-deprez/gsemver"

	FormGithubActionsTitle       = "Add GitHub Actions CI workflows?\n"
	FormGithubActionsDescription = "Automate your build, test, and deployment pipeline.\n\nFor more info → https://docs.github.com/en/actions"

	FormDockerfileTitle       = "Add a Dockerfile?\n"
	FormDockerfileDescription = "Containerize your application.\n"

	FormDockerComposeTitle       = "Add a docker-compose.yml?\n"
	FormDockerComposeDescription = "Orchestrate your local development environment.\n"

	FormNodeCITitle       = "Add GitHub Actions CI (ESLint + test)?\n"
	FormNodeCIDescription = "Automate linting and testing for your Node.js project.\n"

	FormPythonCITitle       = "Add GitHub Actions CI (pytest + ruff lint)?\n"
	FormPythonCIDescription = "Automate linting and testing for your Python project.\n"
)

func ProjectNameInput(name *string) *huh.Input {
	return huh.NewInput().
		Title(FormProjectNameTitle).
		Description(FormProjectNameDescription).
		Prompt(FormPromptSignature).
		Validate(func(s string) error {
			if len(s) < 2 {
				return fmt.Errorf("project name must be at least 2 characters")
			}
			if len(s) > 64 {
				return fmt.Errorf("project name must be at most 64 characters")
			}
			return nil
		}).
		Value(name)
}

func LanguageSelect(lang *string) *huh.Select[string] {
	return huh.NewSelect[string]().
		Title(FormLanguageTitle).
		Description(FormLanguageDescription).
		Options(
			huh.NewOption("Go", "go"),
			huh.NewOption("Node.js", "nodejs"),
			huh.NewOption("Python", "python"),
			huh.NewOption("Frontend", "frontend"),
		).
		Value(lang)
}

func GoTypeSelect(projectType *string) *huh.Select[string] {
	return huh.NewSelect[string]().
		Title(FormGoTypeTitle).
		Description(FormGoTypeDescription).
		Options(
			huh.NewOption("go-api  — REST API (gin/fiber)", "go-api"),
			huh.NewOption("go-embed — Go + embedded frontend", "go-embed"),
			huh.NewOption("go-cli  — CLI tool", "go-cli"),
			huh.NewOption("go-grpc — gRPC service", "go-grpc"),
		).
		Value(projectType)
}

func NodeTypeSelect(projectType *string) *huh.Select[string] {
	return huh.NewSelect[string]().
		Title(FormNodeTypeTitle).
		Description(FormNodeTypeDescription).
		Options(
			huh.NewOption("koa    — Koa.js framework", "nodejs"),
			huh.NewOption("nestjs — NestJS framework", "nodejs-nestjs"),
		).
		Value(projectType)
}

func GoFrameworkSelect(framework *string) *huh.Select[string] {
	return huh.NewSelect[string]().
		Title(FormFrameworkTitle).
		Description(FormFrameworkDescription).
		Options(
			huh.NewOption("Gin", "gin"),
			huh.NewOption("Fiber", "fiber"),
		).
		Value(framework)
}

func NodeFrameworkSelect(framework *string) *huh.Select[string] {
	return huh.NewSelect[string]().
		Title(FormFrameworkTitle).
		Description(FormFrameworkDescription).
		Options(
			huh.NewOption("Koa", "koa"),
			huh.NewOption("NestJS", "nestjs"),
			huh.NewOption("Next.js", "nextjs"),
		).
		Value(framework)
}

func FrontendFrameworkSelect(frontend *string) *huh.Select[string] {
	return huh.NewSelect[string]().
		Title(FormFrontendTitle).
		Description(FormFrontendDescription).
		Options(
			huh.NewOption("React", "react"),
			huh.NewOption("Vue", "vue"),
			huh.NewOption("Angular", "angular"),
		).
		Value(frontend)
}

func FrontendTypeSelect(projectType *string) *huh.Select[string] {
	return huh.NewSelect[string]().
		Title(FormFrontendTypeTitle).
		Description(FormFrontendTypeDescription).
		Options(
			huh.NewOption("React  — Next.js + shadcn/ui + Zustand + TanStack Query", "frontend-react"),
			huh.NewOption("Vue    — Vite + shadcn-vue + Pinia + TanStack Query", "frontend-vue"),
			huh.NewOption("Angular — Vite + spartan/ui + NgRx Signals + Transloco", "frontend-angular"),
		).
		Value(projectType)
}

func GoModuleInput(module *string) *huh.Input {
	return huh.NewInput().
		Title(FormGoModuleTitle).
		Description(FormGoModuleDescription).
		Prompt(FormPromptSignature).
		Validate(func(s string) error {
			if len(s) < 2 {
				return fmt.Errorf("module name must be at least 2 characters")
			}
			return nil
		}).
		Value(module)
}

func OutputDirInput(dir *string) *huh.Input {
	return huh.NewInput().
		Title(FormOutputDirTitle).
		Description(FormOutputDirDescription).
		Prompt(FormPromptSignature).
		Placeholder("output").
		Validate(func(s string) error {
			if s == "" {
				return fmt.Errorf("output directory cannot be empty")
			}
			return nil
		}).
		Value(dir)
}

func GolangCILintConfirm(v *bool) *huh.Confirm {
	return huh.NewConfirm().
		Title(FormGolangCILintTitle).
		Description(FormGolangCILintDescription).
		Affirmative("Yes").Negative("No").
		Value(v)
}

func GoReleaserConfirm(v *bool) *huh.Confirm {
	return huh.NewConfirm().
		Title(FormGoReleaserTitle).
		Description(FormGoReleaserDescription).
		Affirmative("Yes").Negative("No").
		Value(v)
}

func GSemverConfirm(v *bool) *huh.Confirm {
	return huh.NewConfirm().
		Title(FormGSemverTitle).
		Description(FormGSemverDescription).
		Affirmative("Yes").Negative("No").
		Value(v)
}

func GithubActionsConfirm(v *bool) *huh.Confirm {
	return huh.NewConfirm().
		Title(FormGithubActionsTitle).
		Description(FormGithubActionsDescription).
		Affirmative("Yes").Negative("No").
		Value(v)
}

func DockerfileConfirm(v *bool) *huh.Confirm {
	return huh.NewConfirm().
		Title(FormDockerfileTitle).
		Description(FormDockerfileDescription).
		Affirmative("Yes").Negative("No").
		Value(v)
}

func DockerComposeConfirm(v *bool) *huh.Confirm {
	return huh.NewConfirm().
		Title(FormDockerComposeTitle).
		Description(FormDockerComposeDescription).
		Affirmative("Yes").Negative("No").
		Value(v)
}
