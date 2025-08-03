package database

import (
	"shopping-cart-app/models"
)

// SeedItems clears the table and adds default items
func SeedItems() {
	// Truncate the table and reset IDs
	DB.Exec("TRUNCATE TABLE items RESTART IDENTITY CASCADE;")

	// Seed fresh items
	items := []models.Item{
		{Name: "Apple", Price: 10.0, Status: "Available"},
		{Name: "Banana", Price: 5.0, Status: "Available"},
		{Name: "Orange", Price: 7.5, Status: "Available"},
	}

	for _, item := range items {
		DB.Create(&item)
	}
}
