package app

import (
	"context"
	"log/slog"

	grpcServer "github.com/S1riyS/go-whiteboard/collaboration-service/internal/app/grpc"
	"github.com/S1riyS/go-whiteboard/collaboration-service/internal/config"
	collaborationdb "github.com/S1riyS/go-whiteboard/collaboration-service/internal/infrastructure/redis/collaboration"
	collaborationservice "github.com/S1riyS/go-whiteboard/collaboration-service/internal/service/collaboration"
)

type app struct {
	logger     *slog.Logger
	cfg        *config.Config
	grpcServer *grpcServer.Server
}

func New(ctx context.Context, logger *slog.Logger, cfg *config.Config) *app {
	cacheClient := struct{}{} // TODO: replace with redis
	repo := collaborationdb.NewRepository(logger, cacheClient)
	svc := collaborationservice.NewService(logger, repo)

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
