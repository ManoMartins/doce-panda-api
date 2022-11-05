package user

import (
	"doce-panda/domain/user/repository"
	"doce-panda/usecase/user/dtos"
)

type DestroyUserUseCase struct {
	UserRepository repository.UserRepositoryInterface
}

func NewDestroyUserUseCase(productRepository repository.UserRepositoryInterface) *DestroyUserUseCase {
	return &DestroyUserUseCase{UserRepository: productRepository}
}

func (c DestroyUserUseCase) Execute(input dtos.InputDestroyUserDto) error {
	_, err := c.UserRepository.FindById(input.ID)

	if err != nil {
		return err
	}

	err = c.UserRepository.Delete(input.ID)

	if err != nil {
		return err
	}

	return nil
}
