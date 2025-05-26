package app

import "github.com/S1riyS/go-whiteboard/api-gateway/internal/config"

type IContainer interface {
	GetClosers() []func() error
}

// TODO: add remaining configs
type CoreContainer struct {
	AppConfig *config.AppConfig
}

// TODO: GRPC clients and controllers

// type AuthContainer struct {
// 	Client     IAuthClient
// 	Controller *AuthController
// }

// type WhiteboardContainer struct {
// 	Client     IWhiteboardClient
// 	Controller *WhiteboardController
// }

// type CollaborationContainer struct {
// 	Client     ICollaborationClient
// 	Controller *CollaborationController
// }
