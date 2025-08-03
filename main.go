package main

import (
	"shopping-cart-app/database"
	"shopping-cart-app/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	database.Connect()
	database.SeedItems()

	routes.RegisterUserRoutes(r)
	routes.RegisterItemRoutes(r)
	routes.RegisterCartRoutes(r)
	routes.RegisterOrderRoutes(r)

	r.Run(":8080")
}
