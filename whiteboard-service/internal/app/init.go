package app

import (
	"context"
	"log/slog"

	grpcServer "github.com/S1riyS/go-whiteboard/whiteboard-service/internal/app/grpc"
	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/config"
	whiteboardRepo "github.com/S1riyS/go-whiteboard/whiteboard-service/internal/repository/whiteboard"
	whiteboardSvc "github.com/S1riyS/go-whiteboard/whiteboard-service/internal/service/whiteboard"
)

type app struct {
	logger     *slog.Logger
	cfg        *config.Config
	grpcServer *grpcServer.Server
}

func New(ctx context.Context, logger *slog.Logger, cfg *config.Config) *app {
	repo := whiteboardRepo.NewRepository()
	service := whiteboardSvc.NewService(logger, repo)

	return &app{
		logger:     logger,
		cfg:        cfg,
		grpcServer: grpcServer.New(logger, cfg.GRPC, service),
	}
}

func (a *app) MustRun() {
	if err := a.grpcServer.Run(); err != nil {
		panic(err)
	}
}

func (a *app) Stop() {
	a.grpcServer.Stop()
}
