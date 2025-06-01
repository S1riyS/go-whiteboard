package whiteboard_grpc

import (
	"context"

	whiteboardv1 "github.com/S1riyS/go-whiteboard/shared/gen/go/whiteboard"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteWhiteboard(ctx context.Context, req *whiteboardv1.DeleteWhiteboardRequest) (*emptypb.Empty, error) {
	return nil, nil
}
