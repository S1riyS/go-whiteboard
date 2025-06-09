package interceptors

import (
	"log/slog"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Recovery returns a gRPC recovery interceptor.
func Recovery(logger *slog.Logger) grpc.UnaryServerInterceptor {
	const mark = "interceptors.recovery"
	logger = logger.With(slog.String("mark", mark))

	// Options
	var recoveryOpts = []recovery.Option{
		recovery.WithRecoveryHandler(func(p any) (err error) {
			logger.Error("Recovered from panic", slog.Any("panic", p))
			return status.Errorf(codes.Internal, "internal error")
		}),
	}

	return recovery.UnaryServerInterceptor(recoveryOpts...)
}
