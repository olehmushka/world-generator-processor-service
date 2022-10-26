package cli

import (
	"world_generator_processor_service/config"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	config.Module,
)
