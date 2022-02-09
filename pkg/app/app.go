package app

import (
	"github.com/google/wire"
	"go.uber.org/zap"
)

type Application struct {
	name    string
	version string
	logger  *zap.Logger
<<<<<<< HEAD
	// httpServer *http.Server
=======
>>>>>>> 3a2cd47 (feat(init): Init)
}

func New() (*Application, error) {
	return &Application{name: "1", version: "1"}, nil
}

var ProviderSet = wire.NewSet(New)
