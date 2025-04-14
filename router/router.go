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
	productHandler := handler.NewProductHandler(productService)

	router.POST("/products", productHandler.CreateProduct)
	router.GET("/products/:id", productHandler.GetProductByID)
	router.PATCH("/products/:id", productHandler.UpdateProduct)
	router.DELETE("/products/:id", productHandler.DeleteProduct)
	router.GET("/products", productHandler.ListProducts)

	// Category wiring
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	router.POST("/categories", categoryHandler.CreateCategory)
	router.GET("/categories/:id", categoryHandler.GetCategoryByID)
	router.PATCH("/categories/:id", categoryHandler.UpdateCategory)
	router.DELETE("/categories/:id", categoryHandler.DeleteCategory)
	router.GET("/categories", categoryHandler.GetAllCategories)

	// Order wiring
	orderRepo := repository.NewOrderRepository(db)     // Corrected: orderRepo is passed here
	orderService := service.NewOrderService(orderRepo) // Corrected: Pass orderRepo to NewOrderService
	orderHandler := handler.NewOrderHandler(orderService)

	router.POST("/orders", orderHandler.CreateOrder)
	router.GET("/orders/:id", orderHandler.GetOrder)
	router.PATCH("/orders/:id", orderHandler.UpdateOrderStatus)
	router.GET("/orders/user/:user_id", orderHandler.ListOrdersByUser)

	return router
}
