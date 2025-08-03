package controllers

import (
	"net/http"

	"shopping-cart-app/database"
	"shopping-cart-app/models"

	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDVal.(uint)

	var input struct {
		ItemID   uint `json:"item_id"`
		Quantity int  `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var cart models.Cart
	if err := database.DB.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		// Create cart if it doesn't exist
		cart = models.Cart{UserID: userID}
		database.DB.Create(&cart)
	}

	// Check if item is already in cart
	var existingItem models.CartItem
	result := database.DB.Where("cart_id = ? AND item_id = ?", cart.ID, input.ItemID).First(&existingItem)
	if result.RowsAffected > 0 {
		// If already exists, update quantity
		existingItem.Quantity += input.Quantity
		database.DB.Save(&existingItem)
	} else {
		// If not, add new item
		cartItem := models.CartItem{
			CartID:   cart.ID,
			ItemID:   input.ItemID,
			Quantity: input.Quantity,
		}
		database.DB.Create(&cartItem)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})
}

func GetCartItems(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDVal.(uint)

	var cart models.Cart
	if err := database.DB.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"cart": []models.CartItem{}})
		return
	}

	var items []models.CartItem
	database.DB.Preload("Item").Where("cart_id = ?", cart.ID).Find(&items)

	c.JSON(http.StatusOK, items)
}
