package routes

import (
	"shopping-cart-app/controllers"
	"shopping-cart-app/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterCartRoutes(r *gin.Engine) {
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/carts", controllers.AddToCart)
		auth.GET("/carts/:user_id", controllers.GetCartItems)
	}
}
