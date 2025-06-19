package handler

import (
	"context"
	"errors"
	"log/slog"

	"github.com/S1riyS/go-whiteboard/collaboration-service/internal/service"
	collaborationconverter "github.com/S1riyS/go-whiteboard/collaboration-service/internal/transport/grpc/converter/collaboration"
	"github.com/S1riyS/go-whiteboard/collaboration-service/pkg/logger/slogext"
	collaborationv1 "github.com/S1riyS/go-whiteboard/shared/gen/go/collaboration"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CollaborationHandler struct {
	collaborationv1.UnimplementedCollaborationV1Server
	logger *slog.Logger
	svc    *service.CollaborationService
}

func RegisterCollaborationHandler(logger *slog.Logger, gRPCServer *grpc.Server, svc *service.CollaborationService) {
	collaborationv1.RegisterCollaborationV1Server(gRPCServer, &CollaborationHandler{logger: logger, svc: svc})
}

func (s *CollaborationHandler) Draw(ctx context.Context, req *collaborationv1.DrawRequest) (*collaborationv1.DrawResponse, error) {
	const mark = "service.collaboration.Draw"
	logger := s.logger.With(slog.String("mark", mark))

	// Convert draw request to operation
	operation, err := collaborationconverter.FromProtoDrawRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid draw request")
	}

	logger.Debug("Received draw request", slog.Any("operaion", operation))

	// Handle operation
	drawing, err := s.svc.HandleOperation(ctx, *operation)
	if err != nil {
		logger.Error("Failed to draw", slogext.Err(err))
		return nil, status.Error(codes.Internal, "Failed to draw")
	}

	return &collaborationv1.DrawResponse{Id: drawing.ID}, nil
}

func (s *CollaborationHandler) Delete(ctx context.Context, req *collaborationv1.DeleteRequest) (*emptypb.Empty, error) {
	const mark = "service.collaboration.Delete"
	logger := s.logger.With(slog.String("mark", mark))

	// Convert delete request to operation
	operation, err := collaborationconverter.FromProtoDeleteRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid delete request")
	}

	logger.Debug("Received delete request", slog.Any("operaion", operation))

	// Handle operation
	_, err = s.svc.HandleOperation(ctx, *operation)
	if errors.Is(err, service.ErrDrawingNotFound) {
		return nil, status.Error(codes.NotFound, "Drawing not found")
	}
	if err != nil {
		logger.Error("Failed to delete", slogext.Err(err))
		return nil, status.Error(codes.Internal, "Failed to delete")
	}

	return &emptypb.Empty{}, nil
}
