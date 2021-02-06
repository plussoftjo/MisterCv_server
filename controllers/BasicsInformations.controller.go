// Package controllers ...
package controllers

import (
	"path/filepath"
	"server/config"
	"server/models"
	"server/vendors"

	"net/http"

	"github.com/gin-gonic/gin"
)

// StoreBasicsInformations ..
func StoreBasicsInformations(c *gin.Context) {
	var basicsInformations models.BasicsInformations
	if err := c.ShouldBindJSON(&basicsInformations); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&basicsInformations)
	c.JSON(200, gin.H{
		"BasicsInformations": basicsInformations,
	})
}

// UpdateBasicsInformations ..
func UpdateBasicsInformations(c *gin.Context) {
	var basicsInformations models.BasicsInformations
	if err := c.ShouldBindJSON(&basicsInformations); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&basicsInformations)
	c.JSON(200, gin.H{
		"BasicsInformations": basicsInformations,
	})
}

// BasicsInformationsUploadImage ...
func BasicsInformationsUploadImage(c *gin.Context) {
	file, _ := c.FormFile("image")
	filename := filepath.Base(file.Filename)

	if err := c.SaveUploadedFile(file, config.ServerInfo.PublicPath+"public/Images/BasicsInformations/"+filename); err != nil {
		c.JSON(500, gin.H{
			"error":   err.Error(),
			"message": "error",
		})
		return
	}

	vendors.ResizeImage(filename, "min-"+filename, config.ServerInfo.PublicPath+"public/Images/BasicsInformations/")
	c.JSON(200, gin.H{
		"image":   "min-" + filename,
		"message": "success",
	})
}
