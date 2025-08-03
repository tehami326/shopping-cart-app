package controllers

import (
	"net/http"
	"time"

	"shopping-cart-app/database"
	"shopping-cart-app/models"

	"github.com/gin-gonic/gin"
)

func PlaceOrder(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDVal.(uint)

	var cart models.Cart
	if err := database.DB.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No cart found"})
		return
	}

	var cartItems []models.CartItem
	database.DB.Where("cart_id = ?", cart.ID).Find(&cartItems)

	if len(cartItems) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart is empty"})
		return
	}

	var totalPrice float64
	for _, item := range cartItems {
		var dbItem models.Item
		database.DB.First(&dbItem, item.ItemID)
		totalPrice += dbItem.Price * float64(item.Quantity)
	}

	order := models.Order{
		UserID:     userID,
		TotalPrice: totalPrice,
		CreatedAt:  time.Now(),
	}
	database.DB.Create(&order)

	for _, cartItem := range cartItems {
		orderItem := models.OrderItem{
			OrderID:  order.ID,
			ItemID:   cartItem.ItemID,
			Quantity: cartItem.Quantity,
		}
		database.DB.Create(&orderItem)
	}

	database.DB.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{})

	c.JSON(http.StatusOK, gin.H{"message": "Order placed successfully"})
}

func GetOrders(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDVal.(uint)

	var orders []models.Order
	if err := database.DB.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	var result []gin.H
	for _, order := range orders {
		var orderItems []models.OrderItem
		database.DB.Where("order_id = ?", order.ID).Find(&orderItems)

		var items []gin.H
		for _, oi := range orderItems {
			var item models.Item
			database.DB.First(&item, oi.ItemID)

			items = append(items, gin.H{
				"name":     item.Name,
				"price":    item.Price,
				"quantity": oi.Quantity,
			})
		}

		result = append(result, gin.H{
			"order_id":    order.ID,
			"total_price": order.TotalPrice,
			"created_at":  order.CreatedAt,
			"items":       items,
		})
	}

	c.JSON(http.StatusOK, gin.H{"orders": result})
}
