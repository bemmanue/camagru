package main

import (
	"github.com/bemmanue/camagru/internal/app/camagru"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("no .env file found")
	}
}

func main() {
	config, err := camagru.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(config)

	if err := camagru.Start(config); err != nil {
		log.Fatalln(err)
	}
}
