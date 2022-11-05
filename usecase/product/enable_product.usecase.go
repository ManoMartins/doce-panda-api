package product

import (
	"doce-panda/domain/product/repository"
	"doce-panda/usecase/product/dtos"
)

type EnableProductUseCase struct {
	ProductRepository repository.ProductRepositoryInterface
}

func NewEnableProductUseCase(productRepository repository.ProductRepositoryInterface) *EnableProductUseCase {
	return &EnableProductUseCase{ProductRepository: productRepository}
}

func (c EnableProductUseCase) Execute(input dtos.InputEnableProductDto) error {
	product, err := c.ProductRepository.FindById(input.ID)

	err = product.Enable()

	if err != nil {
		return err
	}

	err = c.ProductRepository.Update(*product)

	if err != nil {
		return err
	}

	return nil
}
