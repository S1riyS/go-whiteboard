package config

type CollaborationClientConfig struct {
	Host string `yaml:"host" env-required:"true"`
	Port string `yaml:"port" env-required:"true"`
}

func (c *CollaborationClientConfig) Address() string {
	return c.Host + ":" + c.Port
}
