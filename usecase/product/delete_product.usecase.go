package product

import (
	"doce-panda/domain/product/repository"
	"doce-panda/usecase/product/dtos"
)

type DeleteProductUseCase struct {
	ProductRepository repository.ProductRepositoryInterface
}

func NewDeleteProductUseCase(productRepository repository.ProductRepositoryInterface) *DeleteProductUseCase {
	return &DeleteProductUseCase{ProductRepository: productRepository}
}

func (c DeleteProductUseCase) Execute(input dtos.InputDeleteProductDto) error {
	_, err := c.ProductRepository.FindById(input.ID)

	if err != nil {
		return err
	}

	err = c.ProductRepository.Delete(input.ID)

	if err != nil {
		return err
	}

	return nil
}
