package repository

import (
	"context"

	"github.com/S1riyS/go-whiteboard/collaboration-service/internal/domain/models"
)

type ICollaborationRepo interface {
	Save(ctx context.Context, drawing *models.Drawing) (int, error)
	Delete(ctx context.Context, drawingId int) error
}
