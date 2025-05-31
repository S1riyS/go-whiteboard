package app

import (
	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/app/grpc_server"
)

type app struct {
	grpcServer *grpc_server.GrpcServer
}

func New() *app {
	return &app{
		grpcServer: grpc_server.New(),
	}
}

func (a *app) MustRun() {
	// TODO: implement me
	panic("implement me")
}

func (a *app) Stop() {
	// TODO: implement me
	panic("implement me")
}
