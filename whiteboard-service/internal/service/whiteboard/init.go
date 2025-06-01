package whiteboard

import (
	"log/slog"

	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/repository/whiteboard"
)

type Service struct {
	logger *slog.Logger
	repo   whiteboard.IRepository
}

func NewService(logger *slog.Logger, repo whiteboard.IRepository) *Service {
	return &Service{
		logger: logger,
		repo:   repo,
	}
}
