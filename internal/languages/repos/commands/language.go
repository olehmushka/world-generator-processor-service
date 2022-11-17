package commands

import (
	"context"
	"world_generator_processor_service/internal/languages/entities"

	"github.com/olehmushka/postgres-golang"
)

type Language interface {
	Create(ctx context.Context, lang *entities.Language) error
	CreateTx(ctx context.Context, client postgres.Writer, lang *entities.Language) error
}
