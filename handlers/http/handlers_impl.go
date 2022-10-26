package http

import (
	httpServer "world_generator_processor_service/core/http_server"

	"go.uber.org/fx"
)

type handlers struct {
}

func New() httpServer.Handlers {
	return &handlers{}
}

var Module = fx.Options(
	fx.Provide(New),
)
