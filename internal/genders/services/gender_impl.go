package services

import (
	"context"
	"world_generator_processor_service/core/storage"
	gendersStorageAdp "world_generator_processor_service/internal/genders/adapters/storage"
	gendersEntities "world_generator_processor_service/internal/genders/entities"

	"go.uber.org/fx"
)

type genderSrv struct {
	storageAdp gendersStorageAdp.Adapter
}

func NewGenderSrv(storageAdp gendersStorageAdp.Adapter) Gender {
	return &genderSrv{
		storageAdp: storageAdp,
	}
}

var GenderSrvModule = fx.Options(
	fx.Provide(NewGenderSrv),
)

func (s *genderSrv) QueryGenders(ctx context.Context, origin gendersEntities.Origin, paginate storage.Pagination) ([]gendersEntities.Gender, int, error) {
	genders, err := s.storageAdp.QueryGenders(ctx, origin, paginate)
	if err != nil {
		return nil, 0, err
	}

	count, err := s.storageAdp.CountGenders(ctx, origin)
	if err != nil {
		return nil, 0, err
	}

	return genders, count, nil
}
