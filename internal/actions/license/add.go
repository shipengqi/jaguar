package license

import (
	"github.com/shipengqi/action"
	"github.com/shipengqi/log"
	"github.com/sourcegraph/conc/pool"

	"github.com/shipengqi/jaguar/internal/actions/license/config"
)

func NewAddLicenseAction(cfg *config.Config, args []string) *action.Action {
	act := &action.Action{
		Name: ActionNameAdd,
		Run: func(act *action.Action) error {
			// process at most 1000 files in parallel
			ch := make(chan *file, 1000)
			done := make(chan struct{})
			go addFiles(ch, done, cfg)
			for _, d := range args {
				walk(ch, d, cfg.SkipDirRegs, cfg.SkipFileRegs)
			}
			close(ch)
			<-done
			return nil
		},
	}

	return act
}

func addFiles(ch chan *file, done chan struct{}, cfg *config.Config) {
	p := pool.New().WithMaxGoroutines(100)
	for f := range ch {
		fi := f // https://golang.org/doc/faq#closures_and_goroutines
		p.Go(addFile(fi, cfg))
	}
	p.Wait()
	close(done)
}

func addFile(f *file, cfg *config.Config) func() {
	return func() {
		modified, err := addLicense(f.path, f.mode, cfg.LicenseTmpl, &copyrightInfo{cfg.HeaderOptions.Year,
			cfg.HeaderOptions.Holder})
		if err != nil {
			log.Warnf("%s: %s", f.path, err.Error())
			return
		}
		if modified {
			log.Infof("%s: license added", f.path)
		}
	}
}
