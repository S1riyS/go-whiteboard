package config

type WhiteboardClientConfig struct {
	Host string `yaml:"host" env-required:"true"`
	Port string `yaml:"port" env-required:"true"`
}

func (c *WhiteboardClientConfig) Address() string {
	return c.Host + ":" + c.Port
}
