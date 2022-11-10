package services

import "go.uber.org/fx"

type acceptaceSrv struct{}

func NewAcceptanceSrv() Acceptace {
	return &acceptaceSrv{}
}

var AcceptanceSrvModule = fx.Options(
	fx.Provide(NewAcceptanceSrv),
)
