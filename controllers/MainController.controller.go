// Package controllers ...
package controllers

import (
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

	var exampleCv models.Cvs
	config.DB.Where("id = ?", 13).First(&exampleCv)

	c.JSON(200, gin.H{
		"user":       user,
		"categories": categories,
		"cvs":        cvs,
		"templates":  templates,
		"exampleCv":  exampleCv,
	})
}

// IndexWithoutAuth ...
func IndexWithoutAuth(c *gin.Context) {
	var categories []models.Categories
	config.DB.Find(&categories)

	var templates []models.Templates
	config.DB.Find(&templates)

	var exampleCv models.Cvs
	config.DB.Where("id = ?", 13).First(&exampleCv)

	c.JSON(200, gin.H{
		"categories": categories,
		"templates":  templates,
		"exampleCv":  exampleCv,
	})
}
