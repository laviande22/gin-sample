package database

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// This function injects the instance of database to each Gin context
func Inject(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
