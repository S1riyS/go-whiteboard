package app

import (
	"context"
	"log/slog"

	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/app/grpc_server"
)

type app struct {
	logger     *slog.Logger
	grpcServer *grpc_server.GrpcServer
}

func New(ctx context.Context, logger *slog.Logger) *app {
	return &app{
		logger:     logger,
		grpcServer: grpc_server.New(),
	}
}

func (a *app) MustRun() {
	// TODO: implement me
	panic("implement me")
}

func (a *app) Stop() {
	// TODO: implement me
	panic("implement me")
}
