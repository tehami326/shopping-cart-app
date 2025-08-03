package routes

import (
	"shopping-cart-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	r.POST("/users", controllers.CreateUser)
	r.POST("/login", controllers.LoginUser)
	r.GET("/users", controllers.GetUsers)
}
