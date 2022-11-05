package enable

import (
	"doce-panda/domain/product/entity"
	"doce-panda/domain/product/repository"
)

type EnableProductUseCase struct {
	ProductRepository repository.ProductRepositoryInterface
}

func NewEnableProductUseCase(productRepository repository.ProductRepositoryInterface) *EnableProductUseCase {
	return &EnableProductUseCase{ProductRepository: productRepository}
}

func (c EnableProductUseCase) Execute(input InputEnableProductDto) (*entity.Product, error) {
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

	err = product.Enable()

	if err != nil {
		return nil, err
	}

	_, err = c.ProductRepository.Update(product)

	return &product, nil
}
