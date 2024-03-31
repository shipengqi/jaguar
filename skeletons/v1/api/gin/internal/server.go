package internal

import (
	"github.com/shipengqi/log"

	"github.com/jaguar/apiskeleton/internal/config"
	"github.com/jaguar/apiskeleton/internal/store/mysql"
	genericapiserver "github.com/jaguar/apiskeleton/pkg/server"
	"github.com/jaguar/apiskeleton/pkg/shutdown"
	"github.com/jaguar/apiskeleton/pkg/shutdown/managers"
)

type ApiServer struct {
	gs               *shutdown.GracefulShutdown
	genericAPIServer *genericapiserver.GenericAPIServer
}

func CreateAPIServer(cfg *config.Config) (*ApiServer, error) {
	gs := shutdown.New()
	gs.AddShutdownManager(managers.NewPosixSignalManager())

	genericCfg, err := cfg.BuildGenericServerConfig()
	if err != nil {
		return nil, err
	}
	genericServer, err := genericapiserver.New(genericCfg)
	if err != nil {
		return nil, err
	}
	_, err = mysql.GetMySQLFactoryOr(cfg.MySQLOptions)
	if err != nil {
		return nil, err
	}

	server := &ApiServer{
		gs:               gs,
		genericAPIServer: genericServer,
	}

	return server, nil
}

func (s *ApiServer) PreRun() *ApiServer {
	initRouter(s.genericAPIServer.Engine)

	s.gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
		mysqlStore, _ := mysql.GetMySQLFactoryOr(nil)
		if mysqlStore != nil {
			_ = mysqlStore.Close()
		}

		s.genericAPIServer.Close()

		return nil
	}))

	return s
}

func (s *ApiServer) Run() error {
	// start shutdown managers
	if err := s.gs.Start(); err != nil {
		log.Fatalf("start shutdown manager failed: %s", err.Error())
	}

	return s.genericAPIServer.Run()
}
