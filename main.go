package main

import (
	"github.com/s-bose/project-mgmt-go/app/controller"
	"github.com/s-bose/project-mgmt-go/app/db"

	"github.com/gin-gonic/gin"
)

func main() {
	database := db.InitDatabase()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!",
		})
	})

	userController := controller.UserController{Db: database.Db}
	mainGroup := r.Group("/api")
	{
		userGroup := mainGroup.Group("/user")
		{
			userGroup.POST("/login", userController.LoginUser)
			userGroup.POST("/register", userController.RegisterUser)
		}

	}

	r.Run(":8080")
}
