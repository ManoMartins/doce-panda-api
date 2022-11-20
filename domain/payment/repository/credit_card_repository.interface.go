package repository

import (
	"doce-panda/domain/payment/entity"
)

type CreditCardRepositoryInterface interface {
	FindById(ID string) (*entity.CreditCard, error)
	FindAll() (*[]entity.CreditCard, error)
	Create(order entity.CreditCard) error
	Delete(ID string) error
}
