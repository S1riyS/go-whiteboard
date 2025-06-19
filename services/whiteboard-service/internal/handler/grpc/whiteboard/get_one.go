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

func (s *Server) GetWhiteboard(ctx context.Context, req *whiteboardv1.GetWhiteboardRequest) (*whiteboardv1.WhiteboardResponse, error) {
	const mark = "whiteboard_grpc.Server.GetWhiteboard"

	id := req.GetId()

	logger := s.logger.With(slog.String("mark", mark), slog.String("id", id))

	whiteboard, err := s.svc.GetOne(ctx, id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			logger.Debug("Whiteboard not found", slogext.Err(err))
			return nil, status.Error(codes.NotFound, "Whiteboard not found")
		}

		logger.Error("Internal error", slogext.Err(err))
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return &whiteboardv1.WhiteboardResponse{
		Whiteboard: whiteboardConverter.ToProtoWhiteboard(whiteboard),
	}, nil
}
