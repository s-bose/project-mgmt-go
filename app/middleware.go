package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DbMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dbInstance", db)
		c.Next()
	}
}
