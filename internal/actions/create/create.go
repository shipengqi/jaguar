package create

import (
	"embed"
	"github.com/manifoldco/promptui"
	"github.com/shipengqi/action"
	"github.com/shipengqi/log"
	"io/fs"
	"os"
	"path"

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

func NewAction(opts *options.Options) *action.Action {
	cfg, _ := config.CreateConfigFromOptions(opts)
	act := &action.Action{
		Name:   ActionName,
		PreRun: func(act *action.Action) error { return prerun(cfg) },
		Run:    func(act *action.Action) error { return create(cfg) },
	}

	return act
}

func prerun(cfg *config.Config) error {
	promp := func(label string, items []string) string {
		prompt := promptui.Select{
			Label: label,
			Items: items,
		}
		prompt.HideSelected = true
		_, result, err := prompt.Run()
		if err != nil {
			return ""
		}
		return result
	}

	if cfg.Type == "" {
		var selected string
		switch promp("Select project type", []string{"CLI", "API", "gRPC"}) {
		case "CLI":
			selected = ProjectTypeCLI
			break
		case "API":
			selected = ProjectTypeAPI
			break
		case "gRPC":
			selected = ProjectTypeGRPC
			break
		default:
			return nil
		}
		cfg.Type = selected
	}

	return nil
}

func create(cfg *config.Config) error {
	switch cfg.Type {
	case ProjectTypeCLI:
		err := Copy(skeletons.CLI, "cli", "newcli")
		if err != nil {
			log.Warnf("copy: %s", err.Error())
			return err
		}
		CopyFile(skeletons.GitIgnore, ".gitignore", "newcli/.gitignore")
		break
	case ProjectTypeAPI:
		err := Copy(skeletons.API, "api", "newapi")
		if err != nil {
			log.Warnf("copy: %s", err.Error())
			return err
		}
		break
	case ProjectTypeGRPC:
		err := Copy(skeletons.GRPC, "grpc", "newgrpc")
		if err != nil {
			log.Warnf("copy: %s", err.Error())
			return err
		}
		break
	}
	return nil
}

// CopyFile copies a file from src to dst.
func CopyFile(embedfs embed.FS, src, dst string) (err error) {
	log.Infof("read file %s", src)
	sdata, err := embedfs.ReadFile(src)
	if err != nil {
		log.Errorf("read file %s: %s", src, err.Error())
		return
	}
	info, err := embedfs.Open(src)
	if err != nil {
		return err
	}
	sinfo, err := info.Stat()
	if err != nil {
		log.Errorf("stat embed/%s: %s", src, err.Error())
		return err
	}
	return os.WriteFile(dst, sdata, sinfo.Mode())
}

// Copy copies a file or directory from src to dst.
func Copy(embedfs embed.FS, src, dst string) error {
	var (
		err   error
		fds   []os.DirEntry
		sinfo fs.FileInfo
	)

	sfd, err := embedfs.Open(src)
	if err != nil {
		log.Errorf("open embed/%s: %s", src, err.Error())
		return err
	}
	sinfo, err = sfd.Stat()
	if err != nil {
		log.Errorf("stat embed/%s: %s", src, err.Error())
		return err
	}
	log.Infof("copy embed/%s", src)
	// copies a file
	if !sinfo.IsDir() {
		log.Infof("embed/%s is file", src)
		return CopyFile(embedfs, src, dst)
	}
	log.Infof("embed/%s is dir", src)
	// tries to create dst directory
	if err = os.MkdirAll(dst, sinfo.Mode()); err != nil {
		return err
	}
	log.Infof("read dir embed/%s", src)
	if fds, err = embedfs.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		log.Infof("%s under dir embed/%s", fd.Name(), src)
		//sfp := path.Join(src, fd.Name())
		dfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = Copy(embedfs, src+"/"+fd.Name(), dfp); err != nil {
				return err
			}
		} else {
			if err = CopyFile(embedfs, src+"/"+fd.Name(), dfp); err != nil {
				return err
			}
		}
	}
	return nil
}
