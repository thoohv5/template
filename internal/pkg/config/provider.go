package config

import (
	"flag"

	"github.com/BurntSushi/toml"
	"github.com/google/wire"

	"github.com/thoohv5/template/pkg/util"
)

var (
	conf string
	// ProviderSet is config providers.
	ProviderSet = wire.NewSet(
		New,
	)
)

func init() {
	flag.StringVar(&conf, "conf", util.AbPath("../../../configs/config.toml"), "config path, eg: -conf config.yaml")
}

func New() IConfig {
	defaultConfig := new(config)
	_, err := toml.DecodeFile(conf, defaultConfig)
	if nil != err {
		panic(err)
	}
	return defaultConfig
}
