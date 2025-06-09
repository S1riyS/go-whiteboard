package interceptors

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// result is a wrapper for the response and error
type result struct {
	resp any
	err  error
}

// Timeout returns a new unary server interceptor that adds timeout
func Timeout(logger *slog.Logger, timeout time.Duration) grpc.UnaryServerInterceptor {
	const mark = "interceptors.timeout"

	logger = logger.With(slog.String("mark", mark), slog.Duration("timeout", timeout))

	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		// Create a new context with timeout
		ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		// Call the handler with the new context
		resultChan := make(chan result, 1)
		go func() {
			resp, err := handler(ctxWithTimeout, req)
			resultChan <- result{resp: resp, err: err}
		}()

		select {
		case <-ctxWithTimeout.Done():
			// If timeout occurs, log and return DeadlineExceeded error
			logger.Warn("Request timed out", slog.String("method", info.FullMethod))
			return nil, status.Error(codes.DeadlineExceeded, "request timed out")
		case result := <-resultChan:
			// Return result from handler
			return result.resp, result.err
		}
	}
}
