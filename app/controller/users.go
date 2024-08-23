package controller

import (
	"fmt"
	"net/http"

	"github.com/s-bose/project-mgmt-go/app/services/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const MAX_AGE = 60 * 60 * 24 * 30 // 30 days
type UserController struct {
	Db *gorm.DB
}

func (u *UserController) RegisterUser(c *gin.Context) {

	var userRequest struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	userService := users.CreateUserService(u.Db)
	if err := c.ShouldBindJSON(&userRequest); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	user, err := userService.InsertUser(userRequest.Name, userRequest.Email, userRequest.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not create new user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": *user,
	})

}

func (u *UserController) LoginUser(c *gin.Context) {

	var userRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	userService := users.CreateUserService(u.Db)

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

	// return jwt token as cookie
	tokenString, err := users.CreateAccessToken((user.ID).String())
	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err,
			},
		)
		return
	}

	tokenStr, _ := tokenString["access_token"]

	tokenStr = fmt.Sprintf("Bearer %s", tokenStr)
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenStr, MAX_AGE, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"data": *user,
	})
}
