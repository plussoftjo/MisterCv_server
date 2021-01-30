// Package controllers ...
package controllers

import (
	"server/config"
	"server/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// StoreCategories ..
func StoreCategories(c *gin.Context) {
	var categories models.Categories
	if err := c.ShouldBindJSON(&categories); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&categories)
	c.JSON(200, gin.H{
		"cv": categories,
	})
}

// IndexCategories ..
func IndexCategories(c *gin.Context) {
	var categories []models.Categories
	if err := config.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"categories": categories,
	})
}

// UpdateCategories ..
func UpdateCategories(c *gin.Context) {
	var categories models.Categories
	if err := c.ShouldBindJSON(&categories); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&categories)
	c.JSON(200, gin.H{
		"cv": categories,
	})
}
