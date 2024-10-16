package config

import (
    "github.com/spf13/viper"
)

func LoadConfig(configPath string) (*Config, error) {
    var cfg Config

    viper.SetConfigName("config") // name of the config file without the extension
    viper.SetConfigType("yaml")
    viper.AddConfigPath(configPath)

    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }

    if err := viper.Unmarshal(&cfg); err != nil {
        return nil, err
    }

    return &cfg, nil
}
