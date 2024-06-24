package configs

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	AppPort          string `env:"APP_PORT" envDefault:"9000"`
	PostgresHost     string `env:"POSTGRES_HOST" envDefault:"psql"`
	PostgresPort     string `env:"POSTGRES_PORT" envDefault:"5432"`
	PostgresUser     string `env:"POSTGRES_USER" envDefault:"postgres"`
	PostgresPassword string `env:"POSTGRES_PASSWORD" envDefault:"postgres"`
	PostgresDBName   string `env:"POSTGRES_DB_NAME" envDefault:"postgres"`
	PostgresURL      string
}

func NewConfig() (*Config, error) {
	cfg := new(Config)
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	cfg.PostgresURL = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDBName)
	return cfg, nil
}
