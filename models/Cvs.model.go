// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Cvs ..
type Cvs struct {
	UserID             uint               `json:"user_id" binding:"required"`
	Title              string             `json:"title" binding:"required"`
	Keys               string             `json:"keys"`
	Public             int                `json:"public"`
	CategoriesID       uint               `json:"categories_id" binding:"required"`
	Note               string             `json:"note" `
	TemplateID         uint               `json:"template_id"`
	Active             int                `json:"active"`
	URI                string             `json:"uri"`
	Categories         Categories         `json:"categories" gorm:"foreignKey:CategoriesID;references:ID"`
	BasicsInformations BasicsInformations `json:"basic_informations" gorm:"foreignKey:CvsID;references:ID"`
	Certifications     []Certifications   `json:"certifications"`
	Educations         []Educations       `json:"educations"`
	Skills             []Skills           `json:"skills"`
	Summary            Summary            `json:"summary"`
	WorkExperience     []WorkExperience   `json:"work_experience"`
	gorm.Model
}

// SearchCvModel ..
type SearchCvModel struct {
	Category   uint   `json:"category"`
	Country    string `json:"country"`
	Gender     string `json:"gender"`
	AnyCountry bool   `json:"any_country"`
}
