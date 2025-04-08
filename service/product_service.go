package service

import (
	"ecomventory/model"
	"ecomventory/repository"
)

// ProductService struct
type ProductService struct {
	Repo *repository.ProductRepository
}

// NewProductService creates a new ProductService
func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{Repo: repo}
}

// CreateProduct creates a new product
func (s *ProductService) CreateProduct(product *model.Product) (*model.Product, error) {
	createdProduct, err := s.Repo.CreateProduct(*product)
	if err != nil {
		return nil, err
	}
	return &createdProduct, nil
}

// GetProductByID fetches a product by its ID
func (s *ProductService) GetProductByID(id uint) (*model.Product, error) {
	product, err := s.Repo.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// UpdateProduct updates an existing product
func (s *ProductService) UpdateProduct(id uint, product *model.Product) (*model.Product, error) {
	updatedProduct, err := s.Repo.UpdateProduct(id, *product)
	if err != nil {
		return nil, err
	}
	return &updatedProduct, nil
}

// DeleteProduct deletes a product by its ID
func (s *ProductService) DeleteProduct(id uint) error {
	err := s.Repo.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}

// ListProducts lists all products
func (s *ProductService) ListProducts() ([]model.Product, error) {
	products, err := s.Repo.ListProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}
