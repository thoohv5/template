package adapter

import (
	los "github.com/thoohv5/template/pkg/logx/standard"
	"github.com/thoohv5/template/pkg/logx/zap"
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
