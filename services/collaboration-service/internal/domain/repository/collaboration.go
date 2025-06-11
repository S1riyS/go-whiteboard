package repository

import "github.com/S1riyS/go-whiteboard/collaboration-service/internal/domain/models"

type ICollaborationRepo interface {
	Save(drawing *models.Drawing) error
	Delete(drawingId int, whiteboardId int) error
}
