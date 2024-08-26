package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/s-bose/project-mgmt-go/app/services/auth"
)

func JWTMiddleware(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized, gin.H{
				"error": "user unauthorized",
			},
		)
	}

	if tokenString == "" {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, gin.H{
				"error": "Invalid token format",
			},
		)
	}

	authTokenParts := strings.Split(tokenString, " ")
	if len(authTokenParts) != 2 || authTokenParts[0] != "Bearer" {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized, gin.H{
				"error": "user unauthorized",
			},
		)
	}

	tokenString = authTokenParts[1]
	token, err := auth.ValidateJWTToken(tokenString)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			},
		)
	}

	claims := token.Claims.(*auth.ClaimsPayload)

	c.Set("userId", claims.UserID)
	c.Next()
}
