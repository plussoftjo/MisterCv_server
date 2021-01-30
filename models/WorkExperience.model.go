// Package models ...
package models

import "github.com/jinzhu/gorm"

// WorkExperience ...
type WorkExperience struct {
	UserID  uint   `json:"user_id"`
	CvsID   uint   `json:"cvs_id"`
	Title   string `json:"title"`
	Company string `json:"company"`
	Start   string `json:"start"`
	End     string `json:"end"`
	gorm.Model
}
