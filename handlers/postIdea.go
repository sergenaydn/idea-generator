package handlers

import (
	"idea-generator/database"
	"idea-generator/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostIdea godoc
// @Summary Create Ideas
// @Description Creates New Data of a New Idea
// @Param EnterNameOwner body string true "Idea Name and Owner please"
// @Produce application/json
// @Tags Ideas
// @Success 200 {object} models.Idea
// @Router /idea [post]
func PostIdea(c *gin.Context) {
	// Declare a variable to hold an instance of the models.Idea struct.
	var ideas models.Idea

	// Parse and bind the request JSON data to the ideas variable.
	// This will create a new idea object with the data provided in the JSON request.
	if err := c.ShouldBindJSON(&ideas); err != nil {
		// If there's an error in parsing the JSON or binding it to the idea model, respond with a 400 Bad Request status.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new record in the database using the database.DB instance.
	// The DB.Create method inserts the data from the "ideas" variable into the "ideas" table in the database.
	database.DB.Create(&ideas)

	// If the creation was successful, respond with a 200 OK status and the newly created idea in JSON format.
	c.JSON(http.StatusOK, &ideas)
}
