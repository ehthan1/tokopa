package redis

import "github.com/google/wire"

func New() {

}

var ProviderSet = wire.NewSet(New)
