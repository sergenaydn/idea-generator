package handlers

import (
	"idea-generator/database"
	"idea-generator/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteIdea godoc
// @Summary Delete an Idea
// @Description Deletes the Data of an Idea By it's ID
// @Tags Ideas
// @Accept json
// @Produce json
// @success 200 {object} models.Idea
// @Failure 400 {object} nil
// @Param id path int true "idea ID"
// @Router /idea/{id} [delete]
func DeleteIdea(c *gin.Context) {
	// Declare a variable to hold an instance of the models.Idea struct.
	var ideas models.Idea

	// Use the GORM library to query the database and delete the idea with the given ID.
	// The ID is obtained from the URL path parameter using c.Param("id").
	// The database.DB is the global variable that holds the database connection.
	// The Delete method performs the actual deletion based on the query.
	if err := database.DB.Where("id = ?", c.Param("id")).Delete(&ideas).Error; err != nil {
		// If there was an error during the deletion (e.g., idea not found), respond with an error JSON.
		c.JSON(http.StatusNotFound, gin.H{"error": "idea not found"})
		return
	}

	// If the deletion was successful, respond with a success JSON indicating the idea was deleted.
	c.JSON(http.StatusOK, gin.H{"message": "idea deleted"})
}
