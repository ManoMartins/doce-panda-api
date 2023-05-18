package repository

import (
	"doce-panda/domain/order/entity"
	"time"
)

type InputReport struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

type OrderRepositoryInterface interface {
	FindById(ID string) (*entity.Order, error)
	FindAll() (*[]entity.Order, error)
	FindAllByUserId(UserID string) (*[]entity.Order, error)
	Create(order entity.Order) error
	UpdateStatus(order entity.Order) error
	Report(input InputReport) (*[]entity.OrderItem, error)
}
