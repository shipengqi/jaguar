package create

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/shipengqi/jaguar/internal/create/ui"
	"github.com/shipengqi/jaguar/internal/fsutil"
	"github.com/shipengqi/jaguar/skeletons"
)

type stages struct {
	cfg      *Config
	skeleton skeletons.Skeleton
}

func newStages(cfg *Config) *stages {
	return &stages{cfg: cfg}
}

// outputPath returns the destination path for the project or a sub-path inside it.
func (s *stages) outputPath(parts ...string) string {
	base := filepath.Join(s.cfg.OutputDir, s.cfg.ProjectName)
	if len(parts) == 0 {
		return base
	}
	return filepath.Join(append([]string{base}, parts...)...)
}

func (s *stages) initialize() error {
	ss := skeletons.New()
	switch s.cfg.SkeletonVersion {
	case SkeletonVersion1:
		s.skeleton = ss.V1
	default:
		return fmt.Errorf("unknown skeleton version '%s'", s.cfg.SkeletonVersion)
	}
	return nil
}

func (s *stages) run(cancel context.CancelFunc) error {
	defer cancel()

	slog.Debug("initializing skeletons")
	if err := s.initialize(); err != nil {
		return err
	}

	slog.Debug("removing existing project directory", "name", s.outputPath())
	if err := os.RemoveAll(s.outputPath()); err != nil {
		return err
	}

	data := s.cfg.ExportTemplateData()
	slog.Debug("exporting template data", "type", s.cfg.ProjectType, "framework", s.cfg.Framework)

	if err := s.copyProjectFiles(data); err != nil {
		return err
	}

	if err := s.copyProjectToolchain(data); err != nil {
		return err
	}

	return nil
}

func (s *stages) copyProjectFiles(data *TemplateData) error {
	ver := s.cfg.SkeletonVersion
	dst := s.outputPath()

	switch s.cfg.ProjectType {
	case ProjectTypeGoAPI:
		src := fmt.Sprintf("%s/go-api/%s", ver, s.cfg.Framework)
		return fsutil.CopyAndCompleteFiles(s.skeleton.GoAPI, src, dst, data)

	case ProjectTypeGoEmbed:
		src := fmt.Sprintf("%s/go-embed/%s", ver, s.cfg.Framework)
		if err := fsutil.CopyAndCompleteFiles(s.skeleton.GoEmbed, src, dst, data); err != nil {
			return err
		}
		webDst := filepath.Join(dst, "web")
		fwSrc := fmt.Sprintf("%s/frontend/%s", ver, s.cfg.FrontendFramework)
		return fsutil.CopyAndCompleteFiles(s.skeleton.Frontend, fwSrc, webDst, data)

	case ProjectTypeFrontendReact:
		src := fmt.Sprintf("%s/frontend/react", ver)
		return fsutil.CopyAndCompleteFiles(s.skeleton.Frontend, src, dst, data)

	case ProjectTypeFrontendVue:
		src := fmt.Sprintf("%s/frontend/vue", ver)
		return fsutil.CopyAndCompleteFiles(s.skeleton.Frontend, src, dst, data)

	case ProjectTypeFrontendAngular:
		src := fmt.Sprintf("%s/frontend/angular", ver)
		return fsutil.CopyAndCompleteFiles(s.skeleton.Frontend, src, dst, data)

	case ProjectTypeGoCLI:
		src := fmt.Sprintf("%s/go-cli", ver)
		return fsutil.CopyAndCompleteFiles(s.skeleton.GoCLI, src, dst, data)

	case ProjectTypeGoGRPC:
		src := fmt.Sprintf("%s/go-grpc", ver)
		return fsutil.CopyAndCompleteFiles(s.skeleton.GoGRPC, src, dst, data)

	case ProjectTypeNodeJS:
		src := fmt.Sprintf("%s/nodejs/%s", ver, s.cfg.Framework)
		return fsutil.CopyAndCompleteFiles(s.skeleton.NodeJS, src, dst, data)

	case ProjectTypePython:
		src := fmt.Sprintf("%s/python/fastapi", ver)
		return fsutil.CopyAndCompleteFiles(s.skeleton.Python, src, dst, data)

	default:
		return fmt.Errorf("unknown project type '%s'", s.cfg.ProjectType)
	}
}

func (s *stages) copyProjectToolchain(data *TemplateData) error {
	switch s.cfg.Language {
	case LanguageGo:
		return s.copyGoToolchain(data)
	case LanguageNodeJS, LanguagePython:
		return s.copyNonGoToolchain(data)
	case LanguageFrontend:
		return s.copyFrontendToolchain(data)
	}
	return nil
}

func (s *stages) copyNonGoToolchain(data *TemplateData) error {
	ver := s.cfg.SkeletonVersion
	pf := s.skeleton.ProjectFiles

	// shared .github files (issue templates, dependabot, labeler, etc.)
	actionsdir := s.outputPath(".github")
	slog.Debug("generating github files", "dst", actionsdir)
	if err := fsutil.CopyAndCompleteFiles(pf,
		fmt.Sprintf("%s/projectfiles/.github", ver),
		actionsdir, data); err != nil {
		return err
	}

	// remove Go-specific workflows that were copied from shared .github
	for _, f := range []string{"lint.yaml", "release.yaml"} {
		_ = os.Remove(s.outputPath(".github", "workflows", f))
	}

	// language-specific CI workflow
	var langDir string
	switch s.cfg.Language {
	case LanguageNodeJS:
		langDir = "nodejs"
	case LanguagePython:
		langDir = "python"
	}
	if langDir != "" {
		src := fmt.Sprintf("%s/projectfiles/%s/.github", ver, langDir)
		if err := fsutil.CopyAndCompleteFiles(pf, src, actionsdir, data); err != nil {
			return err
		}
	}

	if !s.cfg.UseGithubActions {
		workflowdir := s.outputPath(".github", "workflows")
		slog.Debug("github actions disabled, removing workflows", "dst", workflowdir)
		if err := os.RemoveAll(workflowdir); err != nil {
			return err
		}
	}

	return nil
}

