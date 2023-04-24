package main

import (
	"github.com/bemmanue/camagru/internal/app/camagru"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("no .env file found")
	}

	gin.SetMode(gin.ReleaseMode)
}

func main() {
	config, err := camagru.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	if err := camagru.Start(config); err != nil {
		log.Fatalln(err)
	}
}
