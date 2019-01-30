package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/callicoder/go-ready/internal/config"
	"github.com/callicoder/go-ready/pkg/logger"
)

type Server struct {
	httpServer *http.Server
	config     config.ServerConfig
}

func NewServer(serverConfig config.ServerConfig, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Handler:      handler,
			ReadTimeout:  time.Duration(serverConfig.ReadTimeoutSec) * time.Second,
			WriteTimeout: time.Duration(serverConfig.WriteTimeoutSec) * time.Second,
			Addr:         fmt.Sprintf("0.0.0.0:%d", serverConfig.Port),
		},
		config: serverConfig,
	}
}

func (s *Server) Start() {
	go func() {
		logger.Infof("Starting http server on port %v", s.config.Port)
		err := s.httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Failed to start http server %v", err)
			return
		}
	}()
}

func (s *Server) Shutdown() {
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.config.GracefulShutdownTimeoutSec)*time.Second)
	defer cancel()
	s.httpServer.Shutdown(ctx)
	logger.Info("Shutting down http server.")
}
