package product

import (
	"doce-panda/businessController/product/dtos"
	"doce-panda/domain/product/repository"
)

type FindProductBusinessController struct {
	productRepository repository.ProductRepositoryInterface
}

func NewFindProductBusinessController(productRepository repository.ProductRepositoryInterface) *FindProductBusinessController {
	return &FindProductBusinessController{
		productRepository: productRepository,
	}
}

func (c FindProductBusinessController) Execute(input dtos.InputFindProductDto) (*dtos.OutputFindProductDto, error) {
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