func (s *stages) copyGoToolchain(data *TemplateData) error {
	ver := s.cfg.SkeletonVersion
	pf := s.skeleton.ProjectFiles

	if s.cfg.UseGolangCILint {
		dst := s.outputPath(".golangci.yaml")
		slog.Debug("generating golangci config", "dst", dst)
		if err := fsutil.CopyAndCompleteFile(pf,
			fmt.Sprintf("%s/projectfiles/.golangci.yaml.tpl", ver),
			dst, data); err != nil {
			return err
		}
	}

	if s.cfg.UseGoReleaser {
		dst := s.outputPath(".goreleaser.yaml")
		slog.Debug("generating goreleaser config", "dst", dst)
		if err := fsutil.CopyAndCompleteFile(pf,
			fmt.Sprintf("%s/projectfiles/.goreleaser.yaml.tpl", ver),
			dst, data); err != nil {
			return err
		}
	}

	if s.cfg.UseGSemver {
		dst := s.outputPath(".gsemver.yaml")
		slog.Debug("generating gsemver config", "dst", dst)
		if err := fsutil.CopyAndCompleteFile(pf,
			fmt.Sprintf("%s/projectfiles/.gsemver.yaml", ver),
			dst, data); err != nil {
			return err
		}
	}

	actionsdir := s.outputPath(".github")
	slog.Debug("generating github files", "dst", actionsdir)
	if err := fsutil.CopyAndCompleteFiles(pf,
		fmt.Sprintf("%s/projectfiles/.github", ver),
		actionsdir, data); err != nil {
		return err
	}

	if !s.cfg.UseGithubActions {
		workflowdir := s.outputPath(".github", "workflows")
		slog.Debug("github actions disabled, removing workflows", "dst", workflowdir)
		if err := os.RemoveAll(workflowdir); err != nil {
			return err
		}
	} else {
		if !s.cfg.UseGolangCILint {
			lintci := s.outputPath(".github", "workflows", "lint.yaml")
			slog.Debug("golangci-lint disabled, removing", "dst", lintci)
			if err := os.Remove(lintci); err != nil && !os.IsNotExist(err) {
				return err
			}
		}
		if !s.cfg.UseGoReleaser {
			releaseci := s.outputPath(".github", "workflows", "release.yaml")
			slog.Debug("goreleaser disabled, removing", "dst", releaseci)
			if err := os.Remove(releaseci); err != nil && !os.IsNotExist(err) {
				return err
			}
		}
	}

	if s.cfg.UseGSemver {
		dst := s.outputPath("Makefile")
		slog.Debug("generating Makefile", "dst", dst)
		if err := fsutil.CopyAndCompleteFile(pf,
			fmt.Sprintf("%s/projectfiles/Makefile", ver),
			dst, data); err != nil {
			return err
		}
	}

	hackdir := s.outputPath("hack")
	slog.Debug("generating hack files", "dst", hackdir)
	if err := fsutil.CopyAndCompleteFiles(pf,
		fmt.Sprintf("%s/projectfiles/hack", ver),
		hackdir, data); err != nil {
		return err
	}

	pkgdir := s.outputPath("pkg")
	slog.Debug("generating pkg files", "dst", pkgdir)
	return fsutil.CopyAndCompleteFiles(pf,
		fmt.Sprintf("%s/projectfiles/go/pkg", ver),
		pkgdir, data)
}

func (s *stages) copyFrontendToolchain(data *TemplateData) error {
	if !s.cfg.UseGithubActions {
		return nil
	}
	ver := s.cfg.SkeletonVersion
	pf := s.skeleton.ProjectFiles

	actionsdir := s.outputPath(".github")
	slog.Debug("generating github files", "dst", actionsdir)
	if err := fsutil.CopyAndCompleteFiles(pf,
		fmt.Sprintf("%s/projectfiles/.github", ver),
		actionsdir, data); err != nil {
		return err
	}

	// remove Go-specific workflows
	for _, f := range []string{"lint.yaml", "release.yaml"} {
		_ = os.Remove(s.outputPath(".github", "workflows", f))
	}

	// add nodejs CI workflow (npm-based frontend projects share it)
	src := fmt.Sprintf("%s/projectfiles/nodejs/.github", ver)
	return fsutil.CopyAndCompleteFiles(pf, src, actionsdir, data)
}

func showSummary(cfg *Config) {
	ui.ShowSummary(&ui.SummaryConfig{
		ProjectName:       cfg.ProjectName,
		ProjectType:       cfg.ProjectType,
		Language:          cfg.Language,
		Framework:         cfg.Framework,
		FrontendFramework: cfg.FrontendFramework,
		ModuleName:        cfg.ModuleName,
		UseGolangCILint:   cfg.UseGolangCILint,
		UseGoReleaser:     cfg.UseGoReleaser,
		UseGSemver:        cfg.UseGSemver,
		UseGithubActions:  cfg.UseGithubActions,
	})
}
