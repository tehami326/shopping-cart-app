package database

import (
	"fmt"
	"log"

	"shopping-cart-app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=Tehami@123 dbname=shopping_cart port=5432 sslmode=disable TimeZone=Asia/Kolkata"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	fmt.Println("✅ Connected to PostgreSQL")

	DB.AutoMigrate(
		&models.User{}, &models.Item{}, &models.Cart{}, &models.CartItem{}, &models.Order{}, &models.OrderItem{})

	SeedItems()
}
