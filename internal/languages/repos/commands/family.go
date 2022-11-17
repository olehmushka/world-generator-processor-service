package commands

import (
	"context"
	"world_generator_processor_service/internal/languages/entities"

	"github.com/olehmushka/postgres-golang"
)

type Family interface {
	Create(ctx context.Context, f *entities.Family) error
	CreateTx(ctx context.Context, client postgres.Writer, f *entities.Family) error
}
