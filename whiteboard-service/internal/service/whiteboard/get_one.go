package whiteboardService

import (
	"context"
	"errors"

	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/entity"
	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/repository"
	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/service"
)

func (s *Service) GetOne(ctx context.Context, id int) (*entity.Whiteboard, error) {
	whiteboard, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, service.ErrNotFound
		}
		if errors.Is(err, repository.ErrInternal) {
			return nil, service.ErrInternal
		}
		return nil, err
	}
	return whiteboard, nil
}
