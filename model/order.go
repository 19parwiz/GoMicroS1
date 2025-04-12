package model

import "time"

// Order represents an order in the system
type Order struct {
	ID         int       `json:"id" db:"id"`
	UserID     int       `json:"user_id" db:"user_id"`
	Status     string    `json:"status" db:"status"` // e.g., "pending", "completed", "cancelled"
	Products   []Product `json:"products"`           // Product details associated with the order
	TotalPrice float64   `json:"total_price" db:"total_price"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
