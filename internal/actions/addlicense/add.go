package addlicense

import (
	"fmt"
	"os"

	"github.com/shipengqi/action"
	"github.com/spf13/pflag"
	"golang.org/x/sync/errgroup"

	"github.com/shipengqi/jaguar/internal/actions/addlicense/config"
)

func newAddAction(cfg *config.Config) *action.Action {
	act := &action.Action{
		Name: SubActionNameAdd,
		Executable: func(_ *action.Action) bool {
			return !cfg.Check
		},
		Run: func(act *action.Action) error {
			// process at most 1000 files in parallel
			ch := make(chan *file, 1000)
			done := make(chan struct{})

			go func() {
				var wg errgroup.Group
				for f := range ch {
					f := f // https://golang.org/doc/faq#closures_and_goroutines
					wg.Go(func() error {
						modified, err := addLicense(f.path, f.mode, cfg.LicenseTmpl, &copyrightInfo{cfg.Year, cfg.Holder})
						if err != nil {
							fmt.Printf("%s: %v\n", f.path, err)
							return err
						}
						if modified {
							fmt.Printf("%s added license\n", f.path)
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
