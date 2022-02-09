package main

import (
<<<<<<< HEAD
	"bitbucket.org/portgo/tokopa/pkg/app"
	"bitbucket.org/portgo/tokopa/pkg/config"
	"bitbucket.org/portgo/tokopa/pkg/database"
	"bitbucket.org/portgo/tokopa/pkg/logger"
	"github.com/google/wire"
=======
	"github.com/google/wire"
	"github.com/portgo/tokopa/pkg/app"
	"github.com/portgo/tokopa/pkg/config"
	"github.com/portgo/tokopa/pkg/database"
	"github.com/portgo/tokopa/pkg/logger"
>>>>>>> 3a2cd47 (feat(init): Init)
)

var providerSet = wire.NewSet(
	logger.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
)

func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
