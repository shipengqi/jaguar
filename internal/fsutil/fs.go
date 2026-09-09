package fsutil

import (
	"embed"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"
	"text/template"
)

const GoTemplateSuffix = ".tpl"

// CopyAndCompleteFile copies a single file from an embedded FS, rendering it as a Go template if it has .gotmpl suffix.
func CopyAndCompleteFile(embedfs embed.FS, src, dst string, data any) error {
	if strings.HasSuffix(src, GoTemplateSuffix) {
		return copyAndCompleteGoTemplate(embedfs, src, dst, data)
	}
	return copyFile(embedfs, src, dst)
}

// CopyAndCompleteFiles copies a file or directory recursively from an embedded FS.
func CopyAndCompleteFiles(embedfs embed.FS, src, dst string, data any) error {
	sfd, err := embedfs.Open(src)
	if err != nil {
		return err
	}
	sinfo, err := sfd.Stat()
	if err != nil {
		return err
	}
	if !sinfo.IsDir() {
		return CopyAndCompleteFile(embedfs, src, dst, data)
	}
	if err = os.MkdirAll(dst, 0o700); err != nil {
		return err
	}
	fds, err := embedfs.ReadDir(src)
	if err != nil {
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

func copyFile(embedfs embed.FS, src, dst string) error {
	data, err := embedfs.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0o600)
}

func copyAndCompleteGoTemplate(embedfs embed.FS, src, dst string, data any) error {
	tmpl, err := template.ParseFS(embedfs, src)
	if err != nil {
		return err
	}
	tmpf, err := os.CreateTemp("", "jaguar_tmpl_")
	if err != nil {
		return err
	}
	defer func() { _ = os.Remove(tmpf.Name()) }() //nolint:gosec // path is constructed internally, no user input

	if err = tmpl.Execute(tmpf, data); err != nil {
		return err
	}
	if _, err = tmpf.Seek(0, 0); err != nil {
		return err
	}

	outpath := strings.TrimSuffix(dst, GoTemplateSuffix)
	if err = os.MkdirAll(path.Dir(outpath), 0o700); err != nil {
		return err
	}
	outf, err := os.Create(outpath)
	if err != nil {
		return err
	}
	defer func() {
		_ = tmpf.Close()
		_ = outf.Close()
	}()

	_, err = io.Copy(outf, tmpf)
	return err
}

// IsDir reports whether path is a directory.
func IsDir(p string) bool {
	fi, err := os.Stat(p)
	return err == nil && fi.IsDir()
}

// ReadDirEntries returns DirEntry slice for a path in an embed.FS.
func ReadDirEntries(embedfs embed.FS, dir string) ([]fs.DirEntry, error) {
	return embedfs.ReadDir(dir)
}
