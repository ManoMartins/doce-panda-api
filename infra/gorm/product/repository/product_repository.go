package repository

import (
	"doce-panda/domain/product/entity"
	"doce-panda/infra/gorm/product/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type ProductRepositoryDb struct {
	Db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepositoryDb {
	return &ProductRepositoryDb{Db: db}
}

func (r ProductRepositoryDb) Find(ID string) (*model.Product, error) {
	var product model.Product

	r.Db.First(&product, "id = ?", ID)

	if product.ID == "" {
		return nil, fmt.Errorf("O produto não foi encontrado")
	}

	return &product, nil
}

func (r ProductRepositoryDb) FindAll() (*[]model.Product, error) {
	var products []model.Product

	err := r.Db.Find(&products).Error

	if err != nil {
		return nil, err
	}

	return &products, nil
}

func (r ProductRepositoryDb) Create(product entity.Product) (*model.Product, error) {
	productModel := model.Product{
		ID:           product.ID,
		Name:         product.Name,
		PriceInCents: product.PriceInCents,
		Status:       product.Status,
		Description:  product.Description,
		Flavor:       product.Flavor,
		Quantity:     product.Quantity,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := r.Db.Create(&productModel).Error

	if err != nil {
		return nil, err
	}

	return &productModel, nil
}

func (r ProductRepositoryDb) Update(product entity.Product) (*model.Product, error) {
	productModel, err := r.Find(product.ID)

	if err != nil {
		return nil, err
	}

	pm := model.Product{
		ID:           product.ID,
		Name:         product.Name,
		PriceInCents: product.PriceInCents,
		Status:       product.Status,
		Description:  product.Description,
		Flavor:       product.Flavor,
		Quantity:     product.Quantity,
		CreatedAt:    productModel.CreatedAt,
		UpdatedAt:    time.Now(),
	}

	err = r.Db.Save(&pm).Error

	if err != nil {
		return nil, err
	}

	return &pm, nil
}

func (r ProductRepositoryDb) Delete(ID string) error {
	product := entity.Product{ID: ID}

	err := r.Db.Delete(&product).Error

	if err != nil {
		return err
	}

	return nil
}

func (r ProductRepositoryDb) Upload(ID string, fileUrl string) (*model.Product, error) {
	product := entity.Product{ID: ID, ImageUrl: fileUrl}

	err := r.Db.Save(&product).Error

	if err != nil {
		return nil, err
	}

	return &model.Product{}, nil
}

func (r ProductRepositoryDb) Disable(ID string) (*model.Product, error) {
	product := entity.Product{ID: ID, Status: entity.DISABLED}

	err := r.Db.Save(&product).Error

	if err != nil {
		return nil, err
	}

	return &model.Product{}, nil
}

func (r ProductRepositoryDb) Enable(ID string) (*model.Product, error) {
	product := entity.Product{ID: ID, Status: entity.ENABLED}

	err := r.Db.Save(&product).Error

	if err != nil {
		return nil, err
	}

	return &model.Product{}, nil
}
