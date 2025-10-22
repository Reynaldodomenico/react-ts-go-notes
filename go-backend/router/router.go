package router

import (
	"github.com/gin-gonic/gin"
	"github.com/reynaldodomenico/react-and-go-notes/go-backend/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/notes", handlers.GetNotes)
		api.POST("/notes", handlers.CreateNote)
		api.DELETE("/notes/:id", handlers.DeleteNote)
	}

	return r
}
