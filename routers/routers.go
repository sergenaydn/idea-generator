package routers

import (
	"idea-generator/handlers"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// IdeaRouter sets up the routes for the "idea" endpoint.
func IdeaRouter(r *gin.Engine) {
	// Setup the route for Swagger documentation.
	// This serves the Swagger UI to visualize and interact with the API documentation.
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Create a new route group for the "/idea" endpoint.
	iR := r.Group("/idea")

	// Define the individual API routes and their corresponding handler functions.
	// The handlers are from the "handlers" package, which contains the logic for each endpoint.
	iR.GET("", handlers.GetAllIdeas)      // Get a list of all ideas
	iR.GET(":id", handlers.GetIdeaByID)   // Get an idea by its ID
	iR.POST("", handlers.PostIdea)        // Create a new idea
	iR.DELETE(":id", handlers.DeleteIdea) // Delete an idea by its ID
	iR.PUT(":id", handlers.UpdateIdea)    // Update an idea by its ID
}
