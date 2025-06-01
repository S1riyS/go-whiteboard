package grpcserver

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/config"
	whiteboardGRPC "github.com/S1riyS/go-whiteboard/whiteboard-service/internal/handlers/grpc/whiteboard"
	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/service/whiteboard"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server is an actual gRPC server that wraps underlying `whiteboard_grpc.Server`.
// It is responsible for starting and stopping gRPC server.
type Server struct {
	logger     *slog.Logger
	cfg        config.GRPCConfig
	gRPCServer *grpc.Server
}

func New(logger *slog.Logger, cfg config.GRPCConfig, whiteboardService *whiteboard.Service) *Server {
	// Interceptors
	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(
			logging.PayloadReceived,
			logging.PayloadSent,
		),
	}
	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p any) (err error) {
			logger.Error("Recovered from panic", slog.Any("panic", p))
			return status.Errorf(codes.Internal, "internal error")
		}),
	}

	// Create gRPC gRPCServer
	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(recoveryOpts...),
		logging.UnaryServerInterceptor(interceptorLogger(logger), loggingOpts...),
	))

	whiteboardGRPC.Register(gRPCServer, whiteboardService)

	return &Server{
		logger:     logger,
		gRPCServer: gRPCServer,
		cfg:        cfg,
	}
}

// interceptorLogger adapts slog logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func interceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

// Run runs gRPC server.
func (s *Server) Run() error {
	const mark = "grpcserver.Run"

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.Port))
	if err != nil {
		return fmt.Errorf("%s: %w", mark, err)
	}

	s.logger.Info("grpc server started", slog.String("addr", listener.Addr().String()))

	if err := s.gRPCServer.Serve(listener); err != nil {
		return fmt.Errorf("%s: %w", mark, err)
	}

	return nil
}

// Stop stops gRPC server.
func (a *Server) Stop() {
	const mark = "grpcserver.Stop"

	a.logger.With(slog.String("mark", mark)).
		Info("stopping gRPC server", slog.Int("port", a.cfg.Port))

	a.gRPCServer.GracefulStop()
}
