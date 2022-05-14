package config

import (
	_ "embed"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
)

type Config struct {
	Mysql Mysql `toml:"mysql"`
	Redis Redis `toml:"redis"`
}

type Mysql struct {
	DSN string `toml:"dsn" env:"MYSQL_DSN"`
}

type Redis struct {
	DSN string `toml:"dsn" env:"REDIS_DSN"`
}

//go:embed config.toml
var configToml string

func Parse() (conf Config, err error) {
	if _, err := toml.Decode(configToml, &conf); err != nil {
		return conf, errors.Wrap(err, "failed to parse toml config")
	}
	if err := env.Parse(&conf); err != nil {
		return conf, errors.Wrap(err, "failed to parse env config")
	}
	return
}
