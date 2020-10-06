package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string
	Text  string
}

// Migrate automigrates models using ORM
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Post{})
	fmt.Println("Auto Migration completed")
}
