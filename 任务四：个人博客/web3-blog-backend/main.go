package main

import (
	"web3-blog-backend/config"
	"web3-blog-backend/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", config.DB)
	})
	routers.SetupRoutes(r)
	r.Run(":8080")
}
