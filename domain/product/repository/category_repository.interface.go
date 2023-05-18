package repository

import (
	"doce-panda/domain/product/entity"
)

type CategoryRepositoryInterface interface {
	Delete(ID string) error
	FindAll() (*[]entity.Category, error)
	Update(category entity.Category) error
	Create(category entity.Category) error
}
