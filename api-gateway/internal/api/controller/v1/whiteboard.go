package v1

import (
	"net/http"

	"github.com/S1riyS/go-whiteboard/api-gateway/internal/dto/request"
	"github.com/gin-gonic/gin"
)

type whiteboardController struct {
}

func NewWhiteboardController() *whiteboardController {
	return &whiteboardController{}
}

func (c *whiteboardController) GetOne(ctx *gin.Context) {
	// TODO: retrieve ID

	// TODO: Implement get one

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}

func (c *whiteboardController) Create(ctx *gin.Context) {
	var req request.CreateWhiteboardRequest
	_ = req

	// TODO: Implement create

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}

func (c *whiteboardController) Update(ctx *gin.Context) {
	var req request.UpdateWhiteboardRequest
	_ = req

	// TODO: retrieve ID

	// TODO: Implement update

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}

func (c *whiteboardController) Delete(ctx *gin.Context) {
	// TODO: retrieve ID

	// TODO: Implement delete

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}
