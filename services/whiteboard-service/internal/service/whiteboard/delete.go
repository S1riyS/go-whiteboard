package whiteboardService

import (
	"context"
	"errors"

	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/repository"
	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/service"
)

func (s *Service) Delete(ctx context.Context, id int) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return service.ErrNotFound
		}
		if errors.Is(err, repository.ErrInternal) {
			return service.ErrInternal
		}
		return err
	}

	return nil
}
