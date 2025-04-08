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

	// Pass the pointer of product to the service
	createdProduct, err := h.Service.CreateProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}

	c.JSON(http.StatusOK, product)


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

	updatedProduct, err := h.Service.UpdateProduct(uint(id), product)
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

// ListProducts handles listing all products
func (h *InventoryHandler) ListProducts(c *gin.Context) {
	products, err := h.Service.ListProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
