package ui

const (
	// FormPromptSignature represents the prompt signature for the form.
	FormPromptSignature = "> "

	// FormGoModuleNameTitle represents the title for the Go module name input.
	FormGoModuleNameTitle = "What's your Go module name in the go.mod file?\n"
	// FormGoModuleNameDescription represents the description for the Go module name input.
	FormGoModuleNameDescription = "Option can be any name of your Go module (e.g., `github.com/user/project`).\n"

	// FormProjectNameTitle represents the title for the frontend project name input.
	FormProjectNameTitle = "What's your project name?\n"
	// FormProjectNameDescription represents the description for the project name input.
	FormProjectNameDescription = "The name of the new workspace and initial project. (e.g., `project`).\n"

	// FormGoFrameworkTitle represents the title for the Go framework select.
	FormGoFrameworkTitle = "Select the Go web framework or router\n"
	// FormGoFrameworkDescription represents the description for the Go framework select.
	FormGoFrameworkDescription = "This framework (or router) will be used to build\nthe backend part of your application.\n"

	// FormProjectTypeTitle represents the title for the reactivity library select.
	FormProjectTypeTitle = "Select the project type\n"
	// FormProjectTypeDescription represents the description for the reactivity library select.
	FormProjectTypeDescription = "The type of your application.\n"

	// FormGoReleaserUsageTitle represents the title for the GoReleaser switch.
	FormGoReleaserUsageTitle = "Use the GoReleaser to deliver your Go binaries?\n"
	// FormGoReleaserUsageDescription represents the description for the GoReleaser switch.
	FormGoReleaserUsageDescription = "This tool will be used to deliver your Go binaries.\n\nFor more info → https://github.com/goreleaser/goreleaser"

	// FormGSemverUsageTitle represents the title for the GSemver switch.
	FormGSemverUsageTitle = "Use the GSemver to generate your next semver version?\n"
	// FormGSemverUsageDescription represents the description for the GSemver switch.
	FormGSemverUsageDescription = "This tool will generate the next semver version using the git commit convention.\n\nFor more info → https://github.com/arnaud-deprez/gsemver"

	// FormGolangCILintUsageTitle represents the title for the Golang CI Lint switch.
	FormGolangCILintUsageTitle = "Use the Golang CI Lint to lint your Go code?\n"
	// FormGolangCILintUsageDescription represents the description for the Golang CI Lint switch.
	FormGolangCILintUsageDescription = "This tool will be used to lint your Go code.\n\nFor more info → https://github.com/golangci/golangci-lint"
)

const (
	// CreateSpinnerTitle represents the title for the command create spinner.
	CreateSpinnerTitle = " Jaguar CLI is creating your project. Please wait..."
	// ErrorSummaryTitle represents the title of the unknown summary.
	ErrorSummaryTitle = "✕ Oops... Something went wrong!\n"
	// CreateSummaryTitle represents the title of the project summary.
	CreateSummaryTitle = "✓ Your project has been created successfully!\n"

	// CreateSummaryHeadingBackend represents the heading of the backend summary.
	CreateSummaryHeadingBackend = "Backend ↘"
	// CreateSummaryHeadingTools represents the heading of the tools summary.
	CreateSummaryHeadingTools = "Tools ↘"

	// MoreInformationTitle represents the title of the more information string.
	MoreInformationTitle string = "\n✱ For more information go to the official docs: https://github.com/shipengqi/jaguar \n"
)
