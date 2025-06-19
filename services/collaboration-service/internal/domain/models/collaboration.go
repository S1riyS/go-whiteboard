package models

import (
	"encoding/json"
	"time"
)

const (
	ElementTypeLine   = "line"
	ElementTypeCircle = "circle"
	ElementTypeStroke = "stroke"
)

type Drawing struct {
	WhiteboardID int             `json:"whiteboard_id"` // TODO: decide if this should be int, string or something else
	Timestamp    time.Time       `json:"timestamp"`
	ClientID     int             `json:"client_id"` // TODO: decide if this should be int, string or something else
	ElementType  string          `json:"element_type"`
	Element      json.RawMessage `json:"element"` //ElementStroke | ElementLine | ElementCircle
}
