package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/app"
	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/config"
	"github.com/S1riyS/go-whiteboard/whiteboard-service/pkg/logger/slogpretty"
)

func main() {
	// Init config
	cfg := config.MustLoad()

	// Init logger
	logger := setupLogger(cfg.Env)

	// Print config in debug mode
	if cfg.Env == config.EnvLocal {
		logger.Info("Config loaded", slog.Any("config", cfg))
	}

	// Init and run app
	ctx := context.Background()
	application := app.New(ctx, logger, cfg)
	application.MustRun()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.Stop()
	logger.Info("Gracefully stopped")
}

func setupLogger(env config.EnvType) *slog.Logger {
	var log *slog.Logger

	switch env {
	case config.EnvLocal:
		log = setupPrettySlog()
	case config.EnvProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
