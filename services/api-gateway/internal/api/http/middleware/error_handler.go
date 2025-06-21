package middlewares

import (
	"errors"
	"net/http"
	"time"

	"github.com/S1riyS/go-whiteboard/api-gateway/internal/api"
	"github.com/S1riyS/go-whiteboard/api-gateway/internal/dto/response"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Process the request
		c.Next()

		// Check if there's an error in the context
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			// Check if it's an ApiError
			var apiErr *api.ApiError
			if errors.As(err, &apiErr) {
				c.JSON(apiErr.Code, apiErrorToResponse(apiErr))
				return
			}

			// Internal error by default
			c.JSON(http.StatusInternalServerError, apiErrorToResponse(api.NewInternalError()))
			return
		}
	}
}

func apiErrorToResponse(err *api.ApiError) response.ApiErrorResponse {
	return response.ApiErrorResponse{
		Title:     err.Title,
		Details:   err.Details,
		Timestamp: time.Now().UTC(),
	}
}
