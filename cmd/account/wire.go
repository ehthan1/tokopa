package main

import (
	"github.com/google/wire"
	"github.com/portgo/tokopa/pkg/app"
	"github.com/portgo/tokopa/pkg/config"
	"github.com/portgo/tokopa/pkg/database"
	"github.com/portgo/tokopa/pkg/logger"
)

var providerSet = wire.NewSet(
	logger.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
)

func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
