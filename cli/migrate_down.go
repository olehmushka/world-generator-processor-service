package cli

import (
	"context"
	"fmt"
	"os"
	"world_generator_processor_service/config"
	"world_generator_processor_service/core/zerodowntime"

	"github.com/golang-migrate/migrate/v4"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

const MigrateDownCommand = "migrate_down"

func runMigrateDownCommand() error {
	app := fx.New(
		fx.StartTimeout(defaultAppStartTimeout),
		fx.StopTimeout(defaultAppStopTimeout),
		fx.Options(
			config.Module,
			fx.Options(
				fx.Invoke(func(lc fx.Lifecycle, cfg *config.Config) {
					log := logrus.New().WithFields(logrus.Fields{
						"command": MigrateDownCommand,
					})
					lc.Append(fx.Hook{
						OnStart: func(ctx context.Context) error {
							log.Info(ctx, "command start")

							dbName := cfg.Postgres.Client.DBName
							dbUser := cfg.Postgres.Client.Username
							dbPass := cfg.Postgres.Client.Password
							dbHost := cfg.Postgres.Client.Host
							dbPort := cfg.Postgres.Client.Port

							dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

							m, err := migrate.New("file://migrations", dbURL)
							if err != nil {
								return err
							}
							if err = m.Down(); err != nil {
								return err
							}
							log.Info(ctx, "command finish")

							os.Exit(0)
							return nil
						},
					})

				}),
			),
		),
	)
	return zerodowntime.HandleApp(app)
}
