package whiteboard_grpc

import (
	"context"
	"errors"
	"log/slog"

	whiteboardv1 "github.com/S1riyS/go-whiteboard/shared/gen/go/whiteboard"
	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/service"
	"github.com/S1riyS/go-whiteboard/whiteboard-service/pkg/logger/slogext"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteWhiteboard(ctx context.Context, req *whiteboardv1.DeleteWhiteboardRequest) (*emptypb.Empty, error) {
	err := s.svc.Delete(ctx, int(req.GetId()))
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			s.logger.Debug("Whiteboard not found", slogext.Err(err), slog.Int64("id", req.GetId()))
			return nil, status.Error(codes.NotFound, "Whiteboard not found")
		}

		s.logger.Error("Internal error", slogext.Err(err), slog.Any("request", req))
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return nil, nil
}
