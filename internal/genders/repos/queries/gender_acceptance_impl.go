package queries

import (
	"github.com/olehmushka/postgres-golang"
	"go.uber.org/fx"
)

type genderAcceptanceRepo struct {
	client postgres.ClientReader
}

func NewGenderAcceptanceRepo(client postgres.ClientReader) GenderAcceptance {
	return &genderAcceptanceRepo{client: client}
}

var GenderAcceptaceModule = fx.Options(
	fx.Provide(NewGenderAcceptanceRepo),
)
