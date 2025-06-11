package collaborationdb

import (
	"log/slog"

	"github.com/S1riyS/go-whiteboard/collaboration-service/internal/domain/models"
	"github.com/S1riyS/go-whiteboard/collaboration-service/internal/domain/repository"
)

type Repository struct {
	logger      *slog.Logger
	cacheClient any // TODO: change to actual redis client
}

func NewRepository(logger *slog.Logger, cacheClient any) *Repository {
	return &Repository{
		logger:      logger,
		cacheClient: cacheClient,
	}
}

// Ensure Repository implements ICollaborationRepo
var _ repository.ICollaborationRepo = (*Repository)(nil)

func (r *Repository) Save(drawing *models.Drawing) error {
	// TODO: Implement actual database logic here
	return nil
}

func (r *Repository) Delete(drawingId int, whiteboardId int) error {
	// TODO: Implement actual database logic here
	return nil
}
