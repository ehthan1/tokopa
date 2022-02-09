package logger

import (
	"github.com/google/wire"
	"go.uber.org/zap"
)

func New() (*zap.Logger, error) {
	return &zap.Logger{}, nil
}

var ProviderSet = wire.NewSet(New)
