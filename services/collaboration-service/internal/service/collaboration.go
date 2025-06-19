package service

import (
	"context"
	"log/slog"

	"github.com/S1riyS/go-whiteboard/collaboration-service/internal/domain/models"
	"github.com/S1riyS/go-whiteboard/collaboration-service/internal/domain/repository"
)

type CollaborationService struct {
	logger *slog.Logger
	repo   repository.ICollaborationRepo
}

func NewCollaborationService(logger *slog.Logger, repo repository.ICollaborationRepo) *CollaborationService {
	return &CollaborationService{
		logger: logger,
		repo:   repo,
	}
}

func (svc *CollaborationService) Draw(ctx context.Context, drawing *models.Drawing) (int, error) {
	// TODO: implement
	return 0, nil
}

func (svc *CollaborationService) Delete(ctx context.Context, drawingId int, whiteboardId int) error {
	// TODO: implement
	return nil
}
