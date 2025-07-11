package httpServer

import (
	"fmt"
	"log/slog"
	"time"

	v1 "github.com/S1riyS/go-whiteboard/api-gateway/internal/api/http/controller/v1"
	middlewares "github.com/S1riyS/go-whiteboard/api-gateway/internal/api/http/middleware"
	clientgrpc "github.com/S1riyS/go-whiteboard/api-gateway/internal/client/grpc"
	"github.com/S1riyS/go-whiteboard/api-gateway/internal/config"
	"github.com/S1riyS/go-whiteboard/api-gateway/pkg/logger/slogext"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	logger *slog.Logger
	config config.Config

	ginInstance *gin.Engine // Gin engine that runs on `httpSrv`
	// httpSrv     *http.Server // Underlying HTTP server
}

func New(logger *slog.Logger, config config.Config) *Server {
	// const mark = "httpServer.New"

	server := &Server{
		logger: logger,
		config: config,
	}

	server.initGin()
	server.initControllers()

	return server
}

// Run starts the http server.
//
// It runs the gin.Engine and returns an error if it can't start the server.
// The port number is taken from the configuration.
func (hs *Server) Run() error {
	const mark = "httpServer.Run"

	logger := hs.logger.With(slog.String("mark", mark))

	port := fmt.Sprintf(":%d", hs.config.HTTP.Port)
	if err := hs.ginInstance.Run(port); err != nil {
		logger.Error("failed to start http server", slog.Int("port", hs.config.HTTP.Port), slogext.Err(err))
	}

	return nil
}

func (hs *Server) Stop() {
	const mark = "httpServer.Stop"

	logger := hs.logger.With(slog.String("mark", mark))
	logger.Warn("httpServer.Stop is NOT implemented yet", slog.Int("port", hs.config.HTTP.Port))
}

func (hs *Server) initGin() {
	// const mark = "httpServer.setupGinEngine"

	hs.ginInstance = gin.New()

	// CORS configuration
	hs.ginInstance.Use(cors.New(cors.Config{
		AllowOrigins:     hs.config.HTTP.AllowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, //nolint:mnd // Later will be retrieved from config (probably)
	}))

	// Middlewares
	hs.ginInstance.Use(
		gin.Recovery(),
		middlewares.ErrorHandler(),
		gin.Logger(),
	)
}

func (hs *Server) initControllers() {
	// const mark = "httpServer.initControllers"

	// API
	apiGroup := hs.ginInstance.Group("/api")
	v1Group := apiGroup.Group("/v1")

	// Whiteboard
	whiteboardGroup := v1Group.Group("/whiteboards")
	whiteboardClient := clientgrpc.MustNewWhiteboardClient(hs.logger, hs.config.Whiteboard)
	whiteboardController := v1.NewWhiteboardController(hs.logger, whiteboardClient)
	{
		whiteboardGroup.POST("/", whiteboardController.Create)
		whiteboardGroup.GET("/:id", whiteboardController.GetOne)
		whiteboardGroup.PUT("/:id", whiteboardController.Update)
		whiteboardGroup.DELETE("/:id", whiteboardController.Delete)
	}
}
