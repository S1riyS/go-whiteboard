package storage

import (
	"context"
	"log/slog"

	"github.com/S1riyS/go-whiteboard/collaboration-service/internal/domain/models"
	"github.com/S1riyS/go-whiteboard/collaboration-service/internal/domain/repository"
	"github.com/redis/go-redis/v9"
)

type CollaborationRedisStorage struct {
	logger      *slog.Logger
	cacheClient *redis.Client
}

func NewCollaborationRedisStorage(logger *slog.Logger, cacheClient *redis.Client) *CollaborationRedisStorage {
	return &CollaborationRedisStorage{
		logger:      logger,
		cacheClient: cacheClient,
	}
}

// Ensure Repository implements ICollaborationRepo
var _ repository.ICollaborationRepo = (*CollaborationRedisStorage)(nil)

func (r *CollaborationRedisStorage) Save(ctx context.Context, drawing *models.Drawing) (int, error) {
	// const mark = "repository.collaboration.Save"
	// logger := r.logger.With(slog.String("mark", mark))

	// drawingId := int(time.Now().UnixNano()) // TODO: Generate unique ID
	// key := fmt.Sprintf("whiteboard:%d:elements:%d", drawing.WhiteboardID, drawingId)

	// // Convert entire struct to JSON
	// jsonData, err := json.Marshal(drawing)
	// if err != nil {
	// 	return 0, err
	// }

	// // Store in Redis
	// err = r.cacheClient.Set(ctx, key, jsonData, 0).Err()
	// if err != nil {
	// 	return 0, err
	// }
	// logger.Debug("Saved drawing to Redis", slog.String("key", key))

	// // Add to whiteboard index (for lookup later)
	// indexKey := fmt.Sprintf("whiteboard:%d:element_keys", drawing.WhiteboardID)
	// err = r.cacheClient.SAdd(ctx, indexKey, key).Err()
	// if err != nil {
	// 	return 0, err
	// }

	// return drawingId, nil
	return 0, nil
}

func (r *CollaborationRedisStorage) Delete(ctx context.Context, drawingId int) error {
	// TODO: Implement actual database logic here
	return nil
}
