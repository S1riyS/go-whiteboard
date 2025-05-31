package app

import (
	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/app/grpc_server"
)

type App struct {
	grpcServer *grpc_server.GrpcServer
}

func New() *App {
	return &App{
		grpcServer: grpc_server.New(),
	}
}

func (a *App) MustRun() {
	// TODO: implement me
	panic("implement me")
}
