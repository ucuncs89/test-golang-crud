package main

import (
	"test-go/models"
	"test-go/routes"
)

func main() {
	db := models.ConfigDB()
	db.AutoMigrate(&models.Product{})

	r := routes.SetupRoutes(db)
	r.Run(":9000")
}
