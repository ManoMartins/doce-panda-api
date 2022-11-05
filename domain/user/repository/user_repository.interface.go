package repository

import (
	"doce-panda/domain/user/entity"
)

type UserRepositoryInterface interface {
	Find(ID string) (*entity.User, error)
	Delete(ID string) error
	FindAll() (*[]entity.User, error)
	Update(user entity.User) error
	Create(user entity.User) error
}
