package cli

import (
	"context"
	"os"
	"world_generator_processor_service/core/zerodowntime"
	gendersServices "world_generator_processor_service/internal/genders/services"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

const SyncDataCommand = "sync_data"

func runSyncDataCommand() error {
	app := fx.New(
		fx.StartTimeout(defaultAppStartTimeout),
		fx.StopTimeout(defaultAppStopTimeout),
		fx.Options(
			Modules,
			fx.Options(
				fx.Invoke(func(
					genderSyncSrv gendersServices.Sync,
				) {
					log := logrus.New().WithFields(logrus.Fields{
						"command": SyncDataCommand,
					})
					ctx := context.Background()
					log.Info(ctx, "command start")
					if err := genderSyncSrv.SyncGenders(ctx); err != nil {
						log.Errorln(err)
						log.Info(ctx, "command finish")
						os.Exit(0)
						return
					}
					if err := genderSyncSrv.SyncGenderAcceptances(ctx); err != nil {
						log.Errorln(err)
						log.Info(ctx, "command finish")
						os.Exit(0)
						return
					}

					log.Info(ctx, "command finish")
					os.Exit(0)
				}),
			),
		),
	)
	return zerodowntime.HandleApp(app)
}
