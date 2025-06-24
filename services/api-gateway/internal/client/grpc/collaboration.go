package clientgrpc

import (
	"context"
	"log/slog"

	"github.com/S1riyS/go-whiteboard/api-gateway/internal/config"
	convertercollab "github.com/S1riyS/go-whiteboard/api-gateway/internal/converter/collaboration"
	"github.com/S1riyS/go-whiteboard/api-gateway/internal/dto/request"
	"github.com/S1riyS/go-whiteboard/api-gateway/pkg/logger/slogext"
	collaborationv1 "github.com/S1riyS/go-whiteboard/shared/gen/go/collaboration"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// WhiteboardClient is a wrapper for whiteboard gRPC client.
//
// Provedes methods to interact with gRPC client without exposing underlying gRPC methods, variables, etc.
type CollaborationClient struct {
	logger     *slog.Logger
	grpcClient collaborationv1.CollaborationV1Client
}

func MustNewCollaborationClient(logger *slog.Logger, cfg config.CollaborationClientConfig) *CollaborationClient {
	conn, err := grpc.NewClient(
		cfg.Address(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic("failed to create new whiteboard client: " + err.Error())
	}

	return &CollaborationClient{
		logger:     logger,
		grpcClient: collaborationv1.NewCollaborationV1Client(conn),
	}
}

func (cc *CollaborationClient) Draw(ctx context.Context, whiteboardID string, req *request.CollaborationDrawPayload) (string, error) {
	const mark = "CollaborationClient.Draw"
	logger := cc.logger.With(slog.String("mark", mark))

	drawRequest, err := convertercollab.ToProtoDrawRequest(whiteboardID, req)
	if err != nil {
		logger.Warn("Failed to convert draw request", slog.Any("request", req), slogext.Err(err))
		return "", err
	}

	result, err := cc.grpcClient.Draw(ctx, drawRequest)
	if err != nil {
		logger.Warn("Failed to call draw", slogext.Err(err))
		return "", err
	}

	return result.GetId(), nil

}
