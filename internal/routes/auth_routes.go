package routes

import (
	"github.com/gin-gonic/gin"

	"notes_api/internal/handler"
	"notes_api/internal/middleware"
)

func AuthRoutes(router *gin.Engine, authHandler *handler.AuthHandler) {

	api := router.Group("/api")

	// Public routes
	api.POST("/register", authHandler.Register)
	api.POST("/login", authHandler.Login)

	// Protected routes
	authorized := api.Group("/")
	authorized.Use(middleware.JWTMiddleware())

	authorized.POST("/logout", authHandler.Logout)
}
