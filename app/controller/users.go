package controller

import (
	"net/http"

	users "github.com/s-bose/project-mgmt-go/app/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var userRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserController struct {
	UserService *users.UserService
}

func RegisterUser(c *gin.Context) {
	db, ok := c.MustGet("dbInstance").(*gorm.DB)
	if !ok {
		panic("Failed to inject Db")
	}

	userService := users.New(db)
	if err := c.Bind(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse request body",
		})

		return
	}

	user, err := userService.InsertUser(userRequest.Email, userRequest.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not create new user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}

func LoginUser(c *gin.Context, db *gorm.DB) {
	userService := users.New(db)

	if err := c.Bind(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse request body",
		})

		return
	}

	user, err := userService.GetUserByEmail(userRequest.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Email not found",
		})

		return
	}

	if userService.ValidateUser(user, userRequest.Password) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password does not match",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": *user,
	})
}
