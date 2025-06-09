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
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) GetByID(ctx context.Context, id int) (*entity.Whiteboard, error) {
	const mark = "repository.whiteboard.GetByID"

	logger := r.logger.With(slog.String("mark", mark), slog.Int("id", id))

	queryBuilder := squirrel.
		Select(IDColumn, TitleColumn, DescriptionColumn).
		From(TableName).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{IDColumn: id}).
		Limit(1)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		logger.Error("Failed to build select query", slogext.Err(err))
		return nil, fmt.Errorf("failed to build select query: %v", err)
	}

	logger.Debug("Executing select query", slog.String("query", query), slog.Any("args", args))

	var foundWhiteboard entity.Whiteboard
	err = r.dbClient.QueryRow(ctx, query, args...).Scan(&foundWhiteboard.ID, &foundWhiteboard.Title, &foundWhiteboard.Description)
	if err != nil {
		// Handle database errors
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Debug("Whiteboard not found", slogext.Err(err))
			return nil, repository.ErrNotFound
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			logger.Error("Failed to get whiteboard due to database error", slogext.Err(pgErr))
			return nil, repository.ErrInternal
		}

		logger.Error("Failed to get whiteboard due to unexpected error", slogext.Err(err))
		return nil, fmt.Errorf("failed to get whiteboard: %v", err)
	}

	return &foundWhiteboard, nil
}
