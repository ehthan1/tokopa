package main

import (
	"bitbucket.org/portgo/tokopa/pkg/app"
	"bitbucket.org/portgo/tokopa/pkg/config"
	"bitbucket.org/portgo/tokopa/pkg/database"
	"bitbucket.org/portgo/tokopa/pkg/logger"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	logger.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
)

func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
