package user

import (
	"doce-panda/businessController/user/dtos"
	"doce-panda/domain/user/entity"
	"doce-panda/domain/user/repository"
	"time"
)

type UpdateUserBusinessController struct {
	userRepository repository.UserRepositoryInterface
}

func NewUpdateUserBusinessController(userRepository repository.UserRepositoryInterface) *UpdateUserBusinessController {
	return &UpdateUserBusinessController{
		userRepository: userRepository,
	}
}

func (c UpdateUserBusinessController) Execute(input dtos.InputUpdateUserDto) (*dtos.OutputUpdateUserDto, error) {
	userFound, err := c.userRepository.FindById(input.ID)

	if err != nil {
		return nil, err
	}

	err = userFound.Update(entity.UserUpdateProps{
		Name:           input.Name,
		Gender:         input.Gender,
		PhoneNumber:    input.PhoneNumber,
		DocumentNumber: input.DocumentNumber,
	})

	userFound.UpdatedAt = time.Now()

	if err != nil {
		return nil, err
	}

	err = c.userRepository.Update(*userFound)

	if err != nil {
		return nil, err
	}

	return &dtos.OutputUpdateUserDto{
		ID:             userFound.ID,
		Name:           userFound.Name,
		Email:          userFound.Email,
		Gender:         userFound.Gender,
		PhoneNumber:    userFound.PhoneNumber,
		DocumentNumber: userFound.DocumentNumber,
		RewardPoints:   userFound.RewardPoints,
		CreatedAt:      userFound.CreatedAt,
		UpdatedAt:      userFound.UpdatedAt,
	}, nil
}
