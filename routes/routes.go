package routes

import (
	"test-go/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/product", controllers.FindProducts)
	r.POST("/product", controllers.CreateProduct)
	r.GET("/product/:id", controllers.FindDetailProduct)
	r.PUT("/product/:id", controllers.UpdateProduct)
	r.DELETE("/product/:id", controllers.DeleteProduct)
	return r
}
