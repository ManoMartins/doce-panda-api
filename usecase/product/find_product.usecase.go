package product

import (
	"doce-panda/domain/product/repository"
	"doce-panda/usecase/product/dtos"
)

type FindProductUseCase struct {
	productRepository repository.ProductRepositoryInterface
}

func NewFindProductUseCase(productRepository repository.ProductRepositoryInterface) *FindProductUseCase {
	return &FindProductUseCase{
		productRepository: productRepository,
	}
}

func (c FindProductUseCase) Execute(input dtos.InputFindProductDto) (*dtos.OutputFindProductDto, error) {
	product, err := c.productRepository.FindById(input.ID)

	if err != nil {
		return nil, err
	}

	return &dtos.OutputFindProductDto{
		ID:           product.ID,
		Name:         product.Name,
		PriceInCents: product.PriceInCents,
		Status:       product.Status,
		Description:  product.Description,
		Flavor:       product.Flavor,
		Quantity:     product.Quantity,
		ImageUrl:     "http://localhost:3333" + product.ImageUrl,
		CreatedAt:    product.CreatedAt,
		UpdatedAt:    product.UpdatedAt,
	}, nil
}
