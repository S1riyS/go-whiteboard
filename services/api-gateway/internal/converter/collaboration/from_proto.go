package convertercollab

import (
	"github.com/S1riyS/go-whiteboard/api-gateway/internal/dto/model"
	collaborationv1 "github.com/S1riyS/go-whiteboard/shared/gen/go/collaboration"
)

func fromProtoColor(color *collaborationv1.Color) model.Color {
	return model.Color{
		Value: color.Value,
	}
}

func fromProtoPoint(point *collaborationv1.Point) model.Point {
	return model.Point{
		X: point.X,
		Y: point.Y,
	}
}
