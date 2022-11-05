package repository

import (
	"doce-panda/domain/product/entity"
	"doce-panda/infra/gorm/product/model"
)

type ProductRepositoryInterface interface {
	FindById(ID string) (*entity.Product, error)
	Delete(ID string) error
	FindAll() (*[]entity.Product, error)
	Update(product entity.Product) error
	Create(product entity.Product) error
	Disable(ID string) (*model.Product, error)
	Enable(ID string) (*model.Product, error)
}
