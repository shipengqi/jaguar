package license

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/shipengqi/jaguar/internal/actions/license/config"
)

func NewAddLicenseAction(cfg *config.Config, args []string) func() error {
	return func() error {
		ch := make(chan *file, 1000)
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			addFiles(ch, cfg)
		}()
		for _, d := range args {
			walk(ch, d, cfg.SkipDirRegs, cfg.SkipFileRegs)
		}
		close(ch)
		wg.Wait()
		return nil
	}
}

func addFiles(ch <-chan *file, cfg *config.Config) {
	sem := make(chan struct{}, 100)
	var wg sync.WaitGroup
	for f := range ch {
		f := f
		sem <- struct{}{}
		wg.Add(1)
		go func() {
			defer func() { <-sem; wg.Done() }()
			modified, err := addLicense(f.path, f.mode, cfg.LicenseTmpl, &copyrightInfo{cfg.HeaderOptions.Year, cfg.HeaderOptions.Holder})
			if err != nil {
				slog.Warn("add license", "path", f.path, "err", err)
				return
			}
			if modified {
				fmt.Printf("%s: license added\n", f.path)
			}
		}()
	}
	wg.Wait()
}
