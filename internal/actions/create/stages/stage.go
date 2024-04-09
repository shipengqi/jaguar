package stages

import (
	"context"
	"embed"
	"fmt"
	"github.com/shipengqi/log"
	"os"

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

type Processor func(cfg *config.Config) error

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

func (s *Stages) Run(cancel context.CancelFunc) error {
	var err error

	defer func() { cancel() }()

	log.Debug("initializing ...")
	if err = s.initialize(); err != nil {
		return err
	}

	log.Debug("initializing template data ...")
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
	log.Debugf("generating code files to '%s'...", s.cfg.ProjectName)
	if err = helpers.CopyAndCompleteFiles(efs, src, s.cfg.ProjectName, data); err != nil {
		return err
	}

	lintfile := fmt.Sprintf("%s/.golangci.yaml", s.cfg.ProjectName)
	log.Debugf("generating golangci-lint configuration file to '%s' ...", lintfile)
	if s.cfg.IsUseGolangCILint {
		if err = helpers.CopyAndCompleteFile(s.skeleton.ProjectFiles,
			fmt.Sprintf("%s/projectfiles/.golangci.yaml.gotmpl", s.cfg.SkeletonVersion),
			lintfile, data); err != nil {
			return err
		}
	}

	releaserfile := fmt.Sprintf("%s/.goreleaser.yaml", s.cfg.ProjectName)
	log.Debugf("generating goreleaser configuration file to '%s'...", releaserfile)
	if s.cfg.IsUseGoReleaser {
		if err = helpers.CopyAndCompleteFile(s.skeleton.ProjectFiles,
			fmt.Sprintf("%s/projectfiles/.goreleaser.yaml.gotmpl", s.cfg.SkeletonVersion),
			releaserfile, data); err != nil {
			return err
		}
	}

	semverfile := fmt.Sprintf("%s/.gsemver.yaml", s.cfg.ProjectName)
	log.Debugf("generating gsemver configuration file to '%s'...", semverfile)
	if s.cfg.IsUseGSemver {
		if err = helpers.CopyAndCompleteFile(s.skeleton.ProjectFiles,
			fmt.Sprintf("%s/projectfiles/.gsemver.yaml", s.cfg.SkeletonVersion),
			semverfile, data); err != nil {
			return err
		}
	}

	actionsdir := fmt.Sprintf("%s/.github", s.cfg.ProjectName)
	log.Debugf("generating github repository configuration files to '%s'...", actionsdir)
	if err = helpers.CopyAndCompleteFiles(s.skeleton.ProjectFiles,
		fmt.Sprintf("%s/projectfiles/.github", s.cfg.SkeletonVersion),
		actionsdir, data); err != nil {
		return err
	}

	if !s.cfg.IsUseGithubActions {
		workflowdir := fmt.Sprintf("%s/.github/workflows", s.cfg.ProjectName)
		log.Debugf("Github Actions is disabled, remove '%s'", workflowdir)
		if err = os.RemoveAll(workflowdir); err != nil {
			return err
		}
	} else {
		if !s.cfg.IsUseGolangCILint {
			lintcifile := fmt.Sprintf("%s/.github/workflows/lint.yaml", s.cfg.ProjectName)
			log.Debugf("golangci-lint is disabled, remove '%s'", lintcifile)
			if err = os.Remove(lintcifile); err != nil {
				return err
			}
		}
	}

	makefile := fmt.Sprintf("%s/Makefile", s.cfg.ProjectName)
	log.Debugf("generating makefile to '%s'...", makefile)
	if s.cfg.IsUseGSemver {
		if err = helpers.CopyAndCompleteFile(s.skeleton.ProjectFiles,
			fmt.Sprintf("%s/projectfiles/Makefile", s.cfg.SkeletonVersion),
			makefile, data); err != nil {
			return err
		}
	}
	
	hackdir := fmt.Sprintf("%s/hack", s.cfg.ProjectName)
	log.Debugf("generating build files '%s'...", hackdir)
	if err = helpers.CopyAndCompleteFiles(s.skeleton.ProjectFiles,
		fmt.Sprintf("%s/projectfiles/hack", s.cfg.SkeletonVersion),
		hackdir, data); err != nil {
		return err
	}
	return nil
}
