package main

import (
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/laviande22/gin-sample/api"
)

func main() {
    // load .env environment variables
    err := godotenv.Load()
    if err != nil {
        panic(err)
    }

    port := os.Getenv("PORT")
    app := gin.Default()
    api.ApplyRoutes(app)
    app.Run(":" + port)
}
