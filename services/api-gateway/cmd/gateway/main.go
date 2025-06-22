package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/S1riyS/go-whiteboard/api-gateway/internal/app"
	"github.com/S1riyS/go-whiteboard/api-gateway/internal/config"
	"github.com/S1riyS/go-whiteboard/api-gateway/pkg/logger/slogpretty"
)

func main() {
	const mark = "main"

	// Read path to config
	var configPath string
	flag.StringVar(&configPath, "config-path", "", "path to config file")
	flag.Parse()

	// Init config
	cfg := config.MustLoad(configPath)

	// Init logger
	logger := setupLogger(cfg.Env)

	// Print config in debug mode
	if cfg.Env == config.EnvLocal {
		logger.With(slog.String("mark", mark)).Debug("Config loaded", slog.Any("config", cfg))
	}

	// Init and run app
	ctx := context.Background()
	application := app.New(ctx, logger, *cfg)
	go application.MustRun() // Run app in separate goroutine

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.Stop()
	logger.With(slog.String("mark", mark)).Info("Gracefully stopped")
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
