package routes

import (
	"shopping-cart-app/controllers"
	"shopping-cart-app/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(r *gin.Engine) {
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/orders", controllers.PlaceOrder)
		auth.GET("/orders/:user_id", controllers.GetOrders)
	}
}
