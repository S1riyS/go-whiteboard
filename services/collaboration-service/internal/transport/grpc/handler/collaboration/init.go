package collaborationgrpc

import (
	"log/slog"

	collaborationservice "github.com/S1riyS/go-whiteboard/collaboration-service/internal/service/collaboration"
	collaborationv1 "github.com/S1riyS/go-whiteboard/shared/gen/go/collaboration"
	"google.golang.org/grpc"
)

type Server struct {
	collaborationv1.UnimplementedCollaborationV1Server
	logger *slog.Logger
	svc    *collaborationservice.Service
}

func Register(logger *slog.Logger, gRPCServer *grpc.Server, svc *collaborationservice.Service) {
	collaborationv1.RegisterCollaborationV1Server(gRPCServer, &Server{logger: logger, svc: svc})
}
