package internal

import (
	"github.com/shipengqi/jcli"
	"github.com/shipengqi/log"

	"github.com/jaguar/grpcskeleton/internal/config"
	"github.com/jaguar/grpcskeleton/internal/options"
)

const desc = `The {{example}} gRPC server validates and configures data
for the api objects. The gRPC Server services REST operations to do 
the api objects management.

Find more {{example}}-grpcserver information at:
    {{example.document.link}}`

func NewApp() *jcli.App {
	opts := options.New()
	application := jcli.New("{{example}} gRPC API Server",
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
	server, err := CreateGRPCServer(cfg)
	if err != nil {
		return err
	}

	return server.PreRun().Run()
}
