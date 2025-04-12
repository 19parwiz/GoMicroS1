package handler

import (
	"ecomventory/model"   // Import model for Product struct
	"ecomventory/service" // Import service for ProductService
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProductHandler struct
type ProductHandler struct {
	ProductService *service.ProductService
}

// NewProductHandler initializes and returns a new ProductHandler
func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{ProductService: productService}
}

// CreateProduct handles creating a new product
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product model.Product

	// Bind incoming JSON to product model
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call CreateProduct from ProductService instead of h.service
	createdProduct, err := h.ProductService.CreateProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdProduct)
}

// GetProductByID handles fetching a product by its ID
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := h.ProductService.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// UpdateProduct handles updating an existing product
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProduct, err := h.ProductService.UpdateProduct(uint(id), &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

// DeleteProduct handles deleting a product by ID
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	err = h.ProductService.DeleteProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

// ListProducts handles fetching all products
func (h *ProductHandler) ListProducts(c *gin.Context) {
	products, err := h.ProductService.ListProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// FilterProducts handles filtering products based on query parameters
func (h *ProductHandler) FilterProducts(c *gin.Context) {
	category := c.Query("category")
	minPriceStr := c.DefaultQuery("min_price", "0")
	maxPriceStr := c.DefaultQuery("max_price", "1000000")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	minPrice, _ := strconv.ParseFloat(minPriceStr, 64)
	maxPrice, _ := strconv.ParseFloat(maxPriceStr, 64)
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	products, err := h.ProductService.FilterProducts(category, minPrice, maxPrice, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
