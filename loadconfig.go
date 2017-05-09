package k8guardlibs

import (
	"github.com/caarlos0/env"
	"github.com/k8guard/k8guardlibs/config"
)

var Cfg = config.Config{}

func LoadConfig() (config.Config, error) {
	err := env.Parse(&Cfg)
	return Cfg, err
}
