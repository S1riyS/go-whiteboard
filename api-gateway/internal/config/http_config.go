package config

type HttpConfig struct {
	Port         int      `yaml:"port" env-default:"8080"`
	AllowOrigins []string `yaml:"allowed_origins" env-required:"true"`
}
