package request

type JoinWhiteboardRequest struct {
	WhiteboardID string `json:"whiteboard_id" validate:"required,uuid"`
}
