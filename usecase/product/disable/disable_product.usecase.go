package disable

import (
	"doce-panda/domain/product/entity"
	"doce-panda/domain/product/repository"
)

type DisableProductUseCase struct {
	ProductRepository repository.ProductRepositoryInterface
}

func NewDisableProductUseCase(productRepository repository.ProductRepositoryInterface) *DisableProductUseCase {
	return &DisableProductUseCase{ProductRepository: productRepository}
}

func (c DisableProductUseCase) Execute(input InputDisableProductDto) (*entity.Product, error) {
	productModel, err := c.ProductRepository.Find(input.ID)

	product := entity.Product{
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
	}

	err = product.Disable()

	if err != nil {
		return nil, err
	}

	_, err = c.ProductRepository.Update(product)

	return &product, nil
}
