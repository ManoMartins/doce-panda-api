package repository

import (
	"doce-panda/domain/product/entity"
	"doce-panda/infra/gorm/product/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

type ProductRepositoryDb struct {
	Db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepositoryDb {
	return &ProductRepositoryDb{Db: db}
}

func (r ProductRepositoryDb) FindById(ID string) (*entity.Product, error) {
	var productModel model.Product

	r.Db.First(&productModel, "id = ?", ID)

	if productModel.ID == "" {
		return nil, fmt.Errorf("O produto não foi encontrado")
	}

	return entity.NewProduct(entity.Product{
		ID:           productModel.ID,
		Name:         productModel.Name,
		PriceInCents: productModel.PriceInCents,
		Status:       productModel.Status,
		Description:  productModel.Description,
		Flavor:       productModel.Flavor,
		Quantity:     productModel.Quantity,
		ImageUrl:     productModel.ImageUrl,
		CreatedAt:    productModel.CreatedAt,
		UpdatedAt:    productModel.UpdatedAt,
	})
}

func (r ProductRepositoryDb) FindAll() (*[]entity.Product, error) {
	var productsModel []model.Product

	err := r.Db.Find(&productsModel).Error

	if err != nil {
		return nil, err
	}

	var products []entity.Product

	for _, productModel := range productsModel {
		product, err := entity.NewProduct(entity.Product{
			ID:           productModel.ID,
			Name:         productModel.Name,
			PriceInCents: productModel.PriceInCents,
			Status:       productModel.Status,
			Description:  productModel.Description,
			Flavor:       productModel.Flavor,
			Quantity:     productModel.Quantity,
			ImageUrl:     productModel.ImageUrl,
			CreatedAt:    productModel.CreatedAt,
			UpdatedAt:    productModel.UpdatedAt,
		})

		if err != nil {
			return nil, err
		}

		products = append(products, *product)
	}

	return &products, nil
}

func (r ProductRepositoryDb) Create(product entity.Product) error {
	productModel := model.Product{
		ID:           product.ID,
		Name:         product.Name,
		PriceInCents: product.PriceInCents,
		Status:       product.Status,
		Description:  product.Description,
		Flavor:       product.Flavor,
		Quantity:     product.Quantity,
		CreatedAt:    product.CreatedAt,
		UpdatedAt:    product.UpdatedAt,
	}

	err := r.Db.Create(&productModel).Error

	if err != nil {
		return err
	}

	return nil
}

func (r ProductRepositoryDb) Update(product entity.Product) error {
	productModel := model.Product{
		ID:           product.ID,
		Name:         product.Name,
		PriceInCents: product.PriceInCents,
		Status:       product.Status,
		Description:  product.Description,
		Flavor:       product.Flavor,
		Quantity:     product.Quantity,
		CreatedAt:    product.CreatedAt,
		UpdatedAt:    product.UpdatedAt,
	}

	err := r.Db.Save(&productModel).Error

	if err != nil {
		return err
	}

	return nil
}

func (r ProductRepositoryDb) Delete(ID string) error {
	product := entity.Product{ID: ID}

	err := r.Db.Delete(&product).Error

	if err != nil {
		return err
	}

	return nil
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
