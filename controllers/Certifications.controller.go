// Package controllers ...
package controllers

import (
	"server/config"
	"server/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// StoreCertifications ..
func StoreCertifications(c *gin.Context) {
	var certifications []models.Certifications
	if err := c.ShouldBindJSON(&certifications); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, certification := range certifications {
		if err := config.DB.Create(&certification).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(200, gin.H{
		"certifications": certifications,
	})
}

// UpdateCertifications ..
func UpdateCertifications(c *gin.Context) {
	var certifications models.Certifications
	if err := c.ShouldBindJSON(&certifications); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&certifications)
	c.JSON(200, gin.H{
		"certifications": certifications,
	})
}

// DeleteCertifications ..
func DeleteCertifications(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Certifications{}, id)

	c.JSON(200, gin.H{
		"message": "Deleted",
	})
}
