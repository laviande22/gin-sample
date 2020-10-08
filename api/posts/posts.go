package posts

import (
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	posts := r.Group("/posts")
	{
		posts.POST("/", create)
		posts.GET("/", list)
		posts.GET("/:id", read)
	}
}
