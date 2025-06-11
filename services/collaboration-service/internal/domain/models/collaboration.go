package models

import (
	"encoding/json"
	"time"
)

type Drawing struct {
	WhiteboardID int             `json:"whiteboard_id"`
	Timestamp    time.Time       `json:"timestamp"`
	ClientID     int             `json:"client_id"`
	ElementType  string          `json:"element_type"`
	Element      json.RawMessage `json:"element"`
}
