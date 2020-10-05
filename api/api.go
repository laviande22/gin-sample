package api

import (
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/ping", ping)
	}
}
