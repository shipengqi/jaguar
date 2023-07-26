package rpcsrv

import (
	"context"
	"crypto/tls"
	"fmt"
	"math"
	"time"

	"github.com/shipengqi/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

const (
	infinity                           = time.Duration(math.MaxInt64)
	defaultMaxMsgSize                  = 4 << 20
	defaultMaxConcurrentStreams        = 100000
	defaultKeepAliveTime               = 30 * time.Second
	defaultConnectionIdleTime          = 10 * time.Second
	defaultMaxServerConnectionAgeGrace = 10 * time.Second
	defaultMiniKeepAliveTimeRate       = 2
)

type Config struct {
	*Options

	id                       string
	domain                   string
	addr                     string
	tls                      *tls.Config
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
