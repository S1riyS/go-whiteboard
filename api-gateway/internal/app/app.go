package app

import (
	"flag"
	"fmt"
	"log"

	v1 "github.com/S1riyS/go-whiteboard/api-gateway/internal/api/controller/v1"
	"github.com/S1riyS/go-whiteboard/api-gateway/internal/config"
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

type app struct {
	config     *config.Config
	httpServer *gin.Engine
}

func New() *app {
	const mark = "app.New"

	app := &app{}
	err := app.runInitSteps()
	if err != nil {
		logger.Fatal("failed to init deps", mark)
	}
	return app
}

func (a *app) Run() error {
	const mark = "app.Run"

	rawPort := a.config.App.Port
	processedPort := fmt.Sprintf(":%d", rawPort)

	err := a.httpServer.Run(processedPort)
	if err != nil {
		logger.Fatal("failed to start http server", mark, zap.Int("port", rawPort))
	}

	return nil
}

func (a *app) runInitSteps() error {
	const mark = "app.runInitSteps"

	initSteps := []func() error{
		a.initEnvironment,
		a.initLogger,
		a.initConfig,
		a.initHttpServer,
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

func (a *app) initEnvironment() error {
	const mark = "app.initEnvironment"

	err := godotenv.Load(ENV_PATH)
	if err != nil {
		log.Fatal("error loading .env file", mark, zap.String("path", ENV_PATH))
		return fmt.Errorf("error loading %v file: %v", ENV_PATH, err)
	}

	return nil
}

func (a *app) initLogger() error {
	const mark = "app.initLogger"

	flag.Parse()                                                 // Parse command-line flags
	logger.Init(logger.GetCore(logger.GetAtomicLevel(logLevel))) // Initialize logger with the specified log level

	logger.Info("Logger initialized", mark)
	return nil
}

func (a *app) initConfig() error {
	const mark = "app.initConfig"

	a.config = config.GetConfig()

	logger.Info("Config initialized", mark)
	return nil
}

func (a *app) initHttpServer() error {
	const mark = "app.initHttpServer"

	a.httpServer = gin.Default()

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

	logger.Info("HTTP server initialized", mark)
	return nil
}

func (a *app) initValidator() error {
	const mark = "app.initValidator"

	// validation.InitValidator()

	logger.Warn("Validator is NOT initialized", mark)
	return nil
}

func (a *app) initControllers() error {
	const mark = "app.initControllers"

	// API
	apiGroup := a.httpServer.Group("/api")
	v1Group := apiGroup.Group("/v1")

	// Whiteboard
	whiteboardGroup := v1Group.Group("/whiteboard")
	whiteboardController := v1.NewWhiteboardController()
	{
		whiteboardGroup.POST("/", whiteboardController.Create)
		whiteboardGroup.GET("/", whiteboardController.GetOne)
		whiteboardGroup.PUT("/:id", whiteboardController.Update)
		whiteboardGroup.DELETE("/:id", whiteboardController.Delete)
	}

	logger.Info("Controllers initialized", mark)
	return nil
}
