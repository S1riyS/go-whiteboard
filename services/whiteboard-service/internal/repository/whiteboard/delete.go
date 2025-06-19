package whiteboard

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/repository"
	"github.com/S1riyS/go-whiteboard/whiteboard-service/pkg/logger/slogext"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) Delete(ctx context.Context, id string) error {
	const mark = "repository.whiteboard.Update"

	logger := r.logger.With(slog.String("mark", mark))

	queryBuilder := squirrel.
		Delete(TableName).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{IDColumn: id})

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		logger.Error("Failed to build delete query", slogext.Err(err))
		return fmt.Errorf("failed to build delete query: %v", err)
	}

	logger.Debug("Executing delete query", slog.String("query", query), slog.Any("args", args))

	res, err := r.dbClient.Exec(ctx, query, args...)
	if err != nil {
		// Handle database errors
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			logger.Error("Failed to update whiteboard due to database error", slogext.Err(pgErr))
			return repository.ErrInternal
		}

		logger.Error("Failed to update whiteboard due to unexpected error", slogext.Err(err))
		return fmt.Errorf("failed to create whiteboard: %v", err)
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		logger.Debug("Whiteboard not found", slog.String("id", id))
		return repository.ErrNotFound
	}

	return nil
}
