package model

import "gorm.io/gorm"

type Category struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null;unique"`
	Description string    `json:"description"`
	Products    []Product `json:"products,omitempty" gorm:"foreignKey:CategoryID"`
}

func MigrateCategory(db *gorm.DB) error {
	return db.AutoMigrate(&Category{})
}
