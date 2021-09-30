package config

import (
	"flag"

	"github.com/BurntSushi/toml"
	"github.com/google/wire"
)

var (
	conf string
	// ProviderSet is config providers.
	ProviderSet = wire.NewSet(
		New,
	)
)

func init() {
	flag.StringVar(&conf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func New() IConfig {
	defaultConfig := new(config)
	_, _ = toml.DecodeFile(conf, defaultConfig)
	return defaultConfig
}
