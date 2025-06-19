package collaborationconverter

import (
	"fmt"

	"github.com/S1riyS/go-whiteboard/collaboration-service/internal/domain/models"
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

func FromProtoPoint(point *collaborationv1.Point) models.Point {
	return models.Point{
		X: point.GetX(),
		Y: point.GetY(),
	}
}

func FromProtoDrawRequest(req *collaborationv1.DrawRequest) (*models.Operation, error) {
	// Set Metadata
	drawing := &models.Operation{
		Type:         models.OperationCreate,
		DrawingID:    "0",
		WhiteboardID: req.GetWhiteboardId(),
		Timestamp:    req.GetMetadata().GetTimestamp().AsTime(),
	}

	// Set Type and Payload
	switch elem := req.GetElementType().(type) {
	// Line
	case *collaborationv1.DrawRequest_Line:
		drawing.ElementType = models.ElementTypeLine
		drawing.Payload = models.ElementLine{
			Start: FromProtoPoint(elem.Line.GetStart()),
			End:   FromProtoPoint(elem.Line.GetEnd()),
			Width: elem.Line.GetWidth(),
			Color: models.Color{Value: elem.Line.GetColor().GetValue()},
		}

	// Circle
	case *collaborationv1.DrawRequest_Circle:
		drawing.ElementType = models.ElementTypeCircle
		drawing.Payload = models.ElementCircle{
			Center: FromProtoPoint(elem.Circle.GetCenter()),
			Radius: elem.Circle.GetRadius(),
			Width:  elem.Circle.GetWidth(),
			Color:  models.Color{Value: elem.Circle.GetColor().GetValue()},
		}

	// Stroke
	case *collaborationv1.DrawRequest_Stroke:
		drawing.ElementType = models.ElementTypeStroke
		drawing.Payload = models.ElementStroke{
			Points: Map(elem.Stroke.GetPoints(), FromProtoPoint),
			Width:  elem.Stroke.GetWidth(),
			Color:  models.Color{Value: elem.Stroke.GetColor().GetValue()},
		}

	default:
		return nil, fmt.Errorf("unknown element type")
	}

	return drawing, nil
}

func FromProtoDeleteRequest(req *collaborationv1.DeleteRequest) (*models.Operation, error) {
	return &models.Operation{
		Type:         models.OperationDelete,
		DrawingID:    req.GetElementId(),
		WhiteboardID: req.GetWhiteboardId(),
		Timestamp:    req.GetMetadata().GetTimestamp().AsTime(),
	}, nil
}
