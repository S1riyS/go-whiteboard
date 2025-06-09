package whiteboard

import (
	"context"
	"log/slog"

	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/entity"
	"github.com/S1riyS/go-whiteboard/whiteboard-service/pkg/database/postgresql"
)

type IRepository interface {
	Create(ctx context.Context, whiteboard *entity.Whiteboard) (int, error)
	GetByID(ctx context.Context, id int) (*entity.Whiteboard, error)
	Update(ctx context.Context, whiteboard *entity.Whiteboard) (*entity.Whiteboard, error)
	Delete(ctx context.Context, id int) error
}

type Repository struct {
	logger   *slog.Logger
	dbClient postgresql.Client
}

func NewRepository(logger *slog.Logger, dbClient postgresql.Client) *Repository {
	return &Repository{
		logger:   logger,
		dbClient: dbClient,
	}
}
