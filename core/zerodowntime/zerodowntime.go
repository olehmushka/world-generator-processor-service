package zerodowntime

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func HandleApp(app *fx.App) error {
	startCtx, cancel := context.WithTimeout(context.Background(), app.StartTimeout())
	defer cancel()

	if err := app.Start(startCtx); err != nil {
		return err
	}

	sigs := app.Done()

	logrus.Info(fmt.Sprintf("received signal: %v", <-sigs))
	logrus.Info(fmt.Sprintf("exiting in %s", app.StartTimeout().String()))

	stopCtx, cancel := context.WithTimeout(context.Background(), app.StopTimeout())
	defer cancel()

	logrus.Info("stopping app")
	return app.Stop(stopCtx)
}
