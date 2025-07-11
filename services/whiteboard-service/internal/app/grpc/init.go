package grpcserver

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/config"
	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/handler/grpc/interceptors"
	whiteboardGRPC "github.com/S1riyS/go-whiteboard/whiteboard-service/internal/handler/grpc/whiteboard"
	whiteboardService "github.com/S1riyS/go-whiteboard/whiteboard-service/internal/service/whiteboard"
	"google.golang.org/grpc"
)

// Server is an actual gRPC server that wraps underlying `whiteboard_grpc.Server`.
// It is responsible for starting and stopping gRPC server, applying interceptors.
type Server struct {
	logger     *slog.Logger
	cfg        config.GRPCConfig
	gRPCServer *grpc.Server
}

func New(logger *slog.Logger, cfg config.GRPCConfig, whiteboardService *whiteboardService.Service) *Server {
	// Create gRPC gRPCServer
	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		interceptors.Timeout(logger, cfg.Timeout),
		interceptors.Recovery(logger),
		interceptors.Logging(logger),
	))

	whiteboardGRPC.Register(logger, gRPCServer, whiteboardService)

	return &Server{
		logger:     logger,
		gRPCServer: gRPCServer,
		cfg:        cfg,
	}
}

// Run runs gRPC server.
func (s *Server) Run() error {
	const mark = "grpcserver.Run"

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
	const mark = "grpcserver.Stop"

	logger := a.logger.With(slog.String("mark", mark))

	a.gRPCServer.GracefulStop()
	logger.Info("stopping gRPC server", slog.Int("port", a.cfg.Port))
}
