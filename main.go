package main

import (
	"github.com/gin-gonic/gin"
	"modaraljazzaa/go-product-api/config"
	"modaraljazzaa/go-product-api/routes"
)

func main() {
	// Initialize database
	config.ConnectDatabase()

	// Setup Gin router
	r := gin.Default()

	// Register routes
	routes.RegisterProductRoutes(r)

	// Start the server
	r.Run(":8080")
}
