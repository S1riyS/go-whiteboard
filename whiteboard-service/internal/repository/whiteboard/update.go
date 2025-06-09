package whiteboard

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/entity"
	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/repository"
	"github.com/S1riyS/go-whiteboard/whiteboard-service/pkg/logger/slogext"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) Update(ctx context.Context, whiteboard *entity.Whiteboard) (*entity.Whiteboard, error) {
	const mark = "repository.whiteboard.Update"

	logger := r.logger.With(slog.String("mark", mark))

	queryBuilder := squirrel.
		Update(TableName).
		PlaceholderFormat(squirrel.Dollar).
		Set(TitleColumn, whiteboard.Title).
		Set(DescriptionColumn, whiteboard.Description).
		Where(squirrel.Eq{IDColumn: whiteboard.ID})

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		logger.Error("Failed to build update query", slogext.Err(err))
		return nil, fmt.Errorf("failed to build update query: %v", err)
	}

	logger.Debug("Executing update query", slog.String("query", query), slog.Any("args", args))

	res, err := r.dbClient.Exec(ctx, query, args...)
	if err != nil {
		// Handle database errors
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			logger.Error("Failed to update whiteboard due to database error", slogext.Err(pgErr))
			return nil, repository.ErrInternal
		}

		logger.Error("Failed to update whiteboard due to unexpected error", slogext.Err(err))
		return nil, fmt.Errorf("failed to create whiteboard: %v", err)
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		logger.Debug("Whiteboard not found", slog.Int("id", whiteboard.ID))
		return nil, repository.ErrNotFound
	}

	return whiteboard, nil
}
