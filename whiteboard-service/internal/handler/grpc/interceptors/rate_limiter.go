package interceptors

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TimeoutInterceptor returns a new unary client interceptor that adds timeout
func TimeoutInterceptor(logger *slog.Logger, timeout time.Duration) grpc.UnaryClientInterceptor {
	const mark = "interceptors.rate_limiter"

	logger = logger.With(slog.String("mark", mark))

	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		errChan := make(chan error, 1)
		go func() {
			errChan <- invoker(ctx, method, req, reply, cc, opts...)
		}()

		select {
		case <-ctx.Done():
			logger.Debug("Client request timed out", slog.String("method", method))
			return status.Error(codes.DeadlineExceeded, "client request timed out")
		case err := <-errChan:
			return err
		}
	}
}
