package create

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/pflag"
)

const (
	FlagProjectType       = "type"
	FlagFramework         = "framework"
	FlagFrontendFramework = "frontend-framework"
	FlagModuleName        = "module"
	FlagOutputDir         = "output-dir"
	FlagUseGolangCILint   = "use-golangci-lint"
	FlagUseGoReleaser     = "use-goreleaser"
	FlagUseGSemver        = "use-gsemver"
	FlagUseGithubActions  = "use-github-actions"
)

type Config struct {
	fs *pflag.FlagSet

	ProjectName       string
	ProjectType       string
	Language          string
	Framework         string
	FrontendFramework string
	ModuleName        string
	OutputDir         string

	UseGolangCILint  bool
	UseGoReleaser    bool
	UseGSemver       bool
	UseGithubActions bool

	SkeletonVersion string
}

func NewConfig() *Config {
	return &Config{
		UseGolangCILint:  true,
		UseGoReleaser:    true,
		UseGSemver:       true,
		UseGithubActions: true,
		SkeletonVersion:  SkeletonVersion1,
		Framework:        FrameworkGin,
		OutputDir:        "output",
	}
}

func (c *Config) AddFlags(fs *pflag.FlagSet) {
	c.fs = fs
	fs.StringVarP(&c.ProjectType, FlagProjectType, "t", c.ProjectType, "the type of your application (go-api, go-embed, go-cli, go-grpc, nodejs, python)")
	fs.StringVar(&c.Framework, FlagFramework, c.Framework, "the web framework (gin, fiber, koa, nestjs, fastapi)")
	fs.StringVar(&c.FrontendFramework, FlagFrontendFramework, c.FrontendFramework, "the frontend framework (react, vue, angular)")
	fs.StringVarP(&c.ModuleName, FlagModuleName, "m", c.ModuleName, "the Go module name in the go.mod file")
	fs.StringVarP(&c.OutputDir, FlagOutputDir, "o", c.OutputDir, "the directory where the project will be generated")
	fs.BoolVar(&c.UseGolangCILint, FlagUseGolangCILint, c.UseGolangCILint, "use golangci-lint to lint your Go code")
	fs.BoolVar(&c.UseGoReleaser, FlagUseGoReleaser, c.UseGoReleaser, "use GoReleaser to deliver your Go binaries")
	fs.BoolVar(&c.UseGSemver, FlagUseGSemver, c.UseGSemver, "use GSemver to generate your next semver version")
	fs.BoolVar(&c.UseGithubActions, FlagUseGithubActions, c.UseGithubActions, "add GitHub Actions CI workflows")
}

func (c *Config) Changed(name string) bool {
	if c.fs == nil {
		return false
	}
	return c.fs.Changed(name)
}

func (c *Config) ExportTemplateData() *TemplateData {
	bin := "apiserver"
	switch c.ProjectType {
	case ProjectTypeGoCLI:
		bin = "examplecli"
	case ProjectTypeGoGRPC:
		bin = "rpcserver"
	}
	return &TemplateData{
		App: AppData{
			Name:           c.ProjectName,
			Type:           c.ProjectType,
			Language:       c.Language,
			Framework:      c.Framework,
			ModuleName:     c.ModuleName,
			Logo:           normalizeAppLogo(c.ProjectName),
			EnvPrefix:      normalizeAppEnv(c.ProjectName),
			NormalizedName: normalizeAppName(c.ProjectName),
			DocumentLink:   fmt.Sprintf("https://%s", c.ModuleName),
		},
		Build: BuildData{
			Bin:  bin,
			Root: fmt.Sprintf("%s/cmd/%s", c.ModuleName, bin),
		},
		Frontend: FrontendData{
			Enabled:   c.FrontendFramework != "",
			Framework: c.FrontendFramework,
		},
	}
}

func normalizeAppName(value string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	value = reg.ReplaceAllString(value, "")
	return strings.ToLower(value)
}

func normalizeAppEnv(value string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	value = reg.ReplaceAllString(value, "_")
	value = strings.Trim(value, "_")
	return strings.ToUpper(value)
}

func normalizeAppLogo(name string) string {
	re, _ := regexp.Compile(`[^a-z\-_?.]+`)
	name = re.ReplaceAllString(strings.ToLower(name), " ")
	// produce a simple ASCII banner via repeated chars
	line := strings.Repeat("=", len(name)+4)
	logo := fmt.Sprintf("%s\n  %s  \n%s", line, strings.ToUpper(name), line)
	return strings.ReplaceAll(logo, "`", "` + \"`\" + `")
}
