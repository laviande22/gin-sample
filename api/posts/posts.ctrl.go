package posts

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/laviande22/gin-sample/common"
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
	c.JSON(201, post.Serialize())
}

func list(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var posts []models.Post
	result := db.Find(&posts)

	post_count := result.RowsAffected
	serialized := make([]common.JSON, post_count, post_count)

	for i := 0; i < int(post_count); i++ {
		serialized[i] = posts[i].Serialize()
	}

	c.JSON(200, serialized)
}

func read(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var post models.Post
	if err := db.Where("id = ?", id).First(&post).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, post.Serialize())
}
