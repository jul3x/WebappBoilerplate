package config

type ServerConfig struct {
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type Config struct {
	Server ServerConfig `mapstructure:"server"`
}
