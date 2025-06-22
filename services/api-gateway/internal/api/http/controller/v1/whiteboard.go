package v1

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/S1riyS/go-whiteboard/api-gateway/internal/api"
	"github.com/S1riyS/go-whiteboard/api-gateway/internal/dto/request"
	"github.com/S1riyS/go-whiteboard/api-gateway/internal/dto/response"
	"github.com/S1riyS/go-whiteboard/api-gateway/pkg/logger/slogext"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WhiteboardClient interface {
	CreateWhiteboard(ctx context.Context, req *request.CreateWhiteboardRequest) (*response.WhiteboardResponse, error)
	GetWhiteboard(ctx context.Context, id uuid.UUID) (*response.WhiteboardResponse, error)
	UpdateWhiteboard(ctx context.Context, id uuid.UUID, req *request.UpdateWhiteboardRequest) (*response.WhiteboardResponse, error)
	DeleteWhiteboard(ctx context.Context, id uuid.UUID) error
}

type WhiteboardController struct {
	logger *slog.Logger
	client WhiteboardClient
}

func NewWhiteboardController(logger *slog.Logger, client WhiteboardClient) *WhiteboardController {
	return &WhiteboardController{
		logger: logger,
		client: client,
	}
}

func (c *WhiteboardController) Create(ctx *gin.Context) {
	const mark = "whiteboardController.Create"
	logger := c.logger.With(slog.String("mark", mark))

	// Bind request
	var req request.CreateWhiteboardRequest
	err := ctx.Bind(&req)
	if err != nil {
		logger.Debug("Failed to bind request", slogext.Err(err))
		api.NewUnprocessableEntityError().WriteToContext(ctx)
		return
	}

	response, err := c.client.CreateWhiteboard(ctx.Request.Context(), &req)
	if err != nil {
		logger.Error("Failed to create whiteboard", slog.Any("response", response), slogext.Err(err))
		api.WriteErrorToContext(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (c *WhiteboardController) GetOne(ctx *gin.Context) {
	const mark = "whiteboardController.GetOne"
	logger := c.logger.With(slog.String("mark", mark))

	// Retrieve ID
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		logger.Warn("Invalid ID format", slog.String("id", ctx.Param("id")), slogext.Err(err))
		api.NewBadRequestError("Invalid ID format").WriteToContext(ctx)
		return
	}

	response, err := c.client.GetWhiteboard(ctx.Request.Context(), id)
	if err != nil {
		logger.Error("Failed to get whiteboard", slog.String("id", ctx.Param("id")), slogext.Err(err))
		api.WriteErrorToContext(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *WhiteboardController) Update(ctx *gin.Context) {
	const mark = "whiteboardController.Update"
	logger := c.logger.With(slog.String("mark", mark))

	// Retrieve ID
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		logger.Warn("Invalid ID format", slog.String("id", ctx.Param("id")), slogext.Err(err))
		api.NewBadRequestError("Invalid ID format").WriteToContext(ctx)
		return
	}

	// Bind request
	var req request.UpdateWhiteboardRequest
	err = ctx.Bind(&req)
	if err != nil {
		logger.Warn("Failed to bind request", slog.Any("request", req), slogext.Err(err))
		api.NewUnprocessableEntityError().WriteToContext(ctx)
		return
	}

	response, err := c.client.UpdateWhiteboard(ctx.Request.Context(), id, &req)
	if err != nil {
		logger.Error("Failed to update whiteboard",
			slog.String("whiteboard_id", id.String()),
			slog.Any("request", req),
			slog.Any("response", response),
			slogext.Err(err),
		)
		api.WriteErrorToContext(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *WhiteboardController) Delete(ctx *gin.Context) {
	const mark = "whiteboardController.Delete"
	logger := c.logger.With(slog.String("mark", mark))

	// Retrieve ID
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		logger.Warn("Invalid ID format", slog.String("id", ctx.Param("id")), slogext.Err(err))
		api.NewBadRequestError("Invalid ID format").WriteToContext(ctx)
		return
	}

	err = c.client.DeleteWhiteboard(ctx.Request.Context(), id)
	if err != nil {
		logger.Error("Failed to delete whiteboard", slog.String("id", ctx.Param("id")), slogext.Err(err))
		api.WriteErrorToContext(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}
