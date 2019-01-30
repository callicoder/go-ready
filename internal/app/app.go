package app

import (
	"github.com/callicoder/go-ready/internal/config"
	"github.com/callicoder/go-ready/pkg/logger"
)

type App struct {
	Srv     *Server
	Deps    *Dependencies
	GrpcSrv *GrpcServer
}

func New(configFile string) (*App, error) {
	config, err := config.Load(configFile)
	if err != nil {
		return nil, err
	}

	logger.InitLogger(config.Logging.Level)

	deps, err := NewDependencies(config)
	if err != nil {
		return nil, err
	}

	router := NewRouter(config, deps)
	server := NewServer(config.Server, router)
	grpcServer := NewGrpcServer(config.Grpc)

	app := &App{
		Srv:     server,
		GrpcSrv: grpcServer,
		Deps:    deps,
	}

	return app, nil
}

func (a *App) Start() {
	a.Srv.Start()
	a.GrpcSrv.Start()
}

func (a *App) Shutdown() {
	a.Srv.Shutdown()
	a.GrpcSrv.Shutdown()
}
