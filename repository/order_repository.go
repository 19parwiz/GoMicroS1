package repository

import (
	"ecomventory/model"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// CreateOrder creates an order and its items in a transaction
func (r *OrderRepository) CreateOrder(order *model.Order) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 1. Save the order first to generate ID
		if err := tx.Create(order).Error; err != nil {
			return err
		}

		// 2. Save all order items with the OrderID
		for i := range order.Items {
			order.Items[i].ID = 0
			order.Items[i].OrderID = order.ID
			if err := tx.Create(&order.Items[i]).Error; err != nil {
				return err
			}
		}

		// 3. Calculate and update total price (optional)
		var total float64
		for _, item := range order.Items {
			total += item.UnitPrice * float64(item.Quantity)
		}
		return tx.Model(order).Update("total_price", total).Error
	})
}

// GetOrderByID fetches an order with its items
func (r *OrderRepository) GetOrderByID(id int) (*model.Order, error) {
	var order model.Order
	err := r.db.
		Preload("Items"). // Eager load items
		First(&order, id).Error
	return &order, err
}

// UpdateOrderStatus updates only the status field of the order
func (r *OrderRepository) UpdateOrderStatus(id int, status string) error {
	return r.db.
		Model(&model.Order{}).
		Where("id = ?", id).
		Update("status", status).Error
}

// ListOrdersByUser gets all orders for a user with their items
func (r *OrderRepository) ListOrdersByUser(userID int) ([]model.Order, error) {
	var orders []model.Order
	err := r.db.
		Preload("Items").             // Eager load items
		Where("user_id = ?", userID). // Search for orders by user ID
		Find(&orders).Error
	return orders, err
}
