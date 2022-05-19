package adapter

import (
	"github.com/thoohv5/template/pkg/logx/impl/zap"
	los "github.com/thoohv5/template/pkg/logx/standard"
)

func GetAdapter(name string, config *los.Config) los.ILogger {
	var logger los.ILogger
	switch name {
	default:
		logger = zap.New(config)
	}
	return logger
}

func New(config *los.Config) los.ILogger {
	return GetAdapter("zap", config)
}
