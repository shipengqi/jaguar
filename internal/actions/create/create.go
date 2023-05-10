package create

import (
	"embed"
	"fmt"
	"os"
	"path"

	"github.com/shipengqi/action"
	"github.com/shipengqi/jaguar/internal/actions/create/config"
	"github.com/shipengqi/jaguar/internal/actions/create/options"
	"github.com/shipengqi/jaguar/skeletons"
)

const (
	ActionName      = "new"
	ActionNameAlias = "n"
)

const (
	ProjectTypeCLI  = "cli"
	ProjectTypeAPI  = "api"
	ProjectTypeGRPC = "grpc"
)

const (
	DefaultSkeletonVersion = "v1"
)

const (
	FilenameGitIgnore = ".gitignore"
)

func NewAction(opts *options.Options, args []string) *action.Action {
	cfg, _ := config.CreateConfigFromOptions(opts, args)

	act := &action.Action{
		Name:   ActionName,
		PreRun: func(act *action.Action) error { return prerun(cfg) },
		Run:    func(act *action.Action) error { return create(cfg) },
	}

	return act
}

func create(cfg *config.Config) error {
	var fs embed.FS
	var skeleton string
	switch cfg.Type {
	case ProjectTypeCLI:
		fs = skeletons.CLI
		skeleton = ProjectTypeCLI
	case ProjectTypeAPI:
		fs = skeletons.API
		skeleton = ProjectTypeAPI
	case ProjectTypeGRPC:
		fs = skeletons.GRPC
		skeleton = ProjectTypeGRPC
	}

	if cfg.Force {
		err := cleanProject(cfg.ProjectName)
		if err != nil {
			return err
		}
	}

	return createProject(fs, DefaultSkeletonVersion, skeleton, cfg.ProjectName)
}

func cleanProject(project string) error {
	return os.RemoveAll(project)
}

func createProject(embedfs embed.FS, version, skeleton, project string) error {
	err := createProjectFiles(embedfs, version, skeleton, project)
	if err != nil {
		return err
	}
	return createGitIgnore(version, project)
}

func createProjectFiles(embedfs embed.FS, version, skeleton, project string) error {
	from := fmt.Sprintf("%s/%s", version, skeleton)
	var count int
	err := CalculateFilesFromEmbedFS(embedfs, from, &count)
	if err != nil {
		return err
	}
	bar := newBar(count, "1/2", "Creating project files")

	err = CopyWithBar(bar, embedfs, from, project)
	if err != nil {
		return err
	}
	return nil
}

func createGitIgnore(version, project string) error {
	bar := newBar(1, "2/2", "Creating .gitignore")
	return CopyFileWithBar(bar, skeletons.GitIgnore,
		fmt.Sprintf("%s/%s", version, FilenameGitIgnore),
		path.Join(project, FilenameGitIgnore))
}

func createLintConfig(version, project string) error {
	bar := newBar(1, "2/2", "Creating .gitignore")
	return CopyFileWithBar(bar, skeletons.GitIgnore,
		fmt.Sprintf("%s/%s", version, FilenameGitIgnore),
		path.Join(project, FilenameGitIgnore))
}
