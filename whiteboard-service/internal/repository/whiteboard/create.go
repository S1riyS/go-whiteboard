package whiteboard

import (
	"context"

	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/entity"
)

func (r *Repository) Create(ctx context.Context, whiteboard *entity.Whiteboard) (int, error) {
	return 1, nil
}
