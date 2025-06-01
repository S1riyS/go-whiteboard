package whiteboard

import (
	"context"

	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/entity"
)

type IRepository interface {
	Create(ctx context.Context, whiteboard *entity.Whiteboard) (int, error)
	GetByID(ctx context.Context, id int) (*entity.Whiteboard, error)
	Update(ctx context.Context, whiteboard *entity.Whiteboard) (*entity.Whiteboard, error)
	Delete(ctx context.Context, id int) error
}

type Repository struct {
	// TODO: replace with actual DB
	db []*entity.Whiteboard
}

func NewRepository() *Repository {
	return &Repository{
		db: make([]*entity.Whiteboard, 0),
	}
}
