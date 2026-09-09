package license

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/shipengqi/jaguar/internal/actions/license/config"
)

func NewCheckLicenseAction(cfg *config.Config, args []string) func() error {
	return func() error {
		ch := make(chan *file, 1000)
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			checkFiles(ch)
		}()
		for _, d := range args {
			walk(ch, d, cfg.SkipDirRegs, cfg.SkipFileRegs)
		}
		close(ch)
		wg.Wait()
		return nil
	}
}

func checkFiles(ch <-chan *file) {
	sem := make(chan struct{}, 100)
	var wg sync.WaitGroup
	for f := range ch {
		fpath := f.path
		sem <- struct{}{}
		wg.Add(1)
		go func() {
			defer func() { <-sem; wg.Done() }()
			_, _, _, unknown := licenseCharsForExt(fpath)
			if unknown {
				slog.Debug("unknown file extension", "path", fpath)
				return
			}
			missing, err := fileHasLicense(fpath)
			if err != nil {
				slog.Warn("check license", "path", fpath, "err", err)
				return
			}
			if missing {
				fmt.Printf("%s: missing license header\n", fpath)
			}
		}()
	}
	wg.Wait()
}
