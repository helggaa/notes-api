package main

import (
	"os"

	"notes_api/configs"
	"notes_api/internal/handler"
	"notes_api/internal/repository"
	"notes_api/internal/routes"
	"notes_api/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	configs.ConnectDatabase()

	router := gin.Default()

	// Repository

	userRepo := repository.NewUserRepository(configs.DB)
	noteRepo := repository.NewNoteRepository(configs.DB)

	// Service

	authService := service.NewAuthService(userRepo)
	noteService := service.NewNoteService(noteRepo)

	// Handler

	authHandler := handler.NewAuthHandler(authService)
	noteHandler := handler.NewNoteHandler(noteService)

	// Routes

	routes.AuthRoutes(router, authHandler)
	routes.NoteRoutes(router, noteHandler)

	router.GET("/health", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"status":  "success",
			"message": "API is running",
		})

	})

	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)

}
