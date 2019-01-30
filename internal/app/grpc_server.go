package app

import (
	"fmt"
	"net"
	"time"

	"github.com/callicoder/go-ready/internal/config"
	"github.com/callicoder/go-ready/pkg/logger"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	server *grpc.Server
	config config.GrpcConfig
}

func NewGrpcServer(grpcConfig config.GrpcConfig) *GrpcServer {
	server := grpc.NewServer(
		grpc.ConnectionTimeout(time.Duration(grpcConfig.ConnectionTimeoutSec) * time.Second),
	)

	return &GrpcServer{
		server: server,
		config: grpcConfig,
	}
}

func (s *GrpcServer) Start() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.Port))
	if err != nil {
		logger.Fatalf("Failed to listen on grpc port %d: %v", s.config.Port, err)
		return
	}

	go func() {
		logger.Infof("Starting grpc server on port %v", s.config.Port)
		if err := s.server.Serve(lis); err != nil {
			logger.Fatalf("Failed to serve grpc: %v", err)
		}
	}()
}

func (s *GrpcServer) Shutdown() {
	stopped := make(chan struct{})

	go func() {
		s.server.GracefulStop()
		close(stopped)
	}()

	t := time.NewTimer(time.Duration(s.config.GracefulShutdownTimeoutSec) * time.Second)
	select {
	case <-t.C:
		s.server.Stop()
	case <-stopped:
		s.server.Stop()
	}

	logger.Info("Shutting down grpc server.")
}
