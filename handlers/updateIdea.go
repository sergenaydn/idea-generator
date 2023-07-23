package handlers

import (
	"idea-generator/database"
	"idea-generator/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateIdea godoc
// @Summary Update Ideas
// @Description Updates Data of an Existing Idea
// @Param id path int true "idea ID"
// @Param EnterNameOwner body string true "Idea Name and Owner please"
// @Produce json
// @Tags Ideas
// @Success 200 {object} models.Idea
// @Failure 400 {object} nil
// @Router /idea/{id} [put]
func UpdateIdea(c *gin.Context) {
	// Declare a variable to hold an instance of the models.Idea struct.
	var ideas models.Idea

	// Use the GORM library to query the database and find the idea with the given ID.
	// The ID is obtained from the URL path parameter using c.Param("id").
	if err := database.DB.Where("id = ?", c.Param("id")).First(&ideas).Error; err != nil {
		// If the idea with the given ID is not found in the database, respond with a 404 Not Found status.
		c.JSON(http.StatusNotFound, gin.H{"Error ": "idea Not Found"})
		return
	}

	// Parse and bind the request JSON data to the ideas variable.
	// This will update the values of the existing idea based on the provided JSON data.
	if err := c.ShouldBindJSON(&ideas); err != nil {
		// If there's an error in parsing the JSON or binding it to the idea model, respond with a 400 Bad Request status.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the updated idea to the database using the Save method of the GORM library.
	// The Updates method is used to update the columns of the idea with the provided data.
	if err := database.DB.Save(&ideas).Updates(&ideas).Error; err != nil {
		// If there's an error while saving the updated idea to the database, respond with a 500 Internal Server Error status.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If the update was successful, respond with a 200 OK status and the updated idea in JSON format.
	c.JSON(http.StatusOK, gin.H{"idea Updated": ideas})
}
