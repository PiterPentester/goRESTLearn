package apiserver

import (
	"github.com/PiterPentester/goRESTLearn/internal/app/store"
)

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func CreateConfig() *Config {
	return &Config{
		BindAddr: ":3939",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
