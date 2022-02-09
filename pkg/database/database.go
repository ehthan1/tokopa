package database

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	return &gorm.DB{}, nil
}

var ProviderSet = wire.NewSet(New)
