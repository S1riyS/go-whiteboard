package convertercollab

import (
	"encoding/json"
	"fmt"

	"github.com/S1riyS/go-whiteboard/api-gateway/internal/dto/model"
	"github.com/S1riyS/go-whiteboard/api-gateway/internal/dto/request"
	collaborationv1 "github.com/S1riyS/go-whiteboard/shared/gen/go/collaboration"
)

// Map applies a function to each element of a slice
func Map[T any, R any](slice []T, f func(T) R) []R { // TODO: Move to utils
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

func toProtoPoint(point model.Point) *collaborationv1.Point {
	return &collaborationv1.Point{
		X: point.X,
		Y: point.Y,
	}
}

func toProtoColor(color model.Color) *collaborationv1.Color {
	return &collaborationv1.Color{
		Value: color.Value,
	}
}

func ToProtoDrawRequest(whiteboardID string, req *request.CollaborationDrawPayload) (*collaborationv1.DrawRequest, error) {
	// FIXME: ElementType switch doesn't ensure proper Unmarshal. Should be fixed
	// TODO: Move Unmarshal somewhere else. For beter error handling

	result := &collaborationv1.DrawRequest{
		Metadata: &collaborationv1.RequestMetadata{ // TODO: create from request
			Timestamp: nil,
			ClientId:  "None",
		},
		WhiteboardId: whiteboardID,
	}

	switch req.ElementType {
	case model.ElementTypeLine:
		var line model.ElementLine
		err := json.Unmarshal(req.Data, &line)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal line: %w", err)
		}
		result.ElementType = &collaborationv1.DrawRequest_Line{
			Line: &collaborationv1.ElementLine{
				Start: toProtoPoint(line.Start),
				End:   toProtoPoint(line.End),
				Width: line.Width,
				Color: toProtoColor(line.Color),
			},
		}

	case model.ElementTypeCircle:
		var circle model.ElementCircle
		err := json.Unmarshal(req.Data, &circle)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal circle: %w", err)
		}
		result.ElementType = &collaborationv1.DrawRequest_Circle{
			Circle: &collaborationv1.ElementCircle{
				Center: toProtoPoint(circle.Center),
				Radius: circle.Radius,
				Width:  circle.Width,
				Color:  toProtoColor(circle.Color),
			},
		}

	case model.ElementTypeStroke:
		var stroke model.ElementStroke
		err := json.Unmarshal(req.Data, &stroke)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal stroke: %w", err)
		}
		result.ElementType = &collaborationv1.DrawRequest_Stroke{
			Stroke: &collaborationv1.ElementStroke{
				Points: Map(stroke.Points, toProtoPoint),
				Width:  stroke.Width,
				Color:  toProtoColor(stroke.Color),
			},
		}

	default:
		return nil, fmt.Errorf("unknown element type: %T", req.Data)
	}

	return result, nil
}
