package services

import "go.uber.org/fx"

type sexesSrv struct{}

func NewSexesSrv() Sexes {
	return &sexesSrv{}
}

var SexesSrvModule = fx.Options(
	fx.Provide(NewSexesSrv),
)
