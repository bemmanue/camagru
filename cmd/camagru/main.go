package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/bemmanue/camagru/internal/app/camagru"
	_ "github.com/lib/pq"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/camagru.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := camagru.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatalln(err)
	}

	if err := camagru.Start(config); err != nil {
		log.Fatalln(err)
	}
}
