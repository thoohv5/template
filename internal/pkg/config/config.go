package config

import (
	dbs "github.com/thoohv5/template/pkg/dbx/standard"
	"github.com/thoohv5/template/pkg/hpx"
	los "github.com/thoohv5/template/pkg/logx/standard"
	"github.com/thoohv5/template/pkg/rdx"
)

type config struct {
	// 服务
	Http *hpx.Config `toml:"http"`
	// 日志
	Log *los.Config `toml:"log"`
	// Database 配置
	Database *dbs.Config `toml:"database"`
	// redis 配置
	Redis *rdx.Config `toml:"redis"`
}

type IConfig interface {
	GetHttp() *hpx.Config
	GetDatabase() *dbs.Config
	GetRedis() *rdx.Config
	GetLog() *los.Config
}

func (c *config) GetHttp() *hpx.Config {
	return c.Http
}

func (c *config) GetDatabase() *dbs.Config {
	return c.Database
}

func (c *config) GetRedis() *rdx.Config {
	return c.Redis
}

func (c *config) GetLog() *los.Config {
	return c.Log
}
