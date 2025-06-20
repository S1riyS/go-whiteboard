package converterwhiteboard

import (
	"fmt"

	"github.com/S1riyS/go-whiteboard/api-gateway/internal/dto/response"
	whiteboardv1 "github.com/S1riyS/go-whiteboard/shared/gen/go/whiteboard"
	"github.com/google/uuid"
)

func FromProtoWhiteboardResponse(proto *whiteboardv1.Whiteboard) (*response.WhiteboardResponse, error) {
	uuid, err := uuid.Parse(proto.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to parse uuid: %w", err)
	}

	return &response.WhiteboardResponse{
		ID:          uuid,
		Title:       proto.Title,
		Description: proto.Description,
	}, nil
}
