// Package controllers ...
package controllers

import (
	"server/config"
	"server/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// StoreSkills ..
func StoreSkills(c *gin.Context) {
	var skills []models.Skills
	if err := c.ShouldBindJSON(&skills); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, skill := range skills {
		if err := config.DB.Create(&skill).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(200, gin.H{
		"skills": skills,
	})
}

// UpdateSkills ...
func UpdateSkills(c *gin.Context) {
	var skills models.Skills
	if err := c.ShouldBindJSON(&skills); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&skills)
	c.JSON(200, gin.H{"skills": skills})

}

// DeleteSkill ...
func DeleteSkill(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Skills{}, id)

	c.JSON(200, gin.H{
		"message": "Deleted",
	})
}
