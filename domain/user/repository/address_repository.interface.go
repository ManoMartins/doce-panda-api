package repository

import (
	"doce-panda/domain/user/entity"
)

type AddressRepositoryInterface interface {
	Find(ID string) (*entity.Address, error)
	FindMain() (*entity.Address, error)
	Delete(ID string) error
	FindAll() (*[]entity.Address, error)
	Update(address entity.Address) error
	Create(address entity.Address) error
}
