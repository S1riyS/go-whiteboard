package request

type CreateWhiteboardRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}

type UpdateWhiteboardRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
