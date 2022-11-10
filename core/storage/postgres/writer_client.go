package postgres

import (
	"context"
	"world_generator_processor_service/config"

	"github.com/olehmushka/golang-toolkit/wrapped_error"
	"github.com/olehmushka/postgres-golang"
	"go.uber.org/fx"
)

func NewWriter(cfg *config.Config, lc fx.Lifecycle) (postgres.ClientWriter, error) {
	client, err := postgres.NewWriter(&postgres.Config{
		Username:          cfg.Postgres.WriterCLient.Username,
		Password:          cfg.Postgres.WriterCLient.Password,
		DBName:            cfg.Postgres.WriterCLient.DBName,
		Host:              cfg.Postgres.WriterCLient.Host,
		Port:              cfg.Postgres.WriterCLient.Port,
		BatchItemsMaxSize: cfg.Postgres.BatchItemMaxSize,
	})
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not create postgres writer client")
	}

	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			client.ClosePool()
			return nil
		},
	})

	return client, nil
}

var WriterModule = fx.Options(
	fx.Provide(NewWriter),
)
