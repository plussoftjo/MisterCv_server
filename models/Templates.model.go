// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Templates ..
type Templates struct {
	Caption     string `json:"caption"`
	TemplateURI string `json:"template_uri"`
	Title       string `json:"title"`
	gorm.Model
}
