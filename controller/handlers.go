package controller

import (
	"idea-generator-api/idea-generator/database"
	"idea-generator-api/idea-generator/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllIdeas(c *gin.Context) {
	var ideas []models.Idea
	if err := database.DB.Find(&ideas).Error; err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, ideas)
}

func GetIdeaByID(c *gin.Context) {
	var idea models.Idea
	if err := database.DB.Where("id = ?", c.Param("id")).Find(&idea).Error; err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, idea)

}

func CreateIdea(c *gin.Context) {
	var idea models.Idea
	if err := c.ShouldBindJSON(&idea); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	database.DB.Create(&idea)
	c.JSON(http.StatusOK, &idea)

}

func UpdateIdea(c *gin.Context) {
	var idea models.Idea
	if err := database.DB.Where("id = ?", c.Param("id")).First(&idea).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error ": "idea Not Found"})
	}

	if err := c.ShouldBindJSON(&idea); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&idea).Updates(&idea).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"idea Updated": idea})
}

func DeleteIdea(c *gin.Context) {
	var idea models.Idea
	if err := database.DB.Where("id = ?", c.Param("id")).Delete(&idea).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "idea not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "idea deleted"})
}
