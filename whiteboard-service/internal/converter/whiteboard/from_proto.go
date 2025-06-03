package whiteboardConverter

import (
	whiteboardv1 "github.com/S1riyS/go-whiteboard/shared/gen/go/whiteboard"
	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/entity"
)

func FromProtoCreateRequest(proto *whiteboardv1.CreateWhiteboardRequest) *entity.Whiteboard {
	return &entity.Whiteboard{
		Title:       proto.Name,
		Description: proto.Description,
	}
}

func FromProtoUpdateRequest(proto *whiteboardv1.UpdateWhiteboardRequest) *entity.Whiteboard {
	return &entity.Whiteboard{
		ID:          int(proto.Id),
		Title:       proto.Name,
		Description: proto.Description,
	}
}

func FromProtoWhiteboard(proto *whiteboardv1.Whiteboard) *entity.Whiteboard {
	return &entity.Whiteboard{
		ID:          int(proto.Id),
		Title:       proto.Title,
		Description: proto.Description,
	}
}
