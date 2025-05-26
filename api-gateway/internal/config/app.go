package config

type AppConfig struct {
	Env  string `env:"ENV" env-default:"development"`
	Port int    `env:"PORT" env-default:"8080"`
}
