package config

import (
	"errors"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Env         string         `env:"ENV" envDefault:"dev"`
	Port        string         `env:"PORT" envDefault:"8080"`
	Database    DatabaseConfig `envPrefix:"DATABASE_"`
	JWTSecretKey string `env:"JWT_SECRET_KEY" envDefault:"secret"`
}

type DatabaseConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     string `env:"PORT" envDefault:"3006"`
	User     string `env:"USER" envDefault:"user"`
	Password string `env:"PASSWORD" envDefault:""`
	Database string `env:"DATABASE" envDefault:"database"`
}

func NewConfig(envPath string) (*Config, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return nil, errors.New("failed to load .env file")
	}

	cfg := new(Config)

	err = env.Parse(cfg)
	if err != nil {
		return nil, errors.New("failed to parse config file")
	}

	return cfg, nil
}