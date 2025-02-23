package handlers

import (
	"go-todo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetTodosHandler handles the GET request to fetch all todos
func GetTodosHandler(c *gin.Context) {
	todos, err := models.GetTodos() // Fetch todos from the database
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch todos"})
		return
	}
	c.JSON(http.StatusOK, todos) // Return the todos as a JSON response
}

// CreateTodoHandler handles the POST request to create a new todo
func CreateTodoHandler(c *gin.Context) {
	var newTodo models.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Create the todo in the database
	if err := newTodo.CreateTodo(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create todo"})
		return
	}
	c.JSON(http.StatusCreated, newTodo) // Return the newly created todo as JSON
}

// UpdateTodoStatusHandler updates the completed status of a todo
func UpdateTodoStatusHandler(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	// Bind the request body to the todo struct
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the todo status in the database
	err := todo.UpdateStatus(id, todo.Completed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo status"})
		return
	}

	c.JSON(http.StatusOK, todo)
}
// UpdateTodoTitleHandler updates the title of a todo
func UpdateTodoTitleHandler(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	// Bind the request body to the todo struct
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the todo title in the database
	err := todo.UpdateTitle(id, todo.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo title"})
		return
	}

	c.JSON(http.StatusOK, todo)
}
// DeleteTodoHandler deletes a todo from the database
func DeleteTodoHandler(c *gin.Context) {
	id := c.Param("id")

	err := models.DeleteTodo(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
// MarkAllTodosCompletedHandler marks all todos as completed
func MarkAllTodosCompletedHandler(c *gin.Context) {
	err := models.MarkAllTodosCompleted()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark all todos as completed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "All todos marked as completed"})
}
