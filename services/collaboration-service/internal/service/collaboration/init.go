package collaborationservice

import (
	"log/slog"

	"github.com/S1riyS/go-whiteboard/collaboration-service/internal/domain/repository"
)

type Service struct {
	logger *slog.Logger
	repo   repository.ICollaborationRepo
}

func NewService(logger *slog.Logger, repo repository.ICollaborationRepo) *Service {
	return &Service{
		logger: logger,
		repo:   repo,
	}
}
