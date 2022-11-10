package postgres

import (
	"context"
	"world_generator_processor_service/config"

	"github.com/olehmushka/golang-toolkit/wrapped_error"
	"github.com/olehmushka/postgres-golang"
	"go.uber.org/fx"
)

func New(cfg *config.Config, lc fx.Lifecycle) (postgres.Client, error) {
	client, err := postgres.New(&postgres.Config{
		Username:          cfg.Postgres.Client.Username,
		Password:          cfg.Postgres.Client.Password,
		DBName:            cfg.Postgres.Client.DBName,
		Host:              cfg.Postgres.Client.Host,
		Port:              cfg.Postgres.Client.Port,
		BatchItemsMaxSize: cfg.Postgres.BatchItemMaxSize,
	})
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not create postgres client")
	}

	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			client.ClosePool()
			return nil
		},
	})

	return client, nil
}

var Module = fx.Options(
	fx.Provide(New),
)
