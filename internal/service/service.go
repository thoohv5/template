// Package service 业务
package service

import (
	ds "github.com/thoohv5/template/internal/data/standard"
	"github.com/thoohv5/template/internal/pkg/config"
)

type service struct {
	// 配置
	cf config.IConfig
	// 日志
	log logx.ILog

	// 数据源
	data ds.IData
}

// IService 业务标准
type IService interface {
}

// New 创建
func New(cf config.IConfig, log logx.ILog, data ds.IData) IService {
	return &service{
		cf:  cf,
		log: log,

		data: data,
	}
}
