package cli

import (
	"world_generator_processor_service/config"
	"world_generator_processor_service/core/engine"
	"world_generator_processor_service/core/storage/mongo"
	"world_generator_processor_service/core/storage/postgres"
	"world_generator_processor_service/core/storage/redis"
	gendersEngineAdp "world_generator_processor_service/internal/genders/adapters/engine"
	gendersStorageAdp "world_generator_processor_service/internal/genders/adapters/storage"
	gendersReposComm "world_generator_processor_service/internal/genders/repos/commands"
	gendersReposQueries "world_generator_processor_service/internal/genders/repos/queries"
	gendersReposTx "world_generator_processor_service/internal/genders/repos/tx"
	gendersServices "world_generator_processor_service/internal/genders/services"
	languagesEngineAdp "world_generator_processor_service/internal/languages/adapters/engine"
	languagesStorageAdp "world_generator_processor_service/internal/languages/adapters/storage"
	languagesReposComm "world_generator_processor_service/internal/languages/repos/commands"
	languagesReposQueries "world_generator_processor_service/internal/languages/repos/queries"
	languagesReposTx "world_generator_processor_service/internal/languages/repos/tx"
	languagesServices "world_generator_processor_service/internal/languages/services"

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

	gendersReposComm.GenderModule,
	gendersReposComm.GenderAcceptaceModule,
	gendersReposQueries.GenderModule,
	gendersReposQueries.GenderAcceptaceModule,
	gendersReposTx.Module,
	gendersEngineAdp.Module,
	gendersStorageAdp.Module,
	gendersServices.GenderAcceptanceSrvModule,
	gendersServices.GenderSrvModule,
	gendersServices.SyncSrvModule,
	languagesReposComm.FamilyModule,
	languagesReposComm.SubfamilyModule,
	languagesReposComm.LanguageModule,
	languagesReposQueries.FamilyModule,
	languagesReposQueries.SubfamilyModule,
	languagesReposQueries.LanguageModule,
	languagesReposTx.Module,
	languagesEngineAdp.Module,
	languagesStorageAdp.Module,
	languagesServices.SyncSrvModule,
)
