package app

import (
	"context"
	"log/slog"

	appgrpc "github.com/S1riyS/go-whiteboard/collaboration-service/internal/app/grpc"
	"github.com/S1riyS/go-whiteboard/collaboration-service/internal/config"
	"github.com/S1riyS/go-whiteboard/collaboration-service/internal/infrastructure/storage"
	"github.com/S1riyS/go-whiteboard/collaboration-service/internal/service"
	"github.com/S1riyS/go-whiteboard/collaboration-service/pkg/cache/redis"
)

type app struct {
	logger     *slog.Logger
	cfg        *config.Config
	grpcServer *appgrpc.Server
}

func New(ctx context.Context, logger *slog.Logger, cfg *config.Config) *app {
	cacheClient := redis.MustNewClient(ctx, logger, cfg.Redis)
	collaborationStorage := storage.NewCollaborationRedisStorage(logger, cacheClient)
	collaborationSvc := service.NewCollaborationService(logger, collaborationStorage)

	return &app{
		logger:     logger,
		cfg:        cfg,
		grpcServer: appgrpc.New(logger, cfg.GRPC, collaborationSvc),
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
