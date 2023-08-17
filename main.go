package main

import (
	middleware "github.com/s-bose/project-mgmt-go/app"
	"github.com/s-bose/project-mgmt-go/app/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	db := Init()

	r := gin.Default()
	r.Use(middleware.DbMiddleware(db))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!",
		})
	})

	r.Run(":8080")
}

func Init() *gorm.DB {
	godotenv.Load()
	database := db.Database{}
	database.ConnectDatabase()
	database.Migrate()
	return database.Db
}
