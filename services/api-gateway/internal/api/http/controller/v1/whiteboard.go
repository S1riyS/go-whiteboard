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

type whiteboardController struct {
	logger *slog.Logger
	client WhiteboardClient
}

func NewWhiteboardController(logger *slog.Logger, client WhiteboardClient) *whiteboardController {
	return &whiteboardController{
		logger: logger,
		client: client,
	}
}

func (c *whiteboardController) Create(ctx *gin.Context) {
	const mark = "whiteboardController.Create"
	logger := c.logger.With(slog.String("mark", mark))

	// Bind request
	var req request.CreateWhiteboardRequest
	err := ctx.Bind(&req)
	if err != nil {
		logger.Warn("Failed to bind request", slogext.Err(err))
		// TODO: set error like api.BadRequest and move error formatting to middleware
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind request",
		})
		return
	}

	response, err := c.client.CreateWhiteboard(ctx.Request.Context(), &req)
	if err != nil {
		logger.Error("Failed to create whiteboard", slog.Any("response", response), slogext.Err(err))
		// TODO: set error like api.BadRequest and move error formatting to middleware
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create whiteboard",
		})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (c *whiteboardController) GetOne(ctx *gin.Context) {
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
		// TODO: set error like api.BadRequest and move error formatting to middleware
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get whiteboard",
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *whiteboardController) Update(ctx *gin.Context) {
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
		logger.Warn("Failed to bind request", slogext.Err(err))
		// TODO: set error like api.BadRequest and move error formatting to middleware
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind request",
		})
		return
	}

	response, err := c.client.UpdateWhiteboard(ctx.Request.Context(), id, &req)
	if err != nil {
		logger.Error("Failed to update whiteboard", slog.Any("response", response), slogext.Err(err))
		// TODO: set error like api.BadRequest and move error formatting to middleware
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update whiteboard",
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *whiteboardController) Delete(ctx *gin.Context) {
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
		// TODO: set error like api.BadRequest and move error formatting to middleware
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete whiteboard",
		})
		return
	}

	ctx.Status(http.StatusOK)
}
