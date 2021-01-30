// Package models ...
package models

import (
	"github.com/jinzhu/gorm"
)

// Educations ..
type Educations struct {
	UserID    uint   `json:"user_id"`
	CvsID     uint   `json:"cvs_id"`
	Title     string `json:"title"`
	Colleague string `json:"colleague"`
	Start     string `json:"start"`
	End       string `json:"end"`
	gorm.Model
}
