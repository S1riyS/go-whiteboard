package config

import "fmt"

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}

func (c *RedisConfig) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
