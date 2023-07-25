package handlers

import (
	"idea-generator/database"
	"idea-generator/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetIdeaByID godoc
// @Summary Get An Idea by ID
// @Description Returns the Data of an Idea By it's ID
// @Param id path int true "idea ID"
// @Tags Ideas
// @Accept json
// @Produce json
// @Success 200 {object} models.Idea
// @Failure 400 {object} nil
// @Router /idea/{id} [get]
func GetIdeaByID(c *gin.Context) {
	// Declare a variable to hold an instance of the Idea model
	var ideas models.Idea

	// Use the database package to query the database for an idea with the given ID
	// The `c.Param("id")` extracts the "id" parameter from the request URL
	// The retrieved idea will be stored in the `ideas` variable
	if err := database.DB.Where("id = ?", c.Param("id")).First(&ideas).Error; err != nil {
		// If there was an error during the database query, respond with a 400 Bad Request status code and an empty response
		c.JSON(http.StatusBadRequest, "Cannot find idea")
		return
	}

	// If the idea was found, respond with a 200 OK status code and the idea data in JSON format
	c.JSON(http.StatusOK, ideas)
}
