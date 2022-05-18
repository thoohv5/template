package log

import (
	"github.com/google/wire"

	"github.com/thoohv5/template/internal/pkg/config"
	loa "github.com/thoohv5/template/pkg/logx/adapter"
	los "github.com/thoohv5/template/pkg/logx/standard"
)

// ProviderSet is log providers.
var ProviderSet = wire.NewSet(
	New,
)

func New(cf config.IConfig) los.ILogger {
	return loa.New(cf.GetLog())
}
