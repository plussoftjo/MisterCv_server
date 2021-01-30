// Package controllers ...
package controllers

import (
	"fmt"
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
)

// MainIndex ..
func MainIndex(c *gin.Context) {
	user, err := AuthWithReturnUser(c.Request.Header["Authorization"][0])
	if err != nil {
		c.JSON(401, gin.H{
			"error": "UnAuthorized",
		})
		return
	}

	var categories []models.Categories
	config.DB.Find(&categories)

	var cvs []models.Cvs
	config.DB.Where("user_id = ?", user.ID).
		Preload("BasicsInformations").
		Preload("Summary").
		Preload("Educations").
		Preload("WorkExperience").
		Preload("Skills").
		Preload("Certifications").
		Find(&cvs)

	var templates []models.Templates
	config.DB.Find(&templates)

	c.JSON(200, gin.H{
		"user":       user,
		"categories": categories,
		"cvs":        cvs,
		"templates":  templates,
	})
}

// ServingImage ..
func ServingImage(c *gin.Context) {
	img := c.Param("img")
	fmt.Println(img)
	c.File(img)
}
