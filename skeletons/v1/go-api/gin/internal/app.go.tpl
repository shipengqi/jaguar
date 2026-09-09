package internal

import (
	"{{ .App.ModuleName }}/internal/config"
	"{{ .App.ModuleName }}/internal/options"
	"{{ .App.ModuleName }}/pkg/app"
	"{{ .App.ModuleName }}/pkg/xlog"
)

const desc = `The {{ .App.Name }} API server validates and configures data for the api objects.
The API Server services REST operations to do the api objects management.

Find more {{ .App.NormalizedName }}-apiserver information at:
    {{ .App.DocumentLink }}`

func NewApp() *app.App {
	opts := options.New()
	return app.New("{{ .App.Name }} API Server",
		app.WithCliOptions(opts),
		app.WithDesc(desc),
		app.WithRunFunc(run(opts)),
	)
}

func run(opts *options.Options) app.RunFunc {
	return func() error {
		xlog.Init(opts.Log)
		defer func() { _ = xlog.Sync() }()

		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}

		return Run(cfg)
	}
}

// Run runs the specified APIServer. This should never exit.
func Run(cfg *config.Config) error {
	server, err := CreateAPIServer(cfg)
	if err != nil {
		return err
	}

	return server.PreRun().Run()
}
