package collaborationconverter

import (
	"encoding/json"
	"fmt"

	"github.com/S1riyS/go-whiteboard/collaboration-service/internal/domain/models"
	collaborationv1 "github.com/S1riyS/go-whiteboard/shared/gen/go/collaboration"
)

func FromProtoDrawRequest(req *collaborationv1.DrawRequest) (*models.Drawing, error) {
	// drawing := &models.Drawing{
	// 	WhiteboardID: req.GetWhiteboardId(),
	// 	Timestamp:    req.GetMetadata().GetTimestamp().AsTime(),
	// 	ClientID:     req.GetMetadata().GetClientId(),
	// }
	// TODO: repalce const IDs with data from request
	drawing := &models.Drawing{
		WhiteboardID: 1,
		Timestamp:    req.GetMetadata().GetTimestamp().AsTime(),
		ClientID:     2,
	}

	// Parse element
	switch elem := req.GetElementType().(type) {
	case *collaborationv1.DrawRequest_Line:
		drawing.ElementType = models.ElementTypeLine
		elementData, err := json.Marshal(elem.Line)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal line element: %w", err)
		}
		drawing.Element = elementData

	case *collaborationv1.DrawRequest_Circle:
		drawing.ElementType = models.ElementTypeCircle
		elementData, err := json.Marshal(elem.Circle)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal circle element: %w", err)
		}
		drawing.Element = elementData

	case *collaborationv1.DrawRequest_Stroke:
		drawing.ElementType = models.ElementTypeStroke
		elementData, err := json.Marshal(elem.Stroke)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal stroke element: %w", err)
		}
		drawing.Element = elementData

	default:
		return nil, fmt.Errorf("unknown element type")
	}

	return drawing, nil
}
