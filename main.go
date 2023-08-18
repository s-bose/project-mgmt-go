package main

import (
	middleware "github.com/s-bose/project-mgmt-go/app"
	"github.com/s-bose/project-mgmt-go/app/db"

	"github.com/gin-gonic/gin"
)

func main() {
	database := db.InitDatabase()

	r := gin.Default()
	r.Use(middleware.DbMiddleware(database.Db))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!",
		})
	})

	r.Run(":8080")
}
