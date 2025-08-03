package routes

import (
	"shopping-cart-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterItemRoutes(r *gin.Engine) {
	r.POST("/items", controllers.CreateItem)
	r.GET("/items", controllers.GetItems)
}
