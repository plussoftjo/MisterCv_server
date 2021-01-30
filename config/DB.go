// Package config ...
package config

import (
	"github.com/jinzhu/gorm"
	// Connect mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// models
	"server/models"
)

// SetupDB ...

// DB ..
var DB *gorm.DB

// SetupDB ..
func SetupDB() {
	database, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3307)/mistercv?charset=utf8mb4&parseTime=True&loc=Local")

	// If Error in Connect
	if err != nil {
		panic(err)
	}
	// User Models Setup
	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.AuthClients{})
	database.AutoMigrate(&models.AuthTokens{})
	database.AutoMigrate(&models.Cvs{})
	database.AutoMigrate(&models.Categories{})
	database.AutoMigrate(&models.BasicsInformations{})
	database.AutoMigrate(&models.Certifications{})
	database.AutoMigrate(&models.Educations{})
	database.AutoMigrate(&models.Skills{})
	database.AutoMigrate(&models.Summary{})
	database.AutoMigrate(&models.WorkExperience{})
	database.AutoMigrate(&models.Templates{})

	// Confirm the DB variables
	DB = database

}
