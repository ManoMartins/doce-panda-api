package repository

import (
	"doce-panda/domain/order/entity"
)

type OrderRepositoryInterface interface {
	FindById(ID string) (*entity.Order, error)
	FindAll() (*[]entity.Order, error)
	Create(order entity.Order) error
}
