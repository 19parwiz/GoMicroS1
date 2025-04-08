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

	// Product wiring
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	inventoryHandler := handler.NewInventoryHandler(productService)

	router.POST("/products", inventoryHandler.CreateProduct)
	router.GET("/products/:id", inventoryHandler.GetProduct)
	router.PATCH("/products/:id", inventoryHandler.UpdateProduct)
	router.DELETE("/products/:id", inventoryHandler.DeleteProduct)
	router.GET("/products", inventoryHandler.ListProducts)

	// Category wiring — Moved inside the function
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	router.POST("/categories", categoryHandler.CreateCategory)
	router.GET("/categories/:id", categoryHandler.GetCategoryByID)
	router.PATCH("/categories/:id", categoryHandler.UpdateCategory)
	router.DELETE("/categories/:id", categoryHandler.DeleteCategory)
	router.GET("/categories", categoryHandler.GetAllCategories)

	return router
}
