package routes

import (
	"github.com/gin-gonic/gin"
	"go-todo/handlers"
)

// SetupRoutes initializes the routes for the API
func SetupRoutes(router *gin.Engine) {
	router.GET("/todos", handlers.GetTodosHandler)
	router.POST("/todos", handlers.CreateTodoHandler)
	router.PUT("/todos/:id", handlers.UpdateTodoStatusHandler)
	router.PUT("/todos/title/:id", handlers.UpdateTodoTitleHandler)   // Edit todo title
	router.DELETE("/todos/:id", handlers.DeleteTodoHandler)          // Delete todo
	router.PUT("/todos/complete", handlers.MarkAllTodosCompletedHandler) // Mark all as completed
}
