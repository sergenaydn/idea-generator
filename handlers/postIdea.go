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
	var idea models.Idea

	// Parse and bind the request JSON data to the idea variable.
	// This will create a new idea object with the data provided in the JSON request.
	if err := c.ShouldBindJSON(&idea); err != nil {
		// If there's an error in parsing the JSON or binding it to the idea model, respond with a 400 Bad Request status.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if "idea" and "owner" fields are not empty.
	if idea.Idea == "" || idea.Owner == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Idea and owner cannot be empty"})
		return
	}

	// If there is a manually entered ID returns error.
	if idea.Id != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you cant enter id"})
		return
	}

	// Create a new record in the database using the database.DB instance.
	// The DB.Create method inserts the data from the "idea" variable into the "ideas" table in the database.
	database.DB.Create(&idea)

	// If the creation was successful, respond with a 200 OK status and the newly created idea in JSON format.
	c.JSON(http.StatusOK, &idea)
}
