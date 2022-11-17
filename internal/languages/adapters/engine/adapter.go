package engine

import (
	"context"
	"world_generator_processor_service/internal/languages/entities"

	"github.com/olehmushka/golang-toolkit/either"
)

type Adapter interface {
	LoadNativeFamilies(context.Context) chan either.Either[[]*entities.Family]
	LoadNativeSubfamilies(context.Context) chan either.Either[[]*entities.Subfamily]
	LoadNativeLanguages(context.Context) chan either.Either[*entities.Language]
}
