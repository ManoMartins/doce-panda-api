package user

import (
	"doce-panda/domain/user/entity"
	"doce-panda/domain/user/repository"
	"doce-panda/usecase/user/dtos"
	"time"
)

type UpdateUserUseCase struct {
	userRepository repository.UserRepositoryInterface
}

func NewUpdateUserUseCase(userRepository repository.UserRepositoryInterface) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		userRepository: userRepository,
	}
}

func (c UpdateUserUseCase) Execute(input dtos.InputUpdateUserDto) (*dtos.OutputUpdateUserDto, error) {
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
