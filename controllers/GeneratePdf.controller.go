// Package controllers ...
package controllers

import (
	"fmt"
	"math/rand"
	"server/config"
	"server/models"
	u "server/vendors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandStringRunes ..
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// GenerateCv ..
func GenerateCv(c *gin.Context) {

	id := c.Param("id")

	var cvs models.Cvs
	config.DB.Where("id = ?", id).First(&cvs)

	var template models.Templates
	config.DB.Where("id = ?", cvs.TemplateID).First(&template)

	r := u.NewRequestPdf("")

	//html template path
	templatePath := config.ServerInfo.PublicPath + "public/Templates/CvTemplates/" + template.TemplateURI

	fileName := id + "-" + strconv.FormatInt(int64(time.Now().Unix()), 10)
	//path for download pdf
	outputPath := config.ServerInfo.PublicPath + "public/PDF/" + fileName + ".pdf"

	var basicsInformations models.BasicsInformations
	var summary models.Summary
	var educations []models.Educations
	var workExperience []models.WorkExperience
	var skills []models.Skills
	var certifications []models.Certifications

	config.DB.Where("cvs_id = ?", id).Find(&educations)
	config.DB.Where("cvs_id = ?", id).Find(&workExperience)
	config.DB.Where("cvs_id = ?", id).Find(&skills)
	config.DB.Where("cvs_id = ?", id).Find(&certifications)
	config.DB.Where("cvs_id = ?", id).First(&basicsInformations)
	config.DB.Where("cvs_id = ?", id).First(&summary)

	// Check has Image
	var HasImage bool
	if basicsInformations.Photo != "" {
		HasImage = true
	} else {
		HasImage = false
	}

	//html template data
	templateData := struct {
		BasicsInformations models.BasicsInformations
		Summary            models.Summary
		Educations         []models.Educations
		WorkExperience     []models.WorkExperience
		Skills             []models.Skills
		Certifications     []models.Certifications
		HasImage           bool
		ServerURI          string
	}{
		BasicsInformations: basicsInformations,
		Summary:            summary,
		Educations:         educations,
		WorkExperience:     workExperience,
		Skills:             skills,
		Certifications:     certifications,
		HasImage:           HasImage,
		ServerURI:          config.ServerInfo.ServerURI,
	}

	if err := r.ParseTemplate(templatePath, templateData); err == nil {
		ok, _ := r.GeneratePDF(outputPath)
		fmt.Println(ok)
		config.DB.Model(&models.Cvs{}).Where("id = ?", id).Update("uri", fileName)
		config.DB.Model(&models.Cvs{}).Where("id = ?", id).Update("active", 1)
	} else {
		fmt.Println(err)
	}
	var cv models.Cvs
	config.DB.Model(&models.Cvs{}).Where("id = ?", id).
		Preload("BasicsInformations").
		Preload("Summary").
		Preload("Educations").
		Preload("WorkExperience").
		Preload("Skills").
		Preload("Certifications").First(&cv)
	c.JSON(200, gin.H{
		"message": "Success",
		"cv":      cv,
	})
}

// ServingHTML ...
func ServingHTML(c *gin.Context) {
	var basicsInformations models.BasicsInformations
	var summary models.Summary
	var educations []models.Educations
	var workExperience []models.WorkExperience
	var skills []models.Skills
	var certifications []models.Certifications

	config.DB.Where("cvs_id = ?", 1).Find(&educations)
	config.DB.Where("cvs_id = ?", 1).Find(&workExperience)
	config.DB.Where("cvs_id = ?", 1).Find(&skills)
	config.DB.Where("cvs_id = ?", 1).Find(&certifications)
	config.DB.Where("cvs_id = ?", 1).First(&basicsInformations)
	config.DB.Where("cvs_id = ?", 1).First(&summary)
	c.HTML(200, "template_1.html", gin.H{
		"BasicsInformations": basicsInformations,
		"Educations":         educations,
		"Summary":            summary,
		"WorkExperience":     workExperience,
		"Skills":             skills,
		"Certifications":     certifications,
	})
}

//ServingCV ..
// func ServingCV(c *gin.Context) {
// 	id := c.Param("id")
// 	fmt.Println("Called --------")
// 	var cvs models.Cvs
// 	// Get DB record
// 	config.DB.Where("id = ?", id).First(&cvs)
// 	// Fetch Uri
// 	uri := cvs.URI
// 	// Get File Link
// 	fileURI := "./public/PDF/" + uri + ".pdf"
// 	c.Header("Content-type", "application/pdf")
// 	c.File(fileURI)
// }
