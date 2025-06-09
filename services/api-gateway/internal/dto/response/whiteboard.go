package response

import (
	"time"

	"github.com/google/uuid"
)

type WhiteboardResponse struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	OwnerID   string    `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WhiteboardListResponse struct {
	Count       int                  `json:"count"`
	Whiteboards []WhiteboardResponse `json:"whiteboards"`
}
