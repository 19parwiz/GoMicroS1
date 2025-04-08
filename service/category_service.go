package service

import (
	"ecomventory/model"
	"ecomventory/repository"
)

type CategoryService interface {
	CreateCategory(category *model.Category) error
	GetCategoryByID(id uint) (*model.Category, error)
	UpdateCategory(category *model.Category) error
	DeleteCategory(id uint) error
	GetAllCategories() ([]model.Category, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo}
}

func (s *categoryService) CreateCategory(category *model.Category) error {
	return s.repo.Create(category)
}

func (s *categoryService) GetCategoryByID(id uint) (*model.Category, error) {
	return s.repo.GetByID(id)
}

func (s *categoryService) UpdateCategory(category *model.Category) error {
	return s.repo.Update(category)
}

func (s *categoryService) DeleteCategory(id uint) error {
	return s.repo.Delete(id)
}

func (s *categoryService) GetAllCategories() ([]model.Category, error) {
	return s.repo.GetAll()
}
