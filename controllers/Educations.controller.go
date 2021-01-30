// Package controllers ...
package controllers

import (
	"server/config"
	"server/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// StoreEducations ..
func StoreEducations(c *gin.Context) {
	var educations []models.Educations
	if err := c.ShouldBindJSON(&educations); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, key := range educations {
		if err := config.DB.Create(&key).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "educations": educations})
			return
		}
	}

	c.JSON(200, gin.H{
		"educations": educations,
	})
}

// UpdateEducations ..
func UpdateEducations(c *gin.Context) {
	var educations models.Educations
	if err := c.ShouldBindJSON(&educations); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&educations)
	c.JSON(200, gin.H{
		"educations": educations,
	})
}

// DeleteEducation ..
func DeleteEducation(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Educations{}, id)

	c.JSON(200, gin.H{
		"message": "Deleted",
	})
}
