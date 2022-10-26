package cli

import (
	httpServer "world_generator_processor_service/core/http_server"
	"world_generator_processor_service/core/zerodowntime"
	httpHandlers "world_generator_processor_service/handlers/http"

	"go.uber.org/fx"
)

const HTTPServerRunCommand = "http_server_run"

func runHTTPServerCommand() error {
	app := fx.New(
		fx.StartTimeout(defaultAppStartTimeout),
		fx.StopTimeout(defaultAppStopTimeout),
		fx.Options(
			Modules,

			httpHandlers.Module,
			httpServer.Module,
		),
	)

	return zerodowntime.HandleApp(app)
}
