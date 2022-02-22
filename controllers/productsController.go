package controllers

import (
	"net/http"
	"test-go/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateProductInput struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

type UpdateProductInput struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

func FindProducts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var products []models.Product
	db.Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}

func CreateProduct(c *gin.Context) {
	var input CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	product := models.Product{Name: input.Name, Price: input.Price, Stock: input.Stock}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&product)
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func FindDetailProduct(c *gin.Context) {
	var product models.Product

	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	if err := db.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func UpdateProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var product models.Product
	id := c.Param("id")

	if err := db.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var detailProduct models.Product
	detailProduct.Name = input.Name
	detailProduct.Price = input.Price
	detailProduct.Stock = input.Stock
	db.Model(&product).Updates(detailProduct)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func DeleteProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var product models.Product
	if err := db.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"data": true, "message": "Product Deleted"})
}
