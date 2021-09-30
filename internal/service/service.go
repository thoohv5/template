// Package service 业务
package service

import (
	"github.com/thoohv5/template/internal/data"
	"github.com/thoohv5/template/internal/pkg/config"
	"github.com/thoohv5/template/pkg/log"
)

type service struct {
	// 配置
	cf config.IConfig
	// 日志
	log log.ILog

	// 数据源
	data data.IData
}

// IService 业务标准
type IService interface{}

// New 创建
func New(cf config.IConfig, log log.ILog, data data.IData) IService {
	return &service{
		cf:  cf,
		log: log,

		data: data,
	}
}
