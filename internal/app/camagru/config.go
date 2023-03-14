package camagru

import "github.com/bemmanue/camagru/internal/store"

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8888",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
