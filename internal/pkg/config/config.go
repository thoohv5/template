package config

import (
	"github.com/thoohv5/template/internal/pkg/dbx"
	"github.com/thoohv5/template/pkg/hpx"
	"github.com/thoohv5/template/pkg/rdx"
)

type config struct {
	// 服务
	Http *hpx.Config `toml:"http"`
	// Database 配置
	Database *dbx.Configs `toml:"database"`
	// redis 配置
	Redis *rdx.Config `toml:"redis"`
}

type IConfig interface {
	GetHttp() *hpx.Config
	GetDatabase() *dbx.Configs
	GetRedis() *rdx.Config
}

func (c *config) GetHttp() *hpx.Config {
	return c.Http
}

func (c *config) GetDatabase() *dbx.Configs {
	return c.Database
}

func (c *config) GetRedis() *rdx.Config {
	return c.Redis
}