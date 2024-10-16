package config

type ServerConfig struct {
	Name         string `mapstructure:"name"`
	Port         int    `mapstructure:"port"`
	Host         string `mapstructure:"host"`
	FrontendPort int    `mapstructure:"frontend_port"`
}

type Config struct {
	Server ServerConfig `mapstructure:"server"`
}
