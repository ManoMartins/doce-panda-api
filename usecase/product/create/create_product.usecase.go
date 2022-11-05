package create

import (
	"doce-panda/domain/product/entity"
	"doce-panda/domain/product/repository"
)

type CreateProductUseCase struct {
	productRepository repository.ProductRepositoryInterface
}

func NewCreateProductUseCase(productRepository repository.ProductRepositoryInterface) *CreateProductUseCase {
	return &CreateProductUseCase{
		productRepository: productRepository,
	}
}

func (c CreateProductUseCase) Execute(input InputCreateProductDto) (*OutputCreateProductDto, error) {
	product, err := entity.NewProduct(entity.ProductProps{
		Name:         input.Name,
		PriceInCents: input.PriceInCents,
		Description:  input.Description,
		Flavor:       input.Flavor,
		Quantity:     input.Quantity,
	})

	if err != nil {
		return nil, err
	}

	productCreated, err := c.productRepository.Create(*product)

	if err != nil {
		return nil, err
	}

	return &OutputCreateProductDto{
		ID:           productCreated.ID,
		Name:         productCreated.Name,
		PriceInCents: productCreated.PriceInCents,
		Description:  productCreated.Description,
		Flavor:       productCreated.Flavor,
		Quantity:     productCreated.Quantity,
		CreatedAt:    productCreated.CreatedAt,
		UpdatedAt:    productCreated.UpdatedAt,
	}, nil
}
