package config

import (
	"sync"

	"github.com/S1riyS/go-whiteboard/api-gateway/pkg/logger"
	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
)

type Config struct {
	App        AppConfig
	Auth       AuthClientConfig
	Whiteboard WhiteboardClientConfig
}

type AppConfig struct {
	Env  string `env:"ENV" env-default:"development"`
	Port int    `env:"API_GATEWAY_PORT" env-default:"8080"`
}

type AuthClientConfig struct {
	Host string `env:"GRPC_AUTH_HOST" env-required:"true"`
	Port string `env:"GRPC_AUTH_PORT" env-required:"true"`
}

type WhiteboardClientConfig struct {
	Host string `env:"GRPC_WHITEBOARD_HOST" env-required:"true"`
	Port string `env:"GRPC_WHITEBOARD_PORT" env-required:"true"`
}

var (
	instance *Config
	once     sync.Once
)

// GetConfig returns the application configuration.
// Note that config.Config is a singleton
func GetConfig() *Config {
	const mark = "config.GetConfig"

	once.Do(func() {
		logger.Info("Read application configuration", mark)
		instance = &Config{}
		if err := cleanenv.ReadEnv(instance); err != nil {
			logger.Fatal("Failed to read application configuration", mark, zap.Error(err))
		}
	})
	return instance
}
