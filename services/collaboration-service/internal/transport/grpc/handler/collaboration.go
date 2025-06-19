package handler

import (
	"context"
	"log/slog"

	"github.com/S1riyS/go-whiteboard/collaboration-service/internal/service"
	collaborationconverter "github.com/S1riyS/go-whiteboard/collaboration-service/internal/transport/grpc/converter/collaboration"
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
	// Convert draw request to drawing model
	drawing, err := collaborationconverter.FromProtoDrawRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid draw request")
	}

	s.logger.Debug("Received draw request", slog.Any("drawing", drawing))

	id, err := s.svc.Draw(ctx, drawing)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to draw")
	}

	return &collaborationv1.DrawResponse{Id: int64(id)}, nil
}

func (s *CollaborationHandler) Delete(ctx context.Context, req *collaborationv1.DeleteRequest) (*emptypb.Empty, error) {
	// TODO: implement
	return nil, nil
}
