package ui

import (
	"fmt"
	"strconv"
)

// SummaryConfig holds the data needed to display the summary.
type SummaryConfig struct {
	ProjectName       string
	ProjectType       string
	Language          string
	Framework         string
	FrontendFramework string
	ModuleName        string
	UseGolangCILint   bool
	UseGoReleaser     bool
	UseGSemver        bool
	UseGithubActions  bool
}

func ShowSummary(cfg *SummaryConfig) {
	frameworkStr := cfg.Framework
	if frameworkStr == "" {
		frameworkStr = "N/A"
	}

	body := fmt.Sprintf(
		"%s\n\nProject name:  %s\nLanguage:      %s\nProject type:  %s\nFramework:     %s",
		MakeStyledString("Application ↘", &StringStyle{Color: ColorGray}),
		MakeStyledString(cfg.ProjectName, &StringStyle{Color: ColorBlue}),
		MakeStyledString(cfg.Language, &StringStyle{Color: ColorBlue}),
		MakeStyledString(cfg.ProjectType, &StringStyle{Color: ColorBlue}),
		MakeStyledString(frameworkStr, &StringStyle{Color: ColorBlue}),
	)

	if cfg.FrontendFramework != "" {
		body += fmt.Sprintf("\nFrontend:      %s",
			MakeStyledString(cfg.FrontendFramework, &StringStyle{Color: ColorBlue}))
	}
	if cfg.ModuleName != "" {
		body += fmt.Sprintf("\nGo module:     %s",
			MakeStyledString(cfg.ModuleName, &StringStyle{Color: ColorBlue}))
	}

	if cfg.Language == "go" {
		body += fmt.Sprintf(
			"\n\n%s\n\ngolangci-lint: %s\nGoReleaser:    %s\nGSemver:       %s\nGitHub Actions:%s",
			MakeStyledString("Tools ↘", &StringStyle{Color: ColorGray}),
			MakeStyledString(strconv.FormatBool(cfg.UseGolangCILint), &StringStyle{Color: ColorBlue}),
			MakeStyledString(strconv.FormatBool(cfg.UseGoReleaser), &StringStyle{Color: ColorBlue}),
			MakeStyledString(strconv.FormatBool(cfg.UseGSemver), &StringStyle{Color: ColorBlue}),
			MakeStyledString(strconv.FormatBool(cfg.UseGithubActions), &StringStyle{Color: ColorBlue}),
		)
	}

	fmt.Println(MakeStyledString(
		"✓ Your project has been created successfully!\n",
		&StringStyle{Color: ColorGreen, IsBold: true},
	))
	fmt.Println(MakeStyledFrame(body, &FrameStyle{Padding: []int{1}, Color: ColorGreen}))
	fmt.Println(MakeStyledString(
		"\n✱ For more information go to the official docs: https://github.com/shipengqi/jaguar \n",
		&StringStyle{Color: ColorYellow},
	))
}
