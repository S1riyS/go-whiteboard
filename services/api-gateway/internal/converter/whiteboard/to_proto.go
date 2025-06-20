package converterwhiteboard

import (
	"github.com/S1riyS/go-whiteboard/api-gateway/internal/dto/request"
	whiteboardv1 "github.com/S1riyS/go-whiteboard/shared/gen/go/whiteboard"
	"github.com/google/uuid"
)

func ToProtoCreateRequest(req *request.CreateWhiteboardRequest) *whiteboardv1.CreateWhiteboardRequest {
	return &whiteboardv1.CreateWhiteboardRequest{
		Name:        req.Title,
		Description: req.Description,
	}
}

func ToProtoUpdateRequest(id uuid.UUID, req *request.UpdateWhiteboardRequest) *whiteboardv1.UpdateWhiteboardRequest {
	return &whiteboardv1.UpdateWhiteboardRequest{
		Id:          id.String(),
		Name:        req.Title,
		Description: req.Description,
	}
}
