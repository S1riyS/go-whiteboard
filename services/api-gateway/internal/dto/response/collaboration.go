package response

type CollaborationSessionResponse struct {
	SessionID string `json:"session_id"`
	WSURL     string `json:"ws_url"` // e.g., "ws://localhost/ws?token=..."
}
