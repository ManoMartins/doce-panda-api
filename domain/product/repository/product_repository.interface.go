package repository

import (
	"doce-panda/domain/product/entity"
	"doce-panda/infra/gorm/product/model"
)

type ProductRepositoryInterface interface {
	Find(ID string) (*model.Product, error)
	Delete(ID string) error
	FindAll() (*[]model.Product, error)
	Update(product entity.Product) (*model.Product, error)
	Create(product entity.Product) (*model.Product, error)
	Upload(ID string, fileUrl string) (*model.Product, error)
	Disable(ID string) (*model.Product, error)
	Enable(ID string) (*model.Product, error)
}
