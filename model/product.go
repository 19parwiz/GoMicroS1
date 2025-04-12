package model

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id,omitempty"` // <- OMIT if not set
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	CategoryID  uint      `json:"category_id"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
