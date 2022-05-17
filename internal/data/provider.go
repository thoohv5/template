// Package data 数据源注入
package data

import (
	"github.com/google/wire"

	"github.com/thoohv5/template/internal/data/dao"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	New,
	dao.NewTest,
)
