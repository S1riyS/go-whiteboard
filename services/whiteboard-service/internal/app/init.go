package app

import (
	"context"
	"log/slog"

	grpcServer "github.com/S1riyS/go-whiteboard/whiteboard-service/internal/app/grpc"
	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/config"
	whiteboardRepo "github.com/S1riyS/go-whiteboard/whiteboard-service/internal/repository/whiteboard"
	whiteboardSvc "github.com/S1riyS/go-whiteboard/whiteboard-service/internal/service/whiteboard"
	"github.com/S1riyS/go-whiteboard/whiteboard-service/pkg/database/postgresql"
)

type app struct {
	logger     *slog.Logger
	cfg        *config.Config
	grpcServer *grpcServer.Server
}

func New(ctx context.Context, logger *slog.Logger, cfg *config.Config) *app {
	dbClient := postgresql.MustNewClient(ctx, logger, cfg.Database)
	repo := whiteboardRepo.NewRepository(logger, dbClient)
	svc := whiteboardSvc.NewService(logger, repo)

	return &app{
		logger:     logger,
		cfg:        cfg,
		grpcServer: grpcServer.New(logger, cfg.GRPC, svc),
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
