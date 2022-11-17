package commands

import (
	"context"
	"world_generator_processor_service/internal/languages/entities"

	"github.com/olehmushka/postgres-golang"
)

type Subfamily interface {
	Create(ctx context.Context, sf *entities.Subfamily) error
	CreateTx(ctx context.Context, client postgres.Writer, sf *entities.Subfamily) error
}
