package main

import (
	"URLShortener/Config"
	"URLShortener/Routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := Config.InitDB()

	r := gin.Default()
	Routes.RegisterRoutes(r, db)

	r.Run(":8080")
}
