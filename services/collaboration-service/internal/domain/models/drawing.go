package models

import (
	"time"
)

type ElementType string

const (
	ElementTypeStroke ElementType = "stroke"
	ElementTypeLine   ElementType = "line"
	ElementTypeCircle ElementType = "circle"
)

type DrawingMetadata struct {
	ID           string    `json:"id"`
	WhiteboardID string    `json:"whiteboardId"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Version      int       `json:"version"`
}

type Drawing struct {
	DrawingMetadata
	Type    ElementType `json:"type"`
	Payload any         `json:"payload"` // ElementStroke | ElementLine | ElementCircle
}
