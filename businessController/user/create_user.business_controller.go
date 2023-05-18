package user

import (
	"doce-panda/businessController/user/dtos"
	"doce-panda/domain/user/entity"
	"doce-panda/domain/user/repository"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserBusinessController struct {
	userRepository repository.UserRepositoryInterface
}

func NewCreateUserBusinessController(userRepository repository.UserRepositoryInterface) *CreateUserBusinessController {
	return &CreateUserBusinessController{
		userRepository: userRepository,
	}
}

func (c CreateUserBusinessController) Execute(input dtos.InputCreateUserDto) (*dtos.OutputCreateUserDto, error) {
	userAlreadyExist, _ := c.userRepository.FindByEmail(input.Email)

	if userAlreadyExist != nil {
		return nil, fmt.Errorf("Usuário já existe")
	}

	password := []byte(input.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	user, err := entity.NewUser(entity.User{
		Name:           input.Name,
		Email:          input.Email,
		Gender:         input.Gender,
		Password:       string(hashedPassword),
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
