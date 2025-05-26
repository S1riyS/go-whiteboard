package response

import "time"

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
