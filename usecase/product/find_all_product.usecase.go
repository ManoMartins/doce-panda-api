package product

import (
	"doce-panda/domain/product/repository"
	"doce-panda/usecase/product/dtos"
)

type FindAllProductUseCase struct {
	productRepository repository.ProductRepositoryInterface
}

func NewFindAllProductUseCase(productRepository repository.ProductRepositoryInterface) *FindAllProductUseCase {
	return &FindAllProductUseCase{
		productRepository: productRepository,
	}
}

func (c FindAllProductUseCase) Execute() (*[]dtos.OutputFindAllProductDto, error) {
	products, err := c.productRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var output []dtos.OutputFindAllProductDto

	for _, product := range *products {
		output = append(output, dtos.OutputFindAllProductDto{
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
		})
	}

	return &output, nil
}
