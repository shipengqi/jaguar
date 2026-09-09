package internal

import (
	"{{ .App.ModuleName }}/pkg/app"
	"{{ .App.ModuleName }}/pkg/xlog"

	"{{ .App.ModuleName }}/internal/config"
	"{{ .App.ModuleName }}/internal/options"
)

const desc = `The {{ .App.Name }} API server validates and configures data for the api objects.
The API Server services REST operations to do the api objects management.

Find more {{ .App.NormalizedName }}-apiserver information at:
    {{ .App.DocumentLink }}`

func NewApp() *app.App {
	opts := options.New()
	application := app.New("{{ .App.Name }} API Server",
		app.WithCliOptions(opts),
		app.WithDesc(desc),
		app.WithRunFunc(run(opts)),
	)
	return application
}

func run(opts *options.Options) app.RunFunc {
    // setting up the global logger before the application runs
    xlog.Init(opts.Log)
	return func() error {
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
