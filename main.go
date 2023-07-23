package main

import (
	"idea-generator/database"
	docs "idea-generator/docs"
	"idea-generator/routers"

	"github.com/gin-gonic/gin"
)

// @title Idea Generator
// @version 1.0
// @description Generate New Ideas and List Existing Ones

// @host localhost:8080
// @BasePath /
func main() {
	// Create a new Gin router instance.
	r := gin.New()

	// Set up the Swagger documentation configuration.
	// This provides information about the API, such as its title, version, and description.
	docs.SwaggerInfo.BasePath = "/"

	// Attach the routes and handlers for the "idea" endpoint to the Gin router.
	// The routers.IdeaRouter function is responsible for defining these routes.
	routers.IdeaRouter(r)

	// Initialize the database connection.
	database.Init()

	// Run the Gin web server on port 8080.
	r.Run(":8080")
}
