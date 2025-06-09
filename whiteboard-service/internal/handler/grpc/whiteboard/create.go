package whiteboard_grpc

import (
	"context"

	whiteboardv1 "github.com/S1riyS/go-whiteboard/shared/gen/go/whiteboard"
	whiteboardConverter "github.com/S1riyS/go-whiteboard/whiteboard-service/internal/converter/whiteboard"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateWhiteboard(ctx context.Context, req *whiteboardv1.CreateWhiteboardRequest) (*whiteboardv1.WhiteboardResponse, error) {
	entity := whiteboardConverter.FromProtoCreateRequest(req)
	id, err := s.svc.Create(ctx, entity)
	if err != nil {
		// TODO: more thorough error handling
		return nil, status.Error(codes.Internal, err.Error())
	}

	// TODO: rework gRPC contract to return only the ID
	// Even though under the hood repository sets the ID to the struct,
	// Here in the handler layer it is set explicitly for better readability.
	entity.ID = id

	return &whiteboardv1.WhiteboardResponse{
		Whiteboard: whiteboardConverter.ToProtoWhiteboard(entity),
	}, nil
}
