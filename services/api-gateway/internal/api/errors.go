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

type Error struct {
	Code    int    `json:"code"`
	Title   string `json:"title"`
	Details string `json:"details"`
}

func NewError(code int, title string, details string) *Error {
	return &Error{
		Code:    code,
		Title:   title,
		Details: details,
	}
}

func NewInternalError() *Error {
	return &Error{
		Code:    http.StatusInternalServerError,
		Title:   TitleInternalError,
		Details: MessageInternalError,
	}
}

func NewUnprocessableEntityError() *Error {
	return &Error{
		Code:    http.StatusUnprocessableEntity,
		Title:   TitleUnprocessableEntity,
		Details: MessageUnprocessableEntity,
	}
}

func NewNotFoundError(details string) *Error {
	return &Error{
		Code:    http.StatusNotFound,
		Title:   TitleNotFound,
		Details: details,
	}
}

func NewBadRequestError(details string) *Error {
	return &Error{
		Code:    http.StatusBadRequest,
		Title:   TitleBadRequest,
		Details: details,
	}
}

func NewUnauthorizedError(details string) *Error {
	return &Error{
		Code:    http.StatusUnauthorized,
		Title:   TitleUnauthorized,
		Details: details,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s (code %d)", e.Title, e.Details, e.Code)
}

// WriteToContext writes the error to the gin context.
func (e *Error) WriteToContext(ctx *gin.Context) {
	WriteErrorToContext(ctx, e)
}

func WriteErrorToContext(ctx *gin.Context, err error) {
	ctx.Error(err) //nolint:errcheck // It's ok here. Returns passed error
}
