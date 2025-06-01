package whiteboard_grpc

import (
	"context"

	whiteboardv1 "github.com/S1riyS/go-whiteboard/shared/gen/go/whiteboard"
)

func (s *Server) GetWhiteboard(ctx context.Context, req *whiteboardv1.GetWhiteboardRequest) (*whiteboardv1.WhiteboardResponse, error) {
	return &whiteboardv1.WhiteboardResponse{
		Whiteboard: &whiteboardv1.Whiteboard{
			Id:          1,
			Title:       "Hello world!",
			Description: "Description of my very first gRPC call! (Get one)",
		},
	}, nil
}
