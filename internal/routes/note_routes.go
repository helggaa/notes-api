package routes

import (
	"github.com/gin-gonic/gin"

	"notes_api/internal/handler"
	"notes_api/internal/middleware"
)

func NoteRoutes(router *gin.Engine, noteHandler *handler.NoteHandler) {

	api := router.Group("/api")

	authorized := api.Group("/notes")
	authorized.Use(middleware.JWTMiddleware())

	{
		authorized.POST("", noteHandler.CreateNote)
		authorized.GET("", noteHandler.GetAllNotes)
		authorized.GET("/:id", noteHandler.GetNoteByID)
		authorized.PUT("/:id", noteHandler.UpdateNote)
		authorized.DELETE("/:id", noteHandler.DeleteNote)
	}
}
