package destroy

import "doce-panda/domain/product/repository"

type DeleteProductUseCase struct {
	ProductRepository repository.ProductRepositoryInterface
}

func NewDeleteProductUseCase(productRepository repository.ProductRepositoryInterface) *DeleteProductUseCase {
	return &DeleteProductUseCase{ProductRepository: productRepository}
}

func (c DeleteProductUseCase) Execute(input InputDeleteProductDto) error {
	_, err := c.ProductRepository.Find(input.ID)

	if err != nil {
		return err
	}

	err = c.ProductRepository.Delete(input.ID)

	if err != nil {
		return err
	}

	return nil
}
