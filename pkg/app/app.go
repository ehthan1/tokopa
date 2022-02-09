package app

import (
	"github.com/google/wire"
	"go.uber.org/zap"
)

type Application struct {
	name    string
	version string
	logger  *zap.Logger
}

func New() (*Application, error) {
	return &Application{name: "1", version: "1"}, nil
}

var ProviderSet = wire.NewSet(New)
