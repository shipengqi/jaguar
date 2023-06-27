package addlicense

import (
	"errors"
	"fmt"
	"os"

	"github.com/shipengqi/action"
	"github.com/spf13/pflag"
	"golang.org/x/sync/errgroup"

	"github.com/shipengqi/jaguar/internal/actions/addlicense/config"
)

func newCheckAction(cfg *config.Config) *action.Action {
	act := &action.Action{
		Name: SubActionNameCheck,
		Executable: func(_ *action.Action) bool {
			return cfg.Check
		},
		Run: func(act *action.Action) error {
			// process at most 1000 files in parallel
			ch := make(chan *file, 1000)
			done := make(chan struct{})

			go func() {
				var wg errgroup.Group
				for f := range ch {
					fpath := f.path // https://golang.org/doc/faq#closures_and_goroutines
					wg.Go(func() error {
						// Check if file extension is known
						lic, err := licenseHeader(fpath, cfg.LicenseTmpl, &copyrightInfo{cfg.Year, cfg.Holder})
						if err != nil {
							fmt.Printf("%s: %v\n", fpath, err)
							return err
						}
						if lic == nil { // Unknown fileExtension
							return nil
						}
						// Check if file has a license
						isMissingLicenseHeader, err := fileHasLicense(fpath)
						if err != nil {
							fmt.Printf("%s: %v\n", fpath, err)
							return err
						}
						if isMissingLicenseHeader {
							fmt.Printf("%s\n", fpath)
							return errors.New("missing license header")
						}
						return nil
					})
				}
				err := wg.Wait()
				close(done)
				if err != nil {
					os.Exit(1)
				}
			}()

			for _, d := range pflag.Args() {
				walk(ch, d)
			}
			close(ch)
			<-done

			return nil
		},
	}

	return act
}
