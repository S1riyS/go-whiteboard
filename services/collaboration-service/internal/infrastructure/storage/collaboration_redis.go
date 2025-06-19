package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/S1riyS/go-whiteboard/collaboration-service/internal/domain/models"
	"github.com/S1riyS/go-whiteboard/collaboration-service/internal/domain/repository"
	"github.com/S1riyS/go-whiteboard/collaboration-service/pkg/logger/slogext"
	"github.com/redis/go-redis/v9"
)

type CollaborationRedisStorage struct {
	logger *slog.Logger
	client *redis.Client
}

func NewCollaborationRedisStorage(logger *slog.Logger, client *redis.Client) *CollaborationRedisStorage {
	return &CollaborationRedisStorage{
		logger: logger,
		client: client,
	}
}

// Ensure Repository implements ICollaborationRepo
var _ repository.ICollaborationRepo = (*CollaborationRedisStorage)(nil)

// Save saves the given drawing with the whiteboard elements index.
func (r *CollaborationRedisStorage) Save(ctx context.Context, drawing *models.Drawing) error {
	const mark = "storage.collaboration.Save"
	logger := r.logger.With(slog.String("mark", mark))

	data, err := json.Marshal(drawing)
	if err != nil {
		logger.Error("Failed to marshal drawing", slogext.Err(err))
		return fmt.Errorf("marshal error: %w", err)
	}

	// Save drawing
	key := fmt.Sprintf("whiteboard:%s:element:%s", drawing.WhiteboardID, drawing.ID)
	logger.Debug("Saving drawing", slog.String("key", key))
	if err := r.client.Set(ctx, key, data, 0).Err(); err != nil {
		logger.Error("Failed to save drawing", slogext.Err(err))
		return fmt.Errorf("redis set error: %w", err)
	}

	// Add to index
	zkey := fmt.Sprintf("whiteboard:%s:elements", drawing.WhiteboardID)
	logger.Debug("Adding drawing to index", slog.String("key", zkey))
	if err := r.client.ZAdd(ctx, zkey, redis.Z{
		Score:  float64(drawing.UpdatedAt.UnixNano()),
		Member: drawing.ID,
	}).Err(); err != nil {
		logger.Error("Failed to add drawing to index", slogext.Err(err))
		return fmt.Errorf("redis zadd error: %w", err)
	}

	logger.Debug("Successfully saved drawing", slog.String("key", key))
	return nil
}

// GetOne retrieves a drawing from the Redis storage by its drawingID and whiteboardID.
func (r *CollaborationRedisStorage) GetOne(ctx context.Context, drawingID string, whiteboardID string) (*models.Drawing, error) {
	// TODO: Implement actual database logic here
	return nil, nil
}

// Delete deletes the given drawing from Redis storage.
func (r *CollaborationRedisStorage) Delete(ctx context.Context, drawingID string, whiteboardID string) error {
	const mark = "storage.collaboration.Delete"
	logger := r.logger.With(
		slog.String("mark", mark),
		slog.String("drawingID", drawingID),
		slog.String("whiteboardID", whiteboardID),
	)

	key := fmt.Sprintf("whiteboard:%s:element:%s", whiteboardID, drawingID)

	// Check existance
	exists, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		logger.Error("Failed to check existance of drawing", slogext.Err(err))
		return fmt.Errorf("redis exists error: %w", err)
	}

	// If drawing not found
	if exists == 0 {
		logger.Debug("Drawing not found", slog.String("key", key))
		return repository.ErrDrawingNotFound
	}

	// Delete drawing
	if err := r.client.Del(ctx, key).Err(); err != nil {
		logger.Error("Failed to delete drawing", slogext.Err(err))
		return fmt.Errorf("redis del error: %w", err)
	}

	// Delete from index
	zkey := fmt.Sprintf("whiteboard:%s:elements", whiteboardID)
	if err := r.client.ZRem(ctx, zkey, drawingID).Err(); err != nil {
		logger.Error("Failed to delete drawing from index", slogext.Err(err))
		return fmt.Errorf("redis zrem error: %w", err)
	}

	return nil
}

// DeleteByWhiteboard deletes all drawings associated with the given whiteboardID.
func (r *CollaborationRedisStorage) DeleteByWhiteboard(ctx context.Context, whiteboardID string) error {
	const mark = "storage.collaboration.DeleteByWhiteboard"
	logger := r.logger.With(slog.String("mark", mark), slog.String("whiteboardID", whiteboardID))

	// Get elements from index
	zkey := fmt.Sprintf("whiteboard:%s:elements", whiteboardID)
	elements, err := r.client.ZRange(ctx, zkey, 0, -1).Result()
	if err != nil {
		logger.Error("Failed to get elements from index", slogext.Err(err))
		return err
	}

	// Delete elements using pipeline
	pipe := r.client.Pipeline()
	for _, elementID := range elements {
		elementKey := "whiteboard:" + whiteboardID + ":element:" + elementID
		pipe.Unlink(ctx, elementKey)
	}
	// Delete index itself
	pipe.Unlink(ctx, zkey)

	// Execute pipeline
	_, err = pipe.Exec(ctx)
	if err != nil {
		logger.Error("Failed to delete whiteboard elements", slogext.Err(err))
		return err
	}

	logger.Info("Successfully deleted whiteboard elements", slog.Int("count", len(elements)))
	return nil
}
