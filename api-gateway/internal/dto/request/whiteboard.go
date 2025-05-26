package request

import "time"

type CreateWhiteboardRequest struct {
	Title       string  `json:"title" validate:"required"`
	Description *string `json:"description"`
}

type UpdateWhiteboardRequest struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
}

// Response DTOs
type WhiteboardResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	OwnerID   string    `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WhiteboardListResponse struct {
	Whiteboards []WhiteboardResponse `json:"whiteboards"`
	Count       int                  `json:"count"`
}
