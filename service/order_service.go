package service

import (
	"ecomventory/model"
	"ecomventory/repository"
)

type OrderService struct {
	OrderRepo *repository.OrderRepository
}

func NewOrderService(orderRepo *repository.OrderRepository) *OrderService {
	return &OrderService{OrderRepo: orderRepo}
}

func (s *OrderService) CreateOrder(order *model.Order) error {
	return s.OrderRepo.CreateOrder(order)
}

func (s *OrderService) GetOrderByID(id int) (*model.Order, error) {
	return s.OrderRepo.GetOrderByID(id)
}

func (s *OrderService) UpdateOrder(order *model.Order) error {
	return s.OrderRepo.UpdateOrderStatus(order.ID, order.Status)
}

func (s *OrderService) ListOrdersByUser(userID int) ([]model.Order, error) {
	return s.OrderRepo.ListOrdersByUser(userID)
}
