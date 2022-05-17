// Package service 业务注入
package service

import (
	"github.com/google/wire"

	"github.com/thoohv5/template/internal/service/example"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	New,
	example.New,
)
