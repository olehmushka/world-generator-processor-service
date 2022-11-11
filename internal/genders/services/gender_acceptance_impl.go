package services

import (
	"context"
	"world_generator_processor_service/core/storage"
	gendersStorageAdp "world_generator_processor_service/internal/genders/adapters/storage"
	gendersEntities "world_generator_processor_service/internal/genders/entities"

	"go.uber.org/fx"
)

type genderAcceptaceSrv struct {
	storageAdp gendersStorageAdp.Adapter
}

func NewGenderAcceptanceSrv(storageAdp gendersStorageAdp.Adapter) GenderAcceptace {
	return &genderAcceptaceSrv{
		storageAdp: storageAdp,
	}
}

var GenderAcceptanceSrvModule = fx.Options(
	fx.Provide(NewGenderAcceptanceSrv),
)

func (s *genderAcceptaceSrv) QueryGenderAcceptances(ctx context.Context, origin gendersEntities.Origin, paginate storage.Pagination) ([]gendersEntities.GenderAcceptance, int, error) {
	genderAcceptances, err := s.storageAdp.QueryGenderAcceptances(ctx, origin, paginate)
	if err != nil {
		return nil, 0, err
	}

	count, err := s.storageAdp.CountGenderAcceptances(ctx, origin)
	if err != nil {
		return nil, 0, err
	}

	return genderAcceptances, count, nil
}
