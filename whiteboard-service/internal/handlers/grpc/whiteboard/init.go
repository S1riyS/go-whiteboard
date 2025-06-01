package whiteboard_grpc

import (
	whiteboardv1 "github.com/S1riyS/go-whiteboard/shared/gen/go/whiteboard"
	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/service/whiteboard"
	"google.golang.org/grpc"
)

type Server struct {
	whiteboardv1.UnimplementedWhiteboardServiceServer
	svc *whiteboard.Service
}

func Register(gRPCServer *grpc.Server, svc *whiteboard.Service) {
	whiteboardv1.RegisterWhiteboardServiceServer(gRPCServer, &Server{svc: svc})
}
