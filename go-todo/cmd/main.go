package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "go-todo/database"
    "go-todo/routes"
)

func main() {
    // Connect to the database
    database.Connect()

    // Set up Gin router
    r := gin.Default()

    // Apply CORS middleware to allow frontend to make requests
    r.Use(cors.Default())

    // Set up routes
    routes.SetupRoutes(r)

    // Run the server
    r.Run(":8080")
}
