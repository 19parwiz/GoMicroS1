package service

import (
	"ecomventory/model"
	"ecomventory/repository"
)

// OrderService handles the business logic for orders
type OrderService struct {
	OrderRepo *repository.OrderRepository // Accept a pointer to OrderRepository
}

// NewOrderService creates a new instance of OrderService
// It should now accept a pointer to OrderRepository
func NewOrderService(orderRepo *repository.OrderRepository) *OrderService {
	return &OrderService{OrderRepo: orderRepo}
}

// CreateOrder creates a new order
func (s *OrderService) CreateOrder(order *model.Order) error {
	return s.OrderRepo.CreateOrder(order)
}

// GetOrderByID retrieves an order by its ID
func (s *OrderService) GetOrderByID(id int) (*model.Order, error) {
	return s.OrderRepo.GetOrderByID(id)
}

// UpdateOrder updates an existing order
func (s *OrderService) UpdateOrder(order *model.Order) error {
	return s.OrderRepo.UpdateOrder(order)
}

// ListOrdersByUser retrieves all orders for a given user
func (s *OrderService) ListOrdersByUser(userID int) ([]model.Order, error) {
	return s.OrderRepo.ListOrdersByUser(userID)
}
