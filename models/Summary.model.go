// Package models ...
package models

import (
	"github.com/jinzhu/gorm"
)

// Summary ..
type Summary struct {
	UserID uint   `json:"user_id"`
	CvsID  uint   `json:"cvs_id"`
	Title  string `json:"title"`
	gorm.Model
}
