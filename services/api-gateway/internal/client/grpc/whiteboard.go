package clientgrpc

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/S1riyS/go-whiteboard/api-gateway/internal/api"
	"github.com/S1riyS/go-whiteboard/api-gateway/internal/config"
	converterwhiteboard "github.com/S1riyS/go-whiteboard/api-gateway/internal/converter/whiteboard"
	"github.com/S1riyS/go-whiteboard/api-gateway/internal/dto/request"
	"github.com/S1riyS/go-whiteboard/api-gateway/internal/dto/response"
	"github.com/S1riyS/go-whiteboard/api-gateway/pkg/logger/slogext"
	whiteboardv1 "github.com/S1riyS/go-whiteboard/shared/gen/go/whiteboard"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

// WhiteboardClient is a wrapper for whiteboard gRPC client.
//
// Provedes methods to interact with gRPC client without exposing underlying gRPC methods, variables, etc.
type WhiteboardClient struct {
	logger     *slog.Logger
	grpcClient whiteboardv1.WhiteboardV1Client
}

func MustNewWhiteboardClient(logger *slog.Logger, cfg config.WhiteboardClientConfig) *WhiteboardClient {
	conn, err := grpc.NewClient(
		cfg.Address(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic("failed to create new whiteboard client: " + err.Error())
	}

	return &WhiteboardClient{
		logger:     logger,
		grpcClient: whiteboardv1.NewWhiteboardV1Client(conn),
	}
}

func (wc *WhiteboardClient) CreateWhiteboard(ctx context.Context, req *request.CreateWhiteboardRequest) (*response.WhiteboardResponse, error) {
	const mark = "WhiteboardClient.CreateWhiteboard"
	logger := wc.logger.With(slog.String("mark", mark))

	result, err := wc.grpcClient.CreateWhiteboard(ctx, converterwhiteboard.ToProtoCreateRequest(req))

	// Error handling
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			//nolint:exhaustive // Later list of cases might be extended
			switch st.Code() {
			case codes.Internal:
				return nil, api.NewInternalError()
			}
		}
		logger.Error("Failed to create whiteboard", slogext.Err(err))
		return nil, fmt.Errorf("failed to create whiteboard: %w", err)
	}

	return converterwhiteboard.FromProtoWhiteboardResponse(result.GetWhiteboard())
}

func (wc *WhiteboardClient) GetWhiteboard(ctx context.Context, id uuid.UUID) (*response.WhiteboardResponse, error) {
	const mark = "WhiteboardClient.CreateWhiteboard"
	logger := wc.logger.With(slog.String("mark", mark))

	result, err := wc.grpcClient.GetWhiteboard(ctx, &whiteboardv1.GetWhiteboardRequest{Id: id.String()})

	// Error handling
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			//nolint:exhaustive // Later list of cases might be extended
			switch st.Code() {
			case codes.NotFound:
				return nil, api.NewNotFoundError("Whiteboard not found")
			}
		}
		logger.Error("Failed to get whiteboard", slogext.Err(err))
		return nil, fmt.Errorf("failed to get whiteboard: %w", err)
	}

	return converterwhiteboard.FromProtoWhiteboardResponse(result.GetWhiteboard())
}

func (wc *WhiteboardClient) UpdateWhiteboard(ctx context.Context, id uuid.UUID, req *request.UpdateWhiteboardRequest) (*response.WhiteboardResponse, error) {
	const mark = "WhiteboardClient.UpdateWhiteboard"
	logger := wc.logger.With(slog.String("mark", mark))

	result, err := wc.grpcClient.UpdateWhiteboard(ctx, converterwhiteboard.ToProtoUpdateRequest(id, req))

	// Error handling
	if err != nil {
		st, ok := status.FromError(err)
		//nolint:exhaustive // Later list of cases might be extended
		if ok {
			switch st.Code() {
			case codes.NotFound:
				return nil, api.NewNotFoundError("Whiteboard not found")
			}
		}
		logger.Error("Failed to update whiteboard", slogext.Err(err))
		return nil, fmt.Errorf("failed to update whiteboard: %w", err)
	}

	return converterwhiteboard.FromProtoWhiteboardResponse(result.GetWhiteboard())
}

func (wc *WhiteboardClient) DeleteWhiteboard(ctx context.Context, id uuid.UUID) error {
	const mark = "WhiteboardClient.UpdateWhiteboard"
	logger := wc.logger.With(slog.String("mark", mark))

	_, err := wc.grpcClient.DeleteWhiteboard(ctx, &whiteboardv1.DeleteWhiteboardRequest{Id: id.String()})

	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			//nolint:exhaustive // Later list of cases might be extended
			switch st.Code() {
			case codes.NotFound:
				return api.NewNotFoundError("Whiteboard not found")
			}
		}
		logger.Error("Failed to delete whiteboard", slogext.Err(err))
		return fmt.Errorf("failed to delete whiteboard: %w", err)
	}

	return nil
}
