package stages

import (
	"embed"
	"fmt"

	"github.com/shipengqi/jaguar/internal/actions/create/config"
	"github.com/shipengqi/jaguar/internal/actions/create/helpers"
	"github.com/shipengqi/jaguar/internal/actions/create/options"
	"github.com/shipengqi/jaguar/internal/actions/create/types"
	"github.com/shipengqi/jaguar/skeletons"
)

const (
	FilenameGitIgnore = ".gitignore"
	FilenameGoCI      = ".golangci.yaml"
	FilenameReleaser  = ".goreleaser.yaml"
	FilenameSemver    = ".gsemver.yaml"
)

const (
	titleClean     = "Cleaning project"
	titleProject   = "Creating %s project files"
	titleMakeRules = "Creating make rules"
	titleSwagger   = "Configuring swagger"
	titleInitMod   = "Initializing Go module"
	titleModTidy   = "Installing dependencies"
)

var fsmap = make(map[string]embed.FS)

type Stages struct {
	cfg      *config.Config
	skeleton skeletons.Skeleton
}

func New(cfg *config.Config) *Stages {
	return &Stages{cfg: cfg}
}

func (s *Stages) initialize() error {
	ss := skeletons.New()
	switch s.cfg.SkeletonVersion {
	case options.SkeletonVersion1:
		s.skeleton = ss.V1
	default:
		return fmt.Errorf("unknown version '%s'", s.cfg.SkeletonVersion)
	}
	return nil
}

func (s *Stages) Run() error {
	var err error

	if err = s.initialize(); err != nil {
		return err
	}

	var efs embed.FS
	data := s.cfg.ExportTemplateData()
	src := fmt.Sprintf("%s/%s", s.cfg.SkeletonVersion, s.cfg.ProjectType)

	switch s.cfg.ProjectType {
	case types.ProjectTypeAPI:
		src = fmt.Sprintf("%s/%s/%s", s.cfg.SkeletonVersion, types.ProjectTypeAPI, s.cfg.GoFramework)
		efs = s.skeleton.API
	case types.ProjectTypeCLI:
		efs = s.skeleton.CLI
	case types.ProjectTypeGRPC:
		efs = s.skeleton.GRPC
	}
	if err = helpers.CopyAndCompleteFiles(efs, src, s.cfg.ProjectName, data); err != nil {
		return err
	}
	if s.cfg.IsUseGolangCILint {
		if err = helpers.CopyAndCompleteFile(s.skeleton.ProjectFiles,
			fmt.Sprintf("%s/projectfiles/.golangci.yaml.gotmpl", s.cfg.SkeletonVersion),
			fmt.Sprintf("%s/.golangci.yaml", s.cfg.ProjectName), data); err != nil {
			return err
		}
	}
	if s.cfg.IsUseGoReleaser {
		if err = helpers.CopyAndCompleteFile(s.skeleton.ProjectFiles,
			fmt.Sprintf("%s/projectfiles/.goreleaser.yaml.gotmpl", s.cfg.SkeletonVersion),
			fmt.Sprintf("%s/.goreleaser.yaml", s.cfg.ProjectName), data); err != nil {
			return err
		}
	}
	if s.cfg.IsUseGSemver {
		if err = helpers.CopyAndCompleteFile(s.skeleton.ProjectFiles,
			fmt.Sprintf("%s/projectfiles/.gsemver.yaml", s.cfg.SkeletonVersion),
			fmt.Sprintf("%s/.gsemver.yaml", s.cfg.ProjectName), data); err != nil {
			return err
		}
	}
	return nil
}
