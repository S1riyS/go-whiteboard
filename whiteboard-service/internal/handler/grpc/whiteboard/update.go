package whiteboard_grpc

import (
	"context"
	"errors"
	"log/slog"

	whiteboardv1 "github.com/S1riyS/go-whiteboard/shared/gen/go/whiteboard"
	whiteboardConverter "github.com/S1riyS/go-whiteboard/whiteboard-service/internal/converter/whiteboard"
	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/service"
	"github.com/S1riyS/go-whiteboard/whiteboard-service/pkg/logger/slogext"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateWhiteboard(ctx context.Context, req *whiteboardv1.UpdateWhiteboardRequest) (*whiteboardv1.WhiteboardResponse, error) {
	whiteboard := whiteboardConverter.FromProtoUpdateRequest(req)
	whiteboard, err := s.svc.Update(ctx, whiteboard)

	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			s.logger.Debug("Whiteboard not found", slogext.Err(err), slog.Int64("id", req.GetId()))
			return nil, status.Error(codes.NotFound, "Whiteboard not found")
		}

		s.logger.Error("Internal error", slogext.Err(err), slog.Any("request", req))
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return &whiteboardv1.WhiteboardResponse{
		Whiteboard: whiteboardConverter.ToProtoWhiteboard(whiteboard),
	}, nil
}
