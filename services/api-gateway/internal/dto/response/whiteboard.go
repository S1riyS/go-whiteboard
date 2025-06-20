package response

import (
	"github.com/google/uuid"
)

type WhiteboardResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

type WhiteboardListResponse struct {
	Count       int                  `json:"count"`
	Whiteboards []WhiteboardResponse `json:"whiteboards"`
}
