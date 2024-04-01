package ui

import (
	"fmt"
	"strconv"

	"github.com/shipengqi/jaguar/internal/actions/create/config"
	"github.com/shipengqi/jaguar/internal/actions/create/helpers"
)

func ShowSummary(cfg *config.Config) {
	// Generate content body.
	contentBody := fmt.Sprintf(
		"%s\n\nApplication type: %s\nGo web framework: %s\n\n%s\n\nIs use golangci-lint to lint your Go code? %s\nIs use GoReleaser to deliver your Go binaries? %s\nIs use GSemver to generate your next semver version? %s",
		helpers.MakeStyledString(CreateSummaryHeadingApp, &helpers.StringStyle{Color: ColorGray}),
		helpers.MakeStyledString(SupportedProjectTypes[cfg.ProjectType][1], &helpers.StringStyle{Color: ColorBlue}),
		helpers.MakeStyledString(SupportedGoFrameworks[cfg.GoFramework][1], &helpers.StringStyle{Color: ColorBlue}),
		helpers.MakeStyledString(CreateSummaryHeadingTools, &helpers.StringStyle{Color: ColorGray}),
		helpers.MakeStyledString(strconv.FormatBool(cfg.IsUseGolangCILint), &helpers.StringStyle{Color: ColorBlue}),
		helpers.MakeStyledString(strconv.FormatBool(cfg.IsUseGoReleaser), &helpers.StringStyle{Color: ColorBlue}),
		helpers.MakeStyledString(strconv.FormatBool(cfg.IsUseGSemver), &helpers.StringStyle{Color: ColorBlue}),
	)

	// Show created project info.
	fmt.Println(helpers.MakeStyledString(
		CreateSummaryTitle,
		&helpers.StringStyle{Color: ColorGreen, IsBold: true},
	))
	fmt.Println(helpers.MakeStyledFrame(
		contentBody,
		&helpers.FrameStyle{Padding: []int{1}, Color: ColorGreen},
	))
	fmt.Println(helpers.MakeStyledString(
		MoreInformationTitle,
		&helpers.StringStyle{Color: ColorYellow},
	))
}
