package services

import (
	"context"
	"world_generator_processor_service/internal/languages/adapters/engine"
	"world_generator_processor_service/internal/languages/adapters/storage"

	"github.com/olehmushka/golang-toolkit/wrapped_error"
	"go.uber.org/fx"
)

type syncSrv struct {
	engineAdp  engine.Adapter
	storageAdp storage.Adapter
}

func NewSyncSrv(engineAdp engine.Adapter, storageAdp storage.Adapter) Sync {
	return &syncSrv{
		engineAdp:  engineAdp,
		storageAdp: storageAdp,
	}
}

var SyncSrvModule = fx.Options(
	fx.Provide(NewSyncSrv),
)

func (s *syncSrv) SyncFamilies(ctx context.Context) error {
	for familiesChunk := range s.engineAdp.LoadNativeFamilies(ctx) {
		if err := familiesChunk.Err; err != nil {
			return wrapped_error.NewInternalServerError(err, "can not load native language families")
		}

		for _, f := range familiesChunk.Value {
			tx, err := s.storageAdp.BeginTx(ctx)
			if err != nil {
				return err
			}
			gottenFamily, err := s.storageAdp.QueryFamilyBySlugTx(ctx, tx, f.Slug)
			if err != nil {
				if err := tx.Rollback(ctx); err != nil {
					return err
				}
				return err
			}
			if gottenFamily != nil {
				if err := tx.Commit(ctx); err != nil {
					return err
				}
				continue
			}
			if err := s.storageAdp.CreateFamilyTx(ctx, tx, f); err != nil {
				if err := tx.Rollback(ctx); err != nil {
					return err
				}
				return err
			}

			if err := tx.Commit(ctx); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *syncSrv) SyncLanguages(ctx context.Context) error {
	for languageChunk := range s.engineAdp.LoadNativeLanguages(ctx) {
		if err := languageChunk.Err; err != nil {
			return wrapped_error.NewInternalServerError(err, "can not load native language")
		}

		tx, err := s.storageAdp.BeginTx(ctx)
		if err != nil {
			return err
		}
		gottenLang, err := s.storageAdp.QueryLanguageBySlugTx(ctx, tx, languageChunk.Value.Slug)
		if err != nil {
			if err := tx.Rollback(ctx); err != nil {
				return err
			}
			return err
		}
		if gottenLang != nil {
			if err := tx.Commit(ctx); err != nil {
				return err
			}
			continue
		}
		if err := s.storageAdp.CreateLanguageTx(ctx, tx, languageChunk.Value); err != nil {
			if err := tx.Rollback(ctx); err != nil {
				return err
			}
			return err
		}

		if err := tx.Commit(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (s *syncSrv) SyncSubfamilies(ctx context.Context) error {
	for subfamiliesChunk := range s.engineAdp.LoadNativeSubfamilies(ctx) {
		if err := subfamiliesChunk.Err; err != nil {
			return wrapped_error.NewInternalServerError(err, "can not load native language subfamilies")
		}

		for _, sf := range subfamiliesChunk.Value {
			tx, err := s.storageAdp.BeginTx(ctx)
			if err != nil {
				return err
			}
			gottenSubfamily, err := s.storageAdp.QuerySubfamilyBySlugTx(ctx, tx, sf.Slug)
			if err != nil {
				if err := tx.Rollback(ctx); err != nil {
					return err
				}
				return err
			}
			if gottenSubfamily != nil {
				if err := tx.Commit(ctx); err != nil {
					return err
				}
				continue
			}
			if err := s.storageAdp.CreateSubfamilyTx(ctx, tx, sf); err != nil {
				if err := tx.Rollback(ctx); err != nil {
					return err
				}
				return err
			}

			if err := tx.Commit(ctx); err != nil {
				return err
			}
		}
	}

	return nil
}
