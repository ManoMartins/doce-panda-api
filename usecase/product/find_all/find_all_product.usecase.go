package find_all

import (
	"doce-panda/domain/product/repository"
)

type FindAllProductUseCase struct {
	productRepository repository.ProductRepositoryInterface
}

func NewFindAllProductUseCase(productRepository repository.ProductRepositoryInterface) *FindAllProductUseCase {
	return &FindAllProductUseCase{
		productRepository: productRepository,
	}
}

func (c FindAllProductUseCase) Execute() (*[]OutputFindAllProductDto, error) {
	products, err := c.productRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var output []OutputFindAllProductDto

	for _, product := range *products {
		output = append(output, OutputFindAllProductDto{
			ID:           product.ID,
			Name:         product.Name,
			PriceInCents: product.PriceInCents,
			Status:       StatusEnum(product.Status),
			Description:  product.Description,
			Flavor:       product.Flavor,
			Quantity:     product.Quantity,
			CreatedAt:    product.CreatedAt,
			UpdatedAt:    product.UpdatedAt,
		})
	}

	return &output, nil
}
