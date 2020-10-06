package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/laviande22/gin-sample/api"
	"github.com/laviande22/gin-sample/database"
)

func main() {
	// load .env environment variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// initialized database
	db, _ := database.Initialize()

	port := os.Getenv("PORT")
	app := gin.Default()
	app.Use(database.Inject(db))
	api.ApplyRoutes(app)
	app.Run(":" + port)
}
