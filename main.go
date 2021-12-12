package main

import (
	"github.com/ardanlabs/conf/v2"
	"github.com/ardanlabs/conf/v2/yaml"
	"os"
)

type Config struct {
	RequiredKey bool `conf:"required" yaml:"required"`
}

func main() {
	var cfg Config

	f, err := os.ReadFile("./config.yml")
	if err != nil {
		panic(err)
	}

	_, err = conf.Parse("", &cfg, yaml.WithData(f))
	if err != nil {
		panic(err)
	}
}
