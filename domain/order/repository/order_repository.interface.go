package repository

import (
	"doce-panda/domain/order/entity"
)

type OrderRepositoryInterface interface {
	FindById(ID string) (*entity.Order, error)
	FindAll() (*[]entity.Order, error)
	FindAllByUserId(UserID string) (*[]entity.Order, error)
	Create(order entity.Order) error
	UpdateStatus(order entity.Order) error
}
