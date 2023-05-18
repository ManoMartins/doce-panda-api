package product

import (
	"doce-panda/businessController/product/dtos"
	"doce-panda/domain/product/repository"
)

type FindAllProductBusinessController struct {
	productRepository repository.ProductRepositoryInterface
}

func NewFindAllProductBusinessController(productRepository repository.ProductRepositoryInterface) *FindAllProductBusinessController {
	return &FindAllProductBusinessController{
		productRepository: productRepository,
	}
}

func (c FindAllProductBusinessController) Execute() (*[]dtos.OutputFindAllProductDto, error) {
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
