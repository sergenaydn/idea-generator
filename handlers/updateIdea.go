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
// @Success 201 {object} models.Idea
// @Failure 400 {object} nil
// @Router /idea/{id} [put]
func UpdateIdea(c *gin.Context) {
	// Declare a variable to hold an instance of the models.Idea struct.
	var idea models.Idea
	var upReq models.UpdateReq

	// Use the GORM library to query the database and find the idea with the given ID.
	// The ID is obtained from the URL path parameter using c.Param("id").
	if err := database.DB.Where("id = ?", c.Param("id")).First(&idea).Error; err != nil {
		// If the idea with the given ID is not found in the database, respond with a 404 Not Found status.
		c.JSON(http.StatusNotFound, gin.H{"Error ": "idea Not Found"})
		return
	}

	// Parse and bind the request JSON data to the idea variable.
	// This will update the values of the existing idea based on the provided JSON data.
	if err := c.ShouldBindJSON(&upReq); err != nil {
		// If there's an error in parsing the JSON or binding it to the idea model, respond with a 400 Bad Request status.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if "idea," and "owner" fields are not empty.
	if upReq.Idea == "" || upReq.Owner == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID, idea, and owner cannot be empty"})
		return
	}

	idea.Idea = upReq.Idea
	idea.Owner = upReq.Idea

	// Save the updated idea to the database using the Save method of the GORM library.
	// The Updates method is used to update the columns of the idea with the provided data.
	if err := database.DB.Save(&idea).Updates(&idea).Error; err != nil {
		// If there's an error while saving the updated idea to the database, respond with a 500 Internal Server Error status.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If the update was successful, respond with a 200 OK status and the updated idea in JSON format.
	c.JSON(http.StatusCreated, gin.H{"idea Updated": idea})
}
