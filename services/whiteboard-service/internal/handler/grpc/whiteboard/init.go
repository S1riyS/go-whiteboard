package whiteboard_grpc

import (
	"log/slog"

	whiteboardv1 "github.com/S1riyS/go-whiteboard/shared/gen/go/whiteboard"
	whiteboardService "github.com/S1riyS/go-whiteboard/whiteboard-service/internal/service/whiteboard"
	"google.golang.org/grpc"
)

type Server struct {
	whiteboardv1.UnimplementedWhiteboardV1Server
	logger *slog.Logger
	svc    *whiteboardService.Service
}

func Register(logger *slog.Logger, gRPCServer *grpc.Server, svc *whiteboardService.Service) {
	whiteboardv1.RegisterWhiteboardV1Server(gRPCServer, &Server{logger: logger, svc: svc})
}
