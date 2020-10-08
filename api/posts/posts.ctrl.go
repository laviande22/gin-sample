package posts

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/laviande22/gin-sample/database/models"
	"gorm.io/gorm"
)

func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		Title string `json:"title" binding:"required"`
		Text  string `json:"text" binding:"required"`
	}
	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println("asdf")
		c.AbortWithStatus(400)
		return
	}

	post := models.Post{Title: requestBody.Title, Text: requestBody.Text}
	db.Create(&post)
	c.JSON(200, post.Serialize())
}
