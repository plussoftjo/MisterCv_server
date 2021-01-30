// Package app ...
package app

import (
	"server/config"
	"server/routes"
	"server/vendors"
)

// Installing ..
func Installing() {
	// Setup The DB
	config.SetupDB()

	// SetupPassport
	vendors.SetupPassport()

	// setup routes
	routes.Setup()
}
