package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized, gin.H{
				"error": "user unauthorized",
			},
		)
	}

}
