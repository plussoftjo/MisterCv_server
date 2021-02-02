// Package controllers ..
package controllers

import (
	"net/http"
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
)

// StoreTemplate ..
func StoreTemplate(c *gin.Context) {
	var template models.Templates
	if err := c.ShouldBindJSON(&template); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&template)
	c.JSON(200, gin.H{
		"message": "Success",
	})
}
