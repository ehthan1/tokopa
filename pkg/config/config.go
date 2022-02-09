package config

import "github.com/google/wire"

type AppConfig struct {
	Port    string `yaml:"port"`
	Version string `yaml: "verion"`
	Name    string `yaml: "name"`
}

type DatabaseConfig struct {
}
type RedisConfig struct {
}
type Config struct {
	App   AppConfig      `yaml: "app"`
	Db    DatabaseConfig `yaml:"database"`
	Redis RedisConfig    `yaml:"redis"`
}

func New(filepath string) (*Config, error) {
	return &Config{}, nil
}

var ProviderSet = wire.NewSet(New)
