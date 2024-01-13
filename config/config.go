package config

import (
	"github.com/naoina/toml"
	"os"
)

type Config struct {
	MongoDB struct {
		Uri string
	}
}

func NewConfig(filePath string) *Config {
	c := new(Config)

	if file, err := os.Open(filePath); err != nil {
		panic(err)
	} else if err := toml.NewDecoder(file).Decode(c); err != nil {
		panic(err)
	} else {
		return c
	}
}
