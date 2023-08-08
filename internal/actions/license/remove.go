package license

import (
	"bytes"
	"os"

	"github.com/shipengqi/action"
	"github.com/shipengqi/log"
	"github.com/sourcegraph/conc/pool"

	"github.com/shipengqi/jaguar/internal/actions/license/config"
)

func NewRemoveLicenseAction(cfg *config.Config, args []string) *action.Action {
	act := &action.Action{
		Name: ActionNameRemove,
		Run: func(act *action.Action) error {
			// process at most 1000 files in parallel
			ch := make(chan *file, 1000)
			done := make(chan struct{})
			go removeFiles(ch, done, cfg)
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

func removeFiles(ch chan *file, done chan struct{}, cfg *config.Config) {
	p := pool.New().WithMaxGoroutines(100)
	for f := range ch {
		fi := f // https://golang.org/doc/faq#closures_and_goroutines
		p.Go(removeFile(fi, cfg))
	}
	p.Wait()
	close(done)
}

func removeFile(f *file, cfg *config.Config) func() {
	return func() {
		var lic []byte
		var err error
		lic, err = licenseHeader(f.path, cfg.LicenseTmpl, &copyrightInfo{cfg.HeaderOptions.Year,
			cfg.HeaderOptions.Holder})
		if err != nil {
			log.Debugf("%s: %s", f.path, err.Error())
			return
		}
		if lic == nil {
			log.Debugf("%s: unknown file extension", f.path)
			return
		}

		b, err := os.ReadFile(f.path)
		if err != nil {
			log.Debugf("%s: %s", f.path, err.Error())
			return
		}

		if !bytes.Contains(b, lic) {
			log.Infof("%s: skipped", f.path)
			return
		}
		modified := bytes.Replace(b, lic, []byte{}, 1)
		err = os.WriteFile(f.path, modified, f.mode)
		if err != nil {
			log.Debugf("%s: %s", f.path, err.Error())
			return
		}

		log.Infof("%s: license removed", f.path)
	}
}
