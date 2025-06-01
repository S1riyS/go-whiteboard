package whiteboard_grpc

import (
	"context"

	whiteboardv1 "github.com/S1riyS/go-whiteboard/shared/gen/go/whiteboard"
)

func (s *Server) CreateWhiteboard(ctx context.Context, req *whiteboardv1.CreateWhiteboardRequest) (*whiteboardv1.WhiteboardResponse, error) {
	return &whiteboardv1.WhiteboardResponse{
		Whiteboard: &whiteboardv1.Whiteboard{
			Id:          1,
			Title:       "Hello world!",
			Description: "Description of my very first gRPC call! (Create)",
		},
	}, nil
}
