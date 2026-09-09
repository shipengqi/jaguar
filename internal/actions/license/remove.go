package license

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"sync"

	"github.com/shipengqi/jaguar/internal/actions/license/config"
)

func NewRemoveLicenseAction(cfg *config.Config, args []string) func() error {
	return func() error {
		ch := make(chan *file, 1000)
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			removeFiles(ch, cfg)
		}()
		for _, d := range args {
			walk(ch, d, cfg.SkipDirRegs, cfg.SkipFileRegs)
		}
		close(ch)
		wg.Wait()
		return nil
	}
}

func removeFiles(ch <-chan *file, cfg *config.Config) {
	sem := make(chan struct{}, 100)
	var wg sync.WaitGroup
	for f := range ch {
		f := f
		sem <- struct{}{}
		wg.Add(1)
		go func() {
			defer func() { <-sem; wg.Done() }()
			removeFile(f, cfg)
		}()
	}
	wg.Wait()
}

func removeFile(f *file, cfg *config.Config) {
	lic, err := licenseHeader(f.path, cfg.LicenseTmpl, &copyrightInfo{cfg.HeaderOptions.Year, cfg.HeaderOptions.Holder})
	if err != nil {
		slog.Debug("remove license", "path", f.path, "err", err)
		return
	}
	if lic == nil {
		slog.Debug("unknown file extension", "path", f.path)
		return
	}

	b, err := os.ReadFile(f.path)
	if err != nil {
		slog.Debug("read file", "path", f.path, "err", err)
		return
	}

	if !bytes.Contains(b, lic) {
		slog.Debug("skipped", "path", f.path)
		return
	}
	modified := bytes.Replace(b, lic, []byte{}, 1)
	if err = os.WriteFile(f.path, modified, f.mode); err != nil {
		slog.Debug("write file", "path", f.path, "err", err)
		return
	}
	fmt.Printf("%s: license removed\n", f.path)
}
