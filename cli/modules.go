package cli

import (
	"world_generator_processor_service/config"
	"world_generator_processor_service/core/engine"
	"world_generator_processor_service/core/storage/mongo"
	"world_generator_processor_service/core/storage/postgres"
	"world_generator_processor_service/core/storage/redis"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	config.Module,
	engine.Module,
	mongo.Module,
	redis.StoreModule,
	redis.PublisherModule,
	postgres.Module,
	postgres.ReaderModule,
	postgres.WriterModule,
)
