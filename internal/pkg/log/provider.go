package log

import (
	"github.com/google/wire"
	
	"github.com/thoohv5/template/internal/pkg/config"
	"github.com/thoohv5/template/pkg/log"
)

// ProviderSet is log providers.
var ProviderSet = wire.NewSet(
	New,
)

func New(cf config.IConfig) log.ILog {
	return log.NewZap()
}
