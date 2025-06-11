package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env   EnvType     `yaml:"env" env-default:"local"`
	GRPC  GRPCConfig  `yaml:"grpc_server"`
	Redis RedisConfig `yaml:"redis"`
}

func MustLoad(configPath string) *Config {
	if configPath == "" {
		panic("config path is empty")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}
