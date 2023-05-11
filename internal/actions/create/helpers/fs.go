package helpers

import (
	"embed"
	"io/fs"
	"os"
	"path"
)

// CopyFile copies a file from src to dst.
func CopyFile(embedfs embed.FS, src, dst string) (err error) {
	sdata, err := embedfs.ReadFile(src)
	if err != nil {
		return
	}
	info, err := embedfs.Open(src)
	if err != nil {
		return err
	}
	sinfo, err := info.Stat()
	if err != nil {
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
		return err
	}
	sinfo, err = sfd.Stat()
	if err != nil {
		return err
	}
	// copies a file
	if !sinfo.IsDir() {
		return CopyFile(embedfs, src, dst)
	}
	// tries to create dst directory
	if err = os.MkdirAll(dst, sinfo.Mode()); err != nil {
		return err
	}
	if fds, err = embedfs.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		sfp := path.Join(src, fd.Name())
		dfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = Copy(embedfs, sfp, dfp); err != nil {
				return err
			}
		} else {
			if err = CopyFile(embedfs, sfp, dfp); err != nil {
				return err
			}
		}
	}
	return nil
}

func CalculateFilesFromEmbedFS(embedfs embed.FS, src string, result *int) error {
	var (
		err error
		fds []os.DirEntry
	)
	if fds, err = embedfs.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		sfp := path.Join(src, fd.Name())
		if fd.IsDir() {
			if err = CalculateFilesFromEmbedFS(embedfs, sfp, result); err != nil {
				return err
			}
		} else {
			*result++
		}
	}
	return nil
}
