package main

import (
	"URLShortener/Config"
	"URLShortener/Routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cfg := Config.LoadConfig()

	db := Config.InitDB(cfg)

	r := gin.Default()
	Routes.RegisterRoutes(r, db)

	r.Run(":" + cfg.AppPort)
}
