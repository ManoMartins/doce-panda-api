package product

import (
	"doce-panda/businessController/product/dtos"
	"doce-panda/domain/product/repository"
)

type DeleteProductBusinessController struct {
	ProductRepository repository.ProductRepositoryInterface
}

func NewDeleteProductBusinessController(productRepository repository.ProductRepositoryInterface) *DeleteProductBusinessController {
	return &DeleteProductBusinessController{ProductRepository: productRepository}
}

func (c DeleteProductBusinessController) Execute(input dtos.InputDeleteProductDto) error {
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
