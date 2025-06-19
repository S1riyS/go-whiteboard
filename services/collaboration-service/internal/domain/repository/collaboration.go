package repository

import (
	"context"

	"github.com/S1riyS/go-whiteboard/collaboration-service/internal/domain/models"
)

type ICollaborationRepo interface {
	Save(ctx context.Context, drawing *models.Drawing) error
	GetOne(ctx context.Context, drawingID string, whiteboardID string) (*models.Drawing, error)
	Delete(ctx context.Context, drawingID string, whiteboardID string) error
	DeleteByWhiteboard(ctx context.Context, whiteboardID string) error
}
