package create

import (
	"embed"
	"io/fs"
	"os"
	"path"

	pb "github.com/schollz/progressbar/v3"
)

// CopyFile copies a file from src to dst.
func CopyFile(embedfs embed.FS, src, dst string) (err error) {
	return CopyFileWithBar(nil, embedfs, src, dst)
}

func CopyFileWithBar(bar *pb.ProgressBar, embedfs embed.FS, src, dst string) (err error) {
	defer func(b *pb.ProgressBar) {
		if b != nil {
			_ = b.Add(1)
		}
	}(bar)
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
	return CopyWithBar(nil, embedfs, src, dst)
}

// CopyWithBar copies a file or directory from src to dst.
func CopyWithBar(bar *pb.ProgressBar, embedfs embed.FS, src, dst string) error {
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
		return CopyFileWithBar(bar, embedfs, src, dst)
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
			if err = CopyWithBar(bar, embedfs, sfp, dfp); err != nil {
				return err
			}
		} else {
			if err = CopyFileWithBar(bar, embedfs, sfp, dfp); err != nil {
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
