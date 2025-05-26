package app

import (
	"flag"
	"fmt"
	"log"

	v1 "github.com/S1riyS/go-whiteboard/api-gateway/internal/api/controller/v1"
	"github.com/S1riyS/go-whiteboard/api-gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

const (
	ENV_PATH = ".env"
)

// logLevel is a command-line flag for specifying the log level.
var logLevel = flag.String("l", "info", "log level")

type App struct {
	httpServer *gin.Engine
}

func New() *App {
	const mark = "app.New"

	app := &App{}
	err := app.runInitSteps()
	if err != nil {
		logger.Fatal("failed to init deps", mark)
	}
	return app
}

func (a *App) Run() error {
	const mark = "app.Run"

	// TODO: replace with real port from config
	rawPort := 8080
	// rawPort := a.provider.Config().App.Port
	processedPort := fmt.Sprintf(":%d", rawPort)

	err := a.httpServer.Run(processedPort)
	if err != nil {
		logger.Fatal("failed to start http server", mark, zap.Int("port", rawPort))
	}

	return nil
}

func (a *App) runInitSteps() error {
	const mark = "app.runInitSteps"

	initSteps := []func() error{
		a.initEnvironment,
		a.initLogger,
		a.initHttpServer,
		a.initServiceProvider,
		a.initValidator,
		a.initControllers,
	}

	for _, step := range initSteps {
		if err := step(); err != nil {
			logger.Fatal("failed to init deps", mark)
		}
	}

	return nil
}

func (a *App) initEnvironment() error {
	const mark = "app.initEnvironment"

	err := godotenv.Load(ENV_PATH)
	if err != nil {
		log.Fatal("error loading .env file", mark, zap.String("path", ENV_PATH))
		return fmt.Errorf("error loading %v file: %v", ENV_PATH, err)
	}
	return nil
}

func (a *App) initLogger() error {
	const mark = "app.initLogger"

	flag.Parse()                                                 // Parse command-line flags
	logger.Init(logger.GetCore(logger.GetAtomicLevel(logLevel))) // Initialize logger with the specified log level

	logger.Info("Logger initialized", mark)
	return nil
}

func (a *App) initHttpServer() error {
	const mark = "app.initHttpServer"

	a.httpServer = gin.Default()
	logger.Info("HTTP server initialized", mark)

	// // CORS policy
	// a.httpServer.Use(cors.New(cors.Config{
	// 	AllowOrigins:     "*",
	// 	AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
	// 	AllowHeaders:     "Accept,Authorization,Content-Type",
	// 	AllowCredentials: false, // credentials require explicit origins
	// 	MaxAge:           300,
	// }))

	// // Apply middlewares
	// a.httpServer.Use(
	// 	middlewares.Logger,       // Logger
	// 	middlewares.ErrorHandler, // Error handler
	// 	recover.New(),            // Recover
	// )

	return nil
}

func (a *App) initServiceProvider() error {
	// a.provider = newServiceProvider()
	return nil
}

func (a *App) initValidator() error {
	// validation.InitValidator()
	return nil
}

func (a *App) initControllers() error {
	api := a.httpServer.Group("/api")
	v1.SetupControllers(api)
	return nil
}
