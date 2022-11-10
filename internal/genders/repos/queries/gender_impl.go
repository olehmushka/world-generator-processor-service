package queries

import (
	"github.com/olehmushka/postgres-golang"
	"go.uber.org/fx"
)

type genderRepo struct {
	client postgres.ClientReader
}

func NewGenderRepo(client postgres.ClientReader) Gender {
	return &genderRepo{client: client}
}

var GenderModule = fx.Options(
	fx.Provide(NewGenderRepo),
)
