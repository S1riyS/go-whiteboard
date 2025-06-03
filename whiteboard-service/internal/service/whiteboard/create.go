package whiteboardService

import (
	"context"

	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/entity"
	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/service"
)

func (s *Service) Create(ctx context.Context, whiteboard *entity.Whiteboard) (int, error) {
	id, err := s.repo.Create(ctx, whiteboard)
	if err != nil {
		// TODO: more thorough error handling
		return 0, service.ErrInternal
	}
	return id, nil
}
