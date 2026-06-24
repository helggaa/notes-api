package main

import (
	"os"

	"notes_api/configs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	configs.ConnectDatabase()
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "API is running",
		})
	})
	port := os.Getenv("APP_PORT")
	router.Run(":" + port)
}
