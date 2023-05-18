package user

import (
	"doce-panda/businessController/user/dtos"
	"doce-panda/domain/user/repository"
)

type FindUserBusinessController struct {
	userRepository repository.UserRepositoryInterface
}

func NewFindUserBusinessController(userRepository repository.UserRepositoryInterface) *FindUserBusinessController {
	return &FindUserBusinessController{userRepository: userRepository}
}

func (c FindUserBusinessController) Execute(input dtos.InputFindUserDto) (*dtos.OutputFindUserDto, error) {
	user, err := c.userRepository.FindById(input.ID)

	if err != nil {
		return nil, err
	}

	return &dtos.OutputFindUserDto{
		ID:             user.ID,
		Name:           user.Name,
		Gender:         user.Gender,
		PhoneNumber:    user.PhoneNumber,
		DocumentNumber: user.DocumentNumber,
		RewardPoints:   user.RewardPoints,
		Email:          user.Email,
		Addresses:      user.Addresses,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}, nil
}
