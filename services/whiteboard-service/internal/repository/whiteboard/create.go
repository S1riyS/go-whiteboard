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

// Create creates a new whiteboard, returns the id of the created whiteboard and sets the id to the struct
func (r *Repository) Create(ctx context.Context, whiteboard *entity.Whiteboard) (int, error) {
	const mark = "repository.whiteboard.Create"

	logger := r.logger.With(slog.String("mark", mark))

	queryBuilder := squirrel.
		Insert(TableName).
		PlaceholderFormat(squirrel.Dollar).
		Columns(TitleColumn, DescriptionColumn).
		Values(whiteboard.Title, whiteboard.Description).
		Suffix(ReturningID)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		logger.Error("Failed to build insert query", slogext.Err(err))
		return 0, fmt.Errorf("failed to build insert query: %v", err)
	}

	logger.Debug("Executing insert query", slog.String("query", query), slog.Any("args", args))

	err = r.dbClient.QueryRow(ctx, query, args...).Scan(&whiteboard.ID)
	if err != nil {
		// Handle database errors
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			logger.Error("Failed to create whiteboard due to database error", slogext.Err(pgErr))
			return 0, repository.ErrInternal
		}

		logger.Error("Failed to create whiteboard due to unexpected error", slogext.Err(err))
		return 0, fmt.Errorf("failed to create whiteboard: %v", err)
	}

	return whiteboard.ID, nil
}
