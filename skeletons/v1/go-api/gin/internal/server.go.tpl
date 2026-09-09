package internal

import (
	"{{ .App.ModuleName }}/internal/config"
	"{{ .App.ModuleName }}/internal/store"
	"{{ .App.ModuleName }}/internal/store/mysql"
	"{{ .App.ModuleName }}/internal/store/sqlite"
	genericapiserver "{{ .App.ModuleName }}/pkg/server"
	"{{ .App.ModuleName }}/pkg/xlog"
)

type ApiServer struct {
	genericAPIServer *genericapiserver.GenericAPIServer
}

func CreateAPIServer(cfg *config.Config) (*ApiServer, error) {
	genericCfg, err := cfg.BuildGenericServerConfig()
	if err != nil {
		return nil, err
	}
	genericServer, err := genericapiserver.New(genericCfg)
	if err != nil {
		return nil, err
	}

	var storeFactory store.Factory
	if cfg.Store == "mysql" {
		storeFactory, err = mysql.GetMySQLFactoryOr(cfg.MySQLOptions)
	} else {
		storeFactory, err = sqlite.GetSQLiteFactoryOr("file::memory:?cache=shared")
	}
	if err != nil {
		return nil, err
	}
	store.SetClient(storeFactory)

	return &ApiServer{genericAPIServer: genericServer}, nil
}

func (s *ApiServer) PreRun() *ApiServer {
	initRouter(s.genericAPIServer.Engine)
	return s
}

func (s *ApiServer) Run() error {
	defer func() {
		if storeFactory := store.Client(); storeFactory != nil {
			_ = storeFactory.Close()
		}
		s.genericAPIServer.Close()
	}()

	xlog.Infof("Starting API server...")
	return s.genericAPIServer.Run()
}
