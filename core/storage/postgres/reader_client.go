package postgres

import (
	"context"
	"world_generator_processor_service/config"

	"github.com/olehmushka/golang-toolkit/wrapped_error"
	"github.com/olehmushka/postgres-golang"
	"go.uber.org/fx"
)

func NewReader(cfg *config.Config, lc fx.Lifecycle) (postgres.ClientReader, error) {
	client, err := postgres.NewReader(&postgres.Config{
		Username:          cfg.Postgres.ReaderClient.Username,
		Password:          cfg.Postgres.ReaderClient.Password,
		DBName:            cfg.Postgres.ReaderClient.DBName,
		Host:              cfg.Postgres.ReaderClient.Host,
		Port:              cfg.Postgres.ReaderClient.Port,
		BatchItemsMaxSize: cfg.Postgres.BatchItemMaxSize,
	})
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not create postgres reader client")
	}

	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			client.ClosePool()
			return nil
		},
	})

	return client, nil
}

var ReaderModule = fx.Options(
	fx.Provide(NewReader),
)
