// Package controllers ...
package controllers

import (
	"server/config"
	"server/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// StoreCv ..
func StoreCv(c *gin.Context) {
	var cvs models.Cvs
	if err := c.ShouldBindJSON(&cvs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&cvs)

	config.DB.Where("id = ?", cvs.ID).
		Preload("BasicsInformations").
		Preload("Summary").
		Preload("Educations").
		Preload("WorkExperience").
		Preload("Skills").
		Preload("Certifications").
		First(&cvs)
	c.JSON(200, gin.H{
		"cv": cvs,
	})
}

// IndexCv ...
func IndexCv(c *gin.Context) {
	var cvs []models.Cvs
	if err := config.DB.Preload("Summary").Find(&cvs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"cv": cvs,
	})
}

// UpdateCV ..
func UpdateCV(c *gin.Context) {
	var cvs models.Cvs
	if err := c.ShouldBindJSON(&cvs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&cvs)
	c.JSON(200, gin.H{
		"cv": cvs,
	})
}

// UpdateCvTemplate ..
func UpdateCvTemplate(c *gin.Context) {
	var cvs models.Cvs
	if err := c.ShouldBindJSON(&cvs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Model(&models.Cvs{}).Where("id = ?", c.Param("id")).Update("template_id", cvs.TemplateID)
}

// SearchInCvs ..
func SearchInCvs(c *gin.Context) {
	var searchCvModel models.SearchCvModel
	if err := c.ShouldBindJSON(&searchCvModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Country => Basics Informations
	// Categories => Cvs
	// Gender => Basics Informations

	SearchAny := false
	var cvsIDs2 []models.BasicsInformations
	if searchCvModel.AnyCountry == true && searchCvModel.Gender == "Any" {
		SearchAny = true
	} else if searchCvModel.AnyCountry == true && searchCvModel.Gender != "Any" {
		config.DB.Select("cvs_id").Where(&models.BasicsInformations{Gender: searchCvModel.Gender}, "Gender").
			Find(&cvsIDs2)
	} else if searchCvModel.AnyCountry == false && searchCvModel.Gender == "Any" {
		config.DB.Select("cvs_id").Where(&models.BasicsInformations{Country: searchCvModel.Country}, "Country").
			Find(&cvsIDs2)
	} else {
		config.DB.Select("cvs_id").Where(&models.BasicsInformations{Country: searchCvModel.Country, Gender: searchCvModel.Gender}, "Country", "Gender").
			Find(&cvsIDs2)
	}

	var cvsIDsCollections []int
	for _, key := range cvsIDs2 {
		cvsIDsCollections = append(cvsIDsCollections, int(key.CvsID))
	}
	var cvs []models.Cvs

	if SearchAny {
		// "categories_id = ?", searchCvModel.Category
		config.DB.Where(&models.Cvs{CategoriesID: searchCvModel.Category, Active: 1, Public: 1}).
			Preload("BasicsInformations").
			Preload("Summary").
			Preload("Educations").
			Preload("WorkExperience").
			Preload("Skills").
			Preload("Certifications").
			Find(&cvs)
	} else {
		config.DB.Where(&models.Cvs{CategoriesID: searchCvModel.Category, Active: 1, Public: 1}).
			Preload("BasicsInformations").
			Preload("Summary").
			Preload("Educations").
			Preload("WorkExperience").
			Preload("Skills").
			Preload("Certifications").
			Find(&cvs, cvsIDsCollections)
	}

	c.JSON(200, gin.H{
		"cvs": cvs,
	})

}
