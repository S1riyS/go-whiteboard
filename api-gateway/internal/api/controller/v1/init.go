package v1

import "github.com/gin-gonic/gin"

func SetupControllers(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		whiteboardController := NewWhiteboardController()
		v1.GET("/whiteboard", whiteboardController.GetOne)
	}
}
