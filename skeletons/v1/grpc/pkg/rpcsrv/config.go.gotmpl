package rpcsrv

import (
	"context"
	"fmt"

	"github.com/shipengqi/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

type Config struct {
	*Options

	id                       string
	domain                   string
	addr                     string
	grpcopts                 []grpc.ServerOption
	unaryServerInterceptors  []grpc.UnaryServerInterceptor
	streamServerInterceptors []grpc.StreamServerInterceptor
	ctx                      context.Context
}

func CreateConfigFromOptions(opts *Options) *Config {
	return &Config{
		Options:                  opts,
		id:                       "",
		domain:                   "",
		addr:                     fmt.Sprintf("%s:%d", opts.BindAddress, opts.BindPort),
		grpcopts:                 initGrpcServerOptions(opts),
		unaryServerInterceptors:  nil,
		streamServerInterceptors: nil,
		ctx:                      nil,
	}
}

func initGrpcServerOptions(opts *Options) []grpc.ServerOption {
	gopts := []grpc.ServerOption{
		grpc.MaxConcurrentStreams(uint32(opts.MaxConcurrentStreams)),
		grpc.MaxRecvMsgSize(opts.MaxMsgSize),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime: opts.Keepalive / defaultMiniKeepAliveTimeRate,
		}),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:                  opts.Keepalive,
			Timeout:               opts.Timeout,
			MaxConnectionAge:      opts.MaxConnectionAge,
			MaxConnectionAgeGrace: opts.MaxConnectionAgeGrace,
		}),
	}

	if opts.ServerCert.CertFile != "" && opts.ServerCert.KeyFile != "" {
		if cs, err := credentials.NewServerTLSFromFile(opts.ServerCert.CertFile, opts.ServerCert.KeyFile); err != nil {
			log.Warnf("failed to generate transport credentials, use insecure mode.")
			log.Debugf("credentials.NewServerTLSFromFile err: %v", err)
		} else {
			gopts = append(gopts, grpc.Creds(cs))
		}
	}

	return gopts
}
