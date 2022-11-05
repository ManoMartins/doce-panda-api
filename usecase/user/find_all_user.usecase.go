package user

import (
	"doce-panda/domain/user/repository"
	"doce-panda/usecase/user/dtos"
)

type FindAllUserUseCase struct {
	userRepository repository.UserRepositoryInterface
}

func NewFindAllUserUseCase(userRepository repository.UserRepositoryInterface) *FindAllUserUseCase {
	return &FindAllUserUseCase{
		userRepository: userRepository,
	}
}

func (c FindAllUserUseCase) Execute() (*[]dtos.OutputFindAllUserDto, error) {
	users, err := c.userRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var output []dtos.OutputFindAllUserDto

	for _, user := range *users {
		output = append(output, dtos.OutputFindAllUserDto{
			ID:             user.ID,
			Name:           user.Name,
			Gender:         user.Gender,
			PhoneNumber:    user.PhoneNumber,
			DocumentNumber: user.DocumentNumber,
			RewardPoints:   user.RewardPoints,
			Email:          user.Email,
			CreatedAt:      user.CreatedAt,
			UpdatedAt:      user.UpdatedAt,
		})
	}

	return &output, nil
}
