package repository

import (
	"ecomventory/model"
	"fmt"
	"gorm.io/gorm"
)

// OrderRepository handles the CRUD operations for orders in the database
type OrderRepository struct {
	DB *gorm.DB
}

// NewOrderRepository creates a new order repository with a GORM DB instance
func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

// CreateOrder saves a new order to the database
func (r *OrderRepository) CreateOrder(order *model.Order) error {
	err := r.DB.Create(order).Error
	if err != nil {
		return fmt.Errorf("could not insert order: %w", err)
	}
	return nil
}

// GetOrderByID retrieves an order by its ID
func (r *OrderRepository) GetOrderByID(id int) (*model.Order, error) {
	var order model.Order
	err := r.DB.First(&order, id).Error
	if err != nil {
		return nil, fmt.Errorf("could not get order by id: %w", err)
	}
	return &order, nil
}

// UpdateOrder updates an existing order in the database
func (r *OrderRepository) UpdateOrder(order *model.Order) error {
	err := r.DB.Save(order).Error
	if err != nil {
		return fmt.Errorf("could not update order: %w", err)
	}
	return nil
}

// ListOrdersByUser retrieves all orders for a given user
func (r *OrderRepository) ListOrdersByUser(userID int) ([]model.Order, error) {
	var orders []model.Order
	err := r.DB.Where("user_id = ?", userID).Find(&orders).Error
	if err != nil {
		return nil, fmt.Errorf("could not list orders by user: %w", err)
	}
	return orders, nil
}
