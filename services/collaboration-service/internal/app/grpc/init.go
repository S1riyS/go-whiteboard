package appgrpc

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/S1riyS/go-whiteboard/collaboration-service/internal/config"
	"github.com/S1riyS/go-whiteboard/collaboration-service/internal/service"
	"github.com/S1riyS/go-whiteboard/collaboration-service/internal/transport/grpc/handler"
	"github.com/S1riyS/go-whiteboard/collaboration-service/internal/transport/grpc/interceptors"
	"google.golang.org/grpc"
)

// Server is an actual gRPC server that wraps underlying `whiteboard_grpc.Server`.
// It is responsible for starting and stopping gRPC server, applying interceptors.
type Server struct {
	logger     *slog.Logger
	cfg        config.GRPCConfig
	gRPCServer *grpc.Server
}

func New(logger *slog.Logger, cfg config.GRPCConfig, collaborationservice *service.CollaborationService) *Server {
	// Create gRPC gRPCServer
	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		interceptors.Timeout(logger, cfg.Timeout),
		interceptors.Recovery(logger),
		interceptors.Logging(logger),
	))

	handler.RegisterCollaborationHandler(logger, gRPCServer, collaborationservice)

	return &Server{
		logger:     logger,
		gRPCServer: gRPCServer,
		cfg:        cfg,
	}
}

// Run runs gRPC server.
func (s *Server) Run() error {
	const mark = "appgrpc.Run"

	localLogger := s.logger.With(slog.String("mark", mark))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.Port))
	if err != nil {
		return fmt.Errorf("%s: %w", mark, err)
	}

	localLogger.Info("gRPC server started", slog.String("addr", listener.Addr().String()))

	if err := s.gRPCServer.Serve(listener); err != nil {
		return fmt.Errorf("%s: %w", mark, err)
	}

	return nil
}

// Stop stops gRPC server.
func (a *Server) Stop() {
	const mark = "appgrpc.Stop"

	logger := a.logger.With(slog.String("mark", mark))

	a.gRPCServer.GracefulStop()
	logger.Info("stopping gRPC server", slog.Int("port", a.cfg.Port))
}
