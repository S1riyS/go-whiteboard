package whiteboardConverter

import (
	whiteboardv1 "github.com/S1riyS/go-whiteboard/shared/gen/go/whiteboard"
	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/entity"
)

func ToProtoWhiteboard(entity *entity.Whiteboard) *whiteboardv1.Whiteboard {
	return &whiteboardv1.Whiteboard{
		Id:          entity.ID,
		Title:       entity.Title,
		Description: entity.Description,
	}
}
