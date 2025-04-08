package router

import (
	"ecomventory/handler"
	"ecomventory/repository"
	"ecomventory/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Initialize repository and service
	productRepo := repository.NewProductRepository(db) // Now passing db to the repository
	productService := service.NewProductService(productRepo)

	// Initialize handler
	inventoryHandler := handler.NewInventoryHandler(productService)

	// Routes
	router.POST("/products", inventoryHandler.CreateProduct)
	router.GET("/products/:id", inventoryHandler.GetProduct)
	router.PATCH("/products/:id", inventoryHandler.UpdateProduct)
	router.DELETE("/products/:id", inventoryHandler.DeleteProduct)
	router.GET("/products", inventoryHandler.ListProducts)

	return router
}
