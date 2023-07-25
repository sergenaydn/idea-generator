package handlers

import (
	"idea-generator/database"
	"idea-generator/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllIdeas godoc
// @Summary Get a list of all ideas
// @Description Returns the Data of All the Ideas
// @Tags Ideas
// @Accept json
// @Produce json
// @success 200 {array} models.Idea
// @Failure 400 {object} nil
// @Router /idea [get]
func GetAllIdeas(c *gin.Context) {
	// Declare a variable to hold a slice of models.Idea structs.
	var idea []models.Idea

	// Use the GORM library to query the database and retrieve all ideas.
	// The Find method is used to retrieve all records from the "ideas" table and store them in the "ideas" variable.
	if err := database.DB.Find(&idea).Error; err != nil {
		// If there was an error during the retrieval (e.g., database connection issue), respond with a 400 Bad Request status.
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	// If the retrieval was successful, respond with the list of ideas in JSON format with a 200 OK status.
	c.JSON(http.StatusOK, idea)
}
