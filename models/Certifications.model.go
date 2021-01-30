// Package models ...
package models

import (
	"github.com/jinzhu/gorm"
)

// Certifications ...
type Certifications struct {
	UserID uint   `json:"user_id"`
	CvsID  uint   `json:"cvs_id"`
	Title  string `json:"title"`
	gorm.Model
}
