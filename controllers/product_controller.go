package controllers

import (
	"github.com/gin-gonic/gin"
	"modaraljazzaa/go-product-api/config"
	"modaraljazzaa/go-product-api/dto"
	"modaraljazzaa/go-product-api/models"
	"net/http"
)

// Get all products
func GetProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

// Get product by ID
func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// Create a new product
func CreateProduct(c *gin.Context) {
	var input dto.CreateProductDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := models.Product{Name: input.Name, Price: input.Price}
	config.DB.Create(&product)
	c.JSON(http.StatusCreated, product)
}

// Update a product
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var input dto.UpdateProductDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Name != "" {
		product.Name = input.Name
	}
	if input.Price > 0 {
		product.Price = input.Price
	}

	config.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

// Delete a product
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
