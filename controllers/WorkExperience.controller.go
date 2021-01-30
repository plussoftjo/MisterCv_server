// Package controllers ...
package controllers

import (
	"server/config"
	"server/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// StoreWorkExperience ..
func StoreWorkExperience(c *gin.Context) {
	var workExperience []models.WorkExperience
	if err := c.ShouldBindJSON(&workExperience); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, experience := range workExperience {
		if err := config.DB.Create(&experience).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(200, gin.H{
		"workExperience": workExperience,
	})
}

// UpdateWorkExperience ..
func UpdateWorkExperience(c *gin.Context) {
	var workExperience models.WorkExperience
	if err := c.ShouldBindJSON(&workExperience); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&workExperience)
	c.JSON(200, gin.H{
		"workExperience": workExperience,
	})
}

// DeleteWorkExperience ..
func DeleteWorkExperience(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.WorkExperience{}, id)

	c.JSON(200, gin.H{
		"message": "Deleted",
	})
}
