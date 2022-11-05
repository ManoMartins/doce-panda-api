package find

import (
	"doce-panda/domain/product/repository"
	"fmt"
)

type FindProductUseCase struct {
	productRepository repository.ProductRepositoryInterface
}

func NewFindProductUseCase(productRepository repository.ProductRepositoryInterface) *FindProductUseCase {
	return &FindProductUseCase{
		productRepository: productRepository,
	}
}

func (c FindProductUseCase) Execute(input InputFindProductDto) (*OutputFindProductDto, error) {
	product, err := c.productRepository.Find(input.ID)

	if err != nil {
		return nil, fmt.Errorf("produto não encontrado")
	}

	return &OutputFindProductDto{
		ID:           product.ID,
		Name:         product.Name,
		PriceInCents: product.PriceInCents,
		Status:       StatusEnum(product.Status),
		Description:  product.Description,
		Flavor:       product.Flavor,
		Quantity:     product.Quantity,
		CreatedAt:    product.CreatedAt,
		UpdatedAt:    product.UpdatedAt,
	}, nil
}
