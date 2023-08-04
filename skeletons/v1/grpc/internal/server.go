package internal

import (
	"github.com/jaguar/grpcskeleton/internal/config"
	v1 "github.com/jaguar/grpcskeleton/internal/service/v1"
	pb "github.com/jaguar/grpcskeleton/pkg/api/proto/v1"
	"github.com/jaguar/grpcskeleton/pkg/rpcsrv"
)

type GRPCServer struct {
	grpcSrv *rpcsrv.Server
}

func CreateGRPCServer(cfg *config.Config) (*GRPCServer, error) {
	srv := rpcsrv.New(cfg.GRPCOptions)
	return &GRPCServer{grpcSrv: srv}, nil
}

func (s *GRPCServer) PreRun() *GRPCServer {
	pb.RegisterUserServer(s.grpcSrv.GRPCServer(), &v1.UserService{})
	return s
}

func (s *GRPCServer) Run() error {
	return s.grpcSrv.Run()
}
