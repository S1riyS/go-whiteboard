package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	TitleInternalError       = "Internal error"
	TitleUnprocessableEntity = "Unprocessable entity"
	TitleNotFound            = "Not found"
	TitleBadRequest          = "Bad request"
	TitleUnauthorized        = "Unauthorized"
	TitleValidationFailed    = "Validation failed"

	MessageInternalError       = "Something went wrong"
	MessageUnprocessableEntity = "Invalid request"
)

type ApiError struct {
	Code    int    `json:"code"`
	Title   string `json:"title"`
	Details string `json:"details"`
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("%s: %s (code %d)", e.Title, e.Details, e.Code)
}

func NewApiError(code int, title string, details string) *ApiError {
	return &ApiError{
		Code:    code,
		Title:   title,
		Details: details,
	}
}

func (e *ApiError) WriteToContext(ctx *gin.Context) {
	ctx.Error(e)
}

func NewInternalError() *ApiError {
	return &ApiError{
		Code:    http.StatusInternalServerError,
		Title:   TitleInternalError,
		Details: MessageInternalError,
	}
}

func NewUnprocessableEntityError() *ApiError {
	return &ApiError{
		Code:    http.StatusUnprocessableEntity,
		Title:   TitleUnprocessableEntity,
		Details: MessageUnprocessableEntity,
	}
}

func NewNotFoundError(details string) *ApiError {
	return &ApiError{
		Code:    http.StatusNotFound,
		Title:   TitleNotFound,
		Details: details,
	}
}

func NewBadRequestError(details string) *ApiError {
	return &ApiError{
		Code:    http.StatusBadRequest,
		Title:   TitleBadRequest,
		Details: details,
	}
}

func NewUnauthorizedError(details string) *ApiError {
	return &ApiError{
		Code:    http.StatusUnauthorized,
		Title:   TitleUnauthorized,
		Details: details,
	}
}
