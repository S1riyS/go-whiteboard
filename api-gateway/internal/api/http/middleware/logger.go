package middlewares

// TODO: rewrite to Gin

// import (
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"go.uber.org/zap"
// )

// func Logger(ctx *fiber.Ctx) error {
// 	const mark = "middleware.Logger"

// 	// Process request and calculate duration
// 	start := time.Now()
// 	err := ctx.Next()
// 	duration := time.Since(start)

// 	// Common fields
// 	details := []zap.Field{
// 		zap.String("method", ctx.Method()),
// 		zap.String("path", ctx.Path()),
// 		zap.Duration("duration", duration),
// 		zap.Int("status", ctx.Response().StatusCode()),
// 	}

// 	if err != nil {
// 		// If error, add it to details
// 		details = append(details, zap.Error(err))
// 		logger.Error("Failed to process request", mark, details...)
// 	} else {
// 		logger.Info("Request processed successfully", mark, details...)
// 	}
// 	return err
// }
