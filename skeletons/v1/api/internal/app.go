package internal

import (
	"github.com/jaguar/apiskeleton/internal/config"
	"github.com/jaguar/apiskeleton/internal/options"
	"github.com/shipengqi/jcli"
	"github.com/shipengqi/log"
)

const desc = `The {{example}} API server validates and configures data
for the api objects. The API Server services REST operations to do 
the api objects management.

Find more {{example}}-apiserver information at:
    {{example.document.link}}`

func NewApp() *jcli.App {
	opts := options.New()
	application := jcli.New("{{example}} API Server",
		jcli.WithCliOptions(opts),
		jcli.WithDesc(desc),
		jcli.WithRunFunc(run(opts)),
	)
	return application
}

func run(opts *options.Options) jcli.RunFunc {
	return func() error {
		log.Configure(opts.Log)
		defer func() { _ = log.Flush() }()

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
