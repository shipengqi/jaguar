package rpcsrv

import (
	"net"

	"github.com/shipengqi/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	cfg *Config
	srv *grpc.Server
}

func New(opts *Options) *Server {
	grpc.EnableTracing = false

	cfg := CreateConfigFromOptions(opts)
	srv := grpc.NewServer(cfg.grpcopts...)

	reflection.Register(srv)

	return &Server{cfg: cfg, srv: srv}
}

func (s *Server) GRPCServer() *grpc.Server {
	return s.srv
}

func (s *Server) Run() error {
	listen, err := net.Listen("tcp", s.cfg.addr)
	if err != nil {
		log.Debugf("failed to listen: %s", err.Error())
		return err
	}

	log.Infof("start gRPC server at %s", s.cfg.addr)
	if err := s.srv.Serve(listen); err != nil {
		log.Debugf("failed to start gRPC server: %s", err.Error())
		return err
	}

	return nil
}

func (s *Server) Shutdown() error {
	log.Infof("gRPC server on %s stopped", s.cfg.addr)
	s.srv.GracefulStop()
	return nil
}
