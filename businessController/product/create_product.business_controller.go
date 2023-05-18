package product

import (
	"doce-panda/businessController/product/dtos"
	"doce-panda/domain/product/entity"
	"doce-panda/domain/product/repository"
)

type CreateProductBusinessController struct {
	productRepository repository.ProductRepositoryInterface
}

func NewCreateProductBusinessController(productRepository repository.ProductRepositoryInterface) *CreateProductBusinessController {
	return &CreateProductBusinessController{
		productRepository: productRepository,
	}
}

func (c CreateProductBusinessController) Execute(input dtos.InputCreateProductDto) (*dtos.OutputCreateProductDto, error) {
	product, err := entity.NewProduct(entity.Product{
		Name:         input.Name,
		PriceInCents: input.PriceInCents,
		Description:  input.Description,
		Flavor:       input.Flavor,
		Quantity:     input.Quantity,
		CategoryID:   input.CategoryID,
	})

	if err != nil {
		return nil, err
	}

	err = c.productRepository.Create(*product)

	if err != nil {
		return nil, err
	}

	return &dtos.OutputCreateProductDto{
		ID:           product.ID,
		Name:         product.Name,
		PriceInCents: product.PriceInCents,
		Description:  product.Description,
		Flavor:       product.Flavor,
		Quantity:     product.Quantity,
		CreatedAt:    product.CreatedAt,
		UpdatedAt:    product.UpdatedAt,
	}, nil
}
