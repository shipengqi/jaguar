package license

import (
	"github.com/shipengqi/action"
	"github.com/shipengqi/log"
	"github.com/sourcegraph/conc/pool"

	"github.com/shipengqi/jaguar/internal/actions/license/config"
)

func NewCheckLicenseAction(cfg *config.Config, args []string) *action.Action {
	act := &action.Action{
		Name: ActionNameCheck,
		Run: func(_ *action.Action) error {
			// process at most 1000 files in parallel
			ch := make(chan *file, 1000)
			done := make(chan struct{})
			go checkFiles(ch, done)
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

func checkFiles(ch chan *file, done chan struct{}) {
	p := pool.New().WithMaxGoroutines(100)
	for f := range ch {
		fpath := f.path // https://golang.org/doc/faq#closures_and_goroutines
		p.Go(checkFile(fpath))
	}
	p.Wait()
	close(done)
}

func checkFile(fpath string) func() {
	return func() {
		// Check if file extension is known
		_, _, _, unknown := licenseCharsForExt(fpath)
		if unknown {
			log.Warnf("%s: unknown file extension", fpath)
			return
		}
		// Check if file has a license
		isMissingLicenseHeader, err := fileHasLicense(fpath)
		if err != nil {
			log.Warnf("%s: %s", fpath, err.Error())
			return
		}
		if isMissingLicenseHeader {
			log.Warnf("%s: missing license header", fpath)
		}
	}
}
