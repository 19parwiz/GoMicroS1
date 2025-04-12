package repository

import (
	"ecomventory/model"

	"gorm.io/gorm"
)

// ProductRepository struct
type ProductRepository struct {
	DB *gorm.DB
}

// NewProductRepository creates a new ProductRepository
func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

// CreateProduct creates a new product in the database
func (r *ProductRepository) CreateProduct(product model.Product) (model.Product, error) {
	if err := r.DB.Create(&product).Error; err != nil {
		return model.Product{}, err
	}
	return product, nil
}

// GetProductByID fetches a product by its ID
func (r *ProductRepository) GetProductByID(id uint) (model.Product, error) {
	var product model.Product
	if err := r.DB.First(&product, id).Error; err != nil {
		return model.Product{}, err
	}
	return product, nil
}

// UpdateProduct updates an existing product
func (r *ProductRepository) UpdateProduct(id uint, product model.Product) (model.Product, error) {
	var existingProduct model.Product
	if err := r.DB.First(&existingProduct, id).Error; err != nil {
		return model.Product{}, err
	}

	// Update product fields here
	if err := r.DB.Model(&existingProduct).Updates(product).Error; err != nil {
		return model.Product{}, err
	}

	return existingProduct, nil
}

// DeleteProduct deletes a product by its ID
func (r *ProductRepository) DeleteProduct(id uint) error {
	if err := r.DB.Delete(&model.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}

// ListProducts lists all products
func (r *ProductRepository) ListProducts() ([]model.Product, error) {
	var products []model.Product
	if err := r.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) FilterProducts(category string, minPrice, maxPrice float64, page, limit int) ([]model.Product, error) {
	var products []model.Product
	query := r.DB

	if category != "" {
		query = query.Where("category_id = ?", category)
	}
	if minPrice > 0 {
		query = query.Where("price >= ?", minPrice)
	}
	if maxPrice > 0 {
		query = query.Where("price <= ?", maxPrice)
	}

	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Find(&products).Error
	return products, err
}
