package v1

import (
	"net/http"

	"github.com/S1riyS/go-whiteboard/api-gateway/internal/dto/request"
	"github.com/gin-gonic/gin"
)

type WhiteboardController struct {
}

func NewWhiteboardController() *WhiteboardController {
	return &WhiteboardController{}
}

func (c *WhiteboardController) GetOne(ctx *gin.Context) {
	var req request.RegisterRequest
	_ = req

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}
