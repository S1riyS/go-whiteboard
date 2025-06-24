package request

import (
	"encoding/json"

	"github.com/S1riyS/go-whiteboard/api-gateway/internal/dto/model"
)

type CollaborationType string

const (
	CollaborationTypeJoin   CollaborationType = "JOIN"
	CollaborationTypeDraw   CollaborationType = "DRAW"
	CollaborationTypeDelete CollaborationType = "DELETE"
)

type CollaborationMessage struct {
	Type    CollaborationType `json:"type"`
	Payload json.RawMessage   `json:"payload"` // CollaborationJoinPayload | CollaborationDrawPayload | CollaborationDeletePayload
}

type CollaborationJoinPayload struct {
	UserID string `json:"user_id"`
}

type CollaborationDrawPayload struct {
	ElementType model.ElementType `json:"element_type"`
	Data        json.RawMessage   `json:"data"` // ElementStroke | ElementLine | ElementCircle
}

type CollaborationDeletePayload struct {
	ElementID string `json:"element_id"`
}
