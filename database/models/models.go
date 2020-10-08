package models

import (
	"fmt"

	"github.com/laviande22/gin-sample/common"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string
	Text  string
}

// Serializer for the model Post into JSON format
func (p Post) Serialize() common.JSON {
	return common.JSON{
		"id":         p.ID,
		"title":      p.Title,
		"text":       p.Text,
		"created_at": p.CreatedAt,
	}
}

// Migrate automigrates models using ORM
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Post{})
	fmt.Println("Auto Migration completed")
}
