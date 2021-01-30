// Package models ...
package models

import "github.com/jinzhu/gorm"

// BasicsInformations ..
type BasicsInformations struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	CvsID    uint   `json:"cvs_id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	WebSite  string `json:"website"`
	Address  string `json:"address"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Photo    string `json:"photo"`
	Gender   string `json:"gender"`
}
