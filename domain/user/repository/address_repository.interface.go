package repository

import (
	"doce-panda/domain/user/entity"
)

type AddressRepositoryInterface interface {
	FindById(ID string) (*entity.Address, error)
	FindByMain() (*entity.Address, error)
	Delete(ID string) error
	FindAllByUserId(userID string) (*[]entity.Address, error)
	Update(address entity.Address) error
	Create(address entity.Address) error
}
