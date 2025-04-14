package model

import "time"

type Order struct {
	ID         int         `json:"id" gorm:"primaryKey"`
	UserID     int         `json:"user_id"`
	Status     string      `json:"status"`
	Items      []OrderItem `json:"items" gorm:"foreignKey:OrderID"` // Must be "Items" (not "Products")
	TotalPrice float64     `json:"total_price"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

type OrderItem struct {
	ID        int     `json:"id" gorm:"primaryKey"`
	OrderID   int     `json:"order_id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
}
