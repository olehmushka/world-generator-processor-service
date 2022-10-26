package config

import (
	"github.com/jinzhu/configor"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

type Config struct {
	HTTPServer HTTPServer
	MongoDB    MongoDB
	Redis      Redis
}

func New() (*Config, error) {
	var cfg Config

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	if err := configor.New(&configor.Config{ErrorOnUnmatchedKeys: true}).Load(&cfg, "config/default.json"); err != nil {
		return nil, err
	}

	return &cfg, nil
}

var Module = fx.Options(
	fx.Provide(New),
)
