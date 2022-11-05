package user

import (
	"doce-panda/domain/user/entity"
	"doce-panda/domain/user/repository"
	"doce-panda/usecase/user/dtos"
	"fmt"
)

type CreateUserUseCase struct {
	userRepository repository.UserRepositoryInterface
}

func NewCreateUserUseCase(userRepository repository.UserRepositoryInterface) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: userRepository,
	}
}

func (c CreateUserUseCase) Execute(input dtos.InputCreateUserDto) (*dtos.OutputCreateUserDto, error) {
	userAlreadyExist, _ := c.userRepository.FindByEmail(input.Email)

	if userAlreadyExist != nil {
		return nil, fmt.Errorf("Usuário já existe")
	}

	user, err := entity.NewUser(entity.User{
		Name:           input.Name,
		Email:          input.Email,
		Gender:         input.Gender,
		Password:       input.Password,
		PhoneNumber:    input.PhoneNumber,
		DocumentNumber: input.DocumentNumber,
	})

	if err != nil {
		return nil, err
	}

	err = c.userRepository.Create(*user)

	if err != nil {
		return nil, err
	}

	return &dtos.OutputCreateUserDto{
		ID:             user.ID,
		Name:           user.Name,
		Email:          user.Email,
		Gender:         user.Gender,
		PhoneNumber:    user.PhoneNumber,
		DocumentNumber: user.DocumentNumber,
		RewardPoints:   user.RewardPoints,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}, nil
}
