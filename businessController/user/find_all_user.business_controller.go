package user

import (
	"doce-panda/businessController/user/dtos"
	"doce-panda/domain/user/repository"
)

type FindAllUserBusinessController struct {
	userRepository repository.UserRepositoryInterface
}

func NewFindAllUserBusinessController(userRepository repository.UserRepositoryInterface) *FindAllUserBusinessController {
	return &FindAllUserBusinessController{
		userRepository: userRepository,
	}
}

func (c FindAllUserBusinessController) Execute() (*[]dtos.OutputFindAllUserDto, error) {
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
