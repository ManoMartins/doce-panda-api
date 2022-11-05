package user

import (
	"doce-panda/domain/user/repository"
	"doce-panda/usecase/user/dtos"
)

type FindUserUseCase struct {
	userRepository repository.UserRepositoryInterface
}

func NewFindUserUseCase(userRepository repository.UserRepositoryInterface) *FindUserUseCase {
	return &FindUserUseCase{userRepository: userRepository}
}

func (c FindUserUseCase) Execute(input dtos.InputFindUserDto) (*dtos.OutputFindUserDto, error) {
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
