package user

import (
	"doce-panda/businessController/user/dtos"
	"doce-panda/domain/user/repository"
)

type DestroyUserBusinessController struct {
	UserRepository repository.UserRepositoryInterface
}

func NewDestroyUserBusinessController(productRepository repository.UserRepositoryInterface) *DestroyUserBusinessController {
	return &DestroyUserBusinessController{UserRepository: productRepository}
}

func (c DestroyUserBusinessController) Execute(input dtos.InputDestroyUserDto) error {
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
