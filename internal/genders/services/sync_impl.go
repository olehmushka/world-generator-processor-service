package services

import (
	"context"
	"world_generator_processor_service/internal/genders/adapters/engine"
	"world_generator_processor_service/internal/genders/adapters/storage"
	"world_generator_processor_service/internal/genders/entities"

	"github.com/olehmushka/golang-toolkit/wrapped_error"
	"go.uber.org/fx"
)

type sync struct {
	engineAdp  engine.Adapter
	storageAdp storage.Adapter
}

func NewSyncSrv(
	engineAdp engine.Adapter,
	storageAdp storage.Adapter,
) Sync {
	return &sync{
		engineAdp:  engineAdp,
		storageAdp: storageAdp,
	}
}

var SyncSrvModule = fx.Options(
	fx.Provide(NewSyncSrv),
)

func (s *sync) SyncGenders(ctx context.Context) error {
	genders, err := s.engineAdp.LoadNativeGenders(ctx)
	if err != nil {
		return wrapped_error.NewInternalServerError(err, "can not load genders for sync")
	}
	for _, g := range genders {
		tx, err := s.storageAdp.BeginTx(ctx)
		if err != nil {
			return err
		}
		gottenGender, err := s.storageAdp.QueryGenderByNameTx(ctx, tx, g.Name)
		if err != nil {
			return err
		}
		if gottenGender != entities.ZeroGender {
			continue
		}
		if err := s.storageAdp.CreateGenderTx(ctx, tx, g); err != nil {
			return err
		}

		if err := tx.Commit(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (s *sync) SyncGenderAcceptances(ctx context.Context) error {
	genderAcceptances, err := s.engineAdp.LoadNativeGenderAcceptances(ctx)
	if err != nil {
		return wrapped_error.NewInternalServerError(err, "can not load gender_acceptances for sync")
	}
	for _, ga := range genderAcceptances {
		tx, err := s.storageAdp.BeginTx(ctx)
		if err != nil {
			return err
		}
		gottenGenderAcceptace, err := s.storageAdp.QueryGenderAcceptanceByNameTx(ctx, tx, ga.Name)
		if err != nil {
			return err
		}
		if gottenGenderAcceptace != entities.ZeroGenderAcceptance {
			continue
		}
		if err := s.storageAdp.CreateGenderAcceptanceTx(ctx, tx, ga); err != nil {
			return err
		}

		if err := tx.Commit(ctx); err != nil {
			return err
		}
	}

	return nil
}
