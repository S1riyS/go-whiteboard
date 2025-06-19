package models

import "time"

type OprerationType string

const (
	OperationCreate OprerationType = "create"
	OperationUpdate OprerationType = "update"
	OperationDelete OprerationType = "delete"
)

type Operation struct {
	Type         OprerationType `json:"type"`
	DrawingID    string         `json:"drawingId"`
	WhiteboardID string         `json:"whiteboardId"`
	ElementType  ElementType    `json:"elementType"`
	Payload      any            `json:"payload,omitempty"`
	Timestamp    time.Time      `json:"timestamp"`
}
