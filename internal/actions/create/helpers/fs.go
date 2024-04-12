package helpers

import (
	"embed"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/shipengqi/jaguar/internal/actions/create/types"
)

const (
	GoTemplateSuffix = ".gotmpl"
)

// CopyAndCompleteFile copies a template from src to dst, and complete it.
func CopyAndCompleteFile(embedfs embed.FS, src, dst string, data *types.TemplateData) (err error) {
	if strings.HasSuffix(src, GoTemplateSuffix) {
		return CopyAndCompleteGoTemplate(embedfs, src, dst, data)
	}
	return CopyFile(embedfs, src, dst)
}

// CopyAndCompleteFiles copies a file or directory from src to dst.
func CopyAndCompleteFiles(embedfs embed.FS, src, dst string, data *types.TemplateData) error {
	var (
		err   error
		fds   []os.DirEntry
		sinfo fs.FileInfo
	)

	sfd, err := embedfs.Open(src)
	if err != nil {
		return err
	}
	sinfo, err = sfd.Stat()
	if err != nil {
		return err
	}
	// copies a file
	if !sinfo.IsDir() {
		return CopyAndCompleteFile(embedfs, src, dst, data)
	}
	// tries to create dst directory
	if err = os.MkdirAll(dst, 0o700); err != nil {
		return err
	}
	if fds, err = embedfs.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		sfp := path.Join(src, fd.Name())
		dfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = CopyAndCompleteFiles(embedfs, sfp, dfp, data); err != nil {
				return err
			}
		} else {
			if err = CopyAndCompleteFile(embedfs, sfp, dfp, data); err != nil {
				return err
			}
		}
	}
	return nil
}

// CopyFile copies a file from src to dst.
func CopyFile(embedfs embed.FS, src, dst string) (err error) {
	sdata, err := embedfs.ReadFile(src)
	if err != nil {
		return
	}
	return os.WriteFile(dst, sdata, 0o600)
}

func CopyAndCompleteGoTemplate(embedfs embed.FS, src, dst string, data *types.TemplateData) (err error) {
	// parse template from embed file system
	tmpl, err := template.ParseFS(embedfs, src)
	if err != nil {
		return err
	}
	// create a new temp file with the given data.
	tempf, err := os.CreateTemp("", "TEMPLATE_")
	if err != nil {
		return err
	}
	defer func() { _ = os.Remove(tempf.Name()) }()

	// set variables to the template.
	if err = tmpl.Execute(tempf, data); err != nil {
		return err
	}

	// Reset the record position to the beginning of the file.
	if _, err = tempf.Seek(0, 0); err != nil {
		return err
	}

	outputf, err := os.Create(strings.TrimSuffix(dst, ".gotmpl"))
	if err != nil {
		return err
	}
	defer func() {
		_ = tempf.Close()
		_ = outputf.Close()
	}()

	// Copy file from the temp file to the output.
	if _, err = io.Copy(outputf, tempf); err != nil {
		return err
	}

	return err
}
