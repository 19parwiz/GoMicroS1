package handler

import (
	"ecomventory/model"
	"ecomventory/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// InventoryHandler struct
type InventoryHandler struct {
	Service *service.ProductService
}

// NewInventoryHandler creates a new InventoryHandler
func NewInventoryHandler(service *service.ProductService) *InventoryHandler {
	return &InventoryHandler{Service: service}
}

// CreateProduct handles product creation
func (h *InventoryHandler) CreateProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := h.Service.CreateProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}

// GetProduct handles getting a product by ID
func (h *InventoryHandler) GetProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// Convert id to uint here
	product, err := h.Service.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// UpdateProduct handles updating a product
func (h *InventoryHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProduct, err := h.Service.UpdateProduct(uint(id), &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

// DeleteProduct handles deleting a product
func (h *InventoryHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	err = h.Service.DeleteProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

func (h *InventoryHandler) ListProducts(c *gin.Context) {
	category := c.Query("category")
	minPrice, _ := strconv.ParseFloat(c.DefaultQuery("minPrice", "0"), 64)
	maxPrice, _ := strconv.ParseFloat(c.DefaultQuery("maxPrice", "0"), 64)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	products, err := h.Service.FilterProducts(category, minPrice, maxPrice, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, products)
}
