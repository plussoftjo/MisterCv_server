// Package controllers ...
package controllers

import (
	"server/config"
	"server/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// StoreSummary ..
func StoreSummary(c *gin.Context) {
	var summary models.Summary
	if err := c.ShouldBindJSON(&summary); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&summary)
	c.JSON(200, gin.H{
		"summary": summary,
	})
}

// UpdateSummary ..
func UpdateSummary(c *gin.Context) {
	var summary models.Summary
	if err := c.ShouldBindJSON(&summary); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&summary)
	c.JSON(200, gin.H{
		"summary": summary,
	})
}
