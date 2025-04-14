package handler

import (
	"ecomventory/model"
	"ecomventory/service"
	"log"
	"net/http"
	"strconv" // Import this for conversion

	"github.com/gin-gonic/gin"
)

// OrderHandler handles HTTP requests related to orders
type OrderHandler struct {
	OrderService *service.OrderService
}

// NewOrderHandler creates a new instance of OrderHandler
func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{OrderService: orderService}
}

// CreateOrder handles POST requests to create a new order
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("%+v", order)
	err := h.OrderService.CreateOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, order)
}

// GetOrder handles GET requests to retrieve an order by ID
func (h *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id")
	// Convert the id from string to int
	orderID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	order, err := h.OrderService.GetOrderByID(orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	id := c.Param("id")
	orderID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var update struct { // Only accept status
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.OrderService.UpdateOrder(&model.Order{ID: orderID, Status: update.Status})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

// ListOrdersByUser handles GET requests to retrieve orders for a user by userID
func (h *OrderHandler) ListOrdersByUser(c *gin.Context) {
	userID := c.Param("user_id")
	// Convert the user_id from string to int
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID format"})
		return
	}

	orders, err := h.OrderService.ListOrdersByUser(userIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}
