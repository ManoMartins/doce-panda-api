package product

import (
	"doce-panda/domain/product/entity"
	"doce-panda/domain/product/repository"
	"doce-panda/usecase/product/dtos"
)

type CreateProductUseCase struct {
	productRepository repository.ProductRepositoryInterface
}

func NewCreateProductUseCase(productRepository repository.ProductRepositoryInterface) *CreateProductUseCase {
	return &CreateProductUseCase{
		productRepository: productRepository,
	}
}

func (c CreateProductUseCase) Execute(input dtos.InputCreateProductDto) (*dtos.OutputCreateProductDto, error) {
	product, err := entity.NewProduct(entity.Product{
		Name:         input.Name,
		PriceInCents: input.PriceInCents,
		Description:  input.Description,
		Flavor:       input.Flavor,
		Quantity:     input.Quantity,
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
