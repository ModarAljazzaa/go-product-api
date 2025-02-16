package routes

import (
	"github.com/gin-gonic/gin"
	"modaraljazzaa/go-product-api/controllers"
)

func RegisterProductRoutes(r *gin.Engine) {
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/:id", controllers.GetProductByID)
	r.POST("/products", controllers.CreateProduct)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)
}
