package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/S1riyS/go-whiteboard/collaboration-service/internal/domain/models"
	"github.com/S1riyS/go-whiteboard/collaboration-service/internal/domain/repository"
	"github.com/S1riyS/go-whiteboard/collaboration-service/pkg/logger/slogext"
	"github.com/google/uuid"
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

func (svc *CollaborationService) HandleOperation(ctx context.Context, op models.Operation) (*models.Drawing, error) {
	switch op.Type {
	case models.OperationCreate:
		return svc.draw(ctx, op)
	case models.OperationUpdate:
		return svc.update(ctx, op)
	case models.OperationDelete:
		return svc.delete(ctx, op)
	}

	return nil, fmt.Errorf("unknown operation type: %s", op.Type)
}

func (svc *CollaborationService) draw(ctx context.Context, op models.Operation) (*models.Drawing, error) {
	const mark = "service.collaboration.Draw"
	logger := svc.logger.With(slog.String("mark", mark))

	logger.Debug("Handling draw request", slog.Any("operaion", op))

	// Convert to drawing model
	drawing := &models.Drawing{
		DrawingMetadata: models.DrawingMetadata{
			ID:           uuid.New().String(),
			WhiteboardID: op.WhiteboardID,
			CreatedAt:    op.Timestamp,
			UpdatedAt:    op.Timestamp,
			Version:      1,
		},
		Type:    op.ElementType,
		Payload: op.Payload,
	}

	err := svc.repo.Save(ctx, drawing)
	if err != nil {
		logger.Error("Failed to create drawing", slogext.Err(err))
		return nil, err
	}

	return drawing, nil

}

func (svc *CollaborationService) update(ctx context.Context, op models.Operation) (*models.Drawing, error) {
	// TODO: implement
	return nil, nil
}

func (svc *CollaborationService) delete(ctx context.Context, op models.Operation) (*models.Drawing, error) {
	const mark = "service.collaboration.Delete"
	logger := svc.logger.With(slog.String("mark", mark))

	logger.Debug("Handling delete request", slog.Any("operaion", op))

	err := svc.repo.Delete(ctx, op.DrawingID, op.WhiteboardID)
	if errors.Is(err, repository.ErrDrawingNotFound) {
		logger.Debug("Drawing not found", slog.String("id", op.DrawingID))
		return nil, ErrDrawingNotFound
	}
	if err != nil {
		logger.Error("Failed to delete drawing", slogext.Err(err))
		return nil, err
	}
	return nil, err
}
