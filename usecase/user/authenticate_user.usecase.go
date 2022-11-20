package user

import (
	"doce-panda/domain/user/repository"
	"doce-panda/usecase/user/dtos"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthenticateUserUseCase struct {
	userRepo repository.UserRepositoryInterface
}

func NewAuthenticateUserUseCase(userRepo repository.UserRepositoryInterface) *AuthenticateUserUseCase {
	return &AuthenticateUserUseCase{userRepo: userRepo}
}

func (a AuthenticateUserUseCase) Execute(input dtos.InputAuthenticationUserDto) (*dtos.OutputAuthenticationUserDto, error) {
	user, _ := a.userRepo.FindByEmail(input.Email)

	if user == nil {
		return nil, fmt.Errorf("Email / Senha incorreto")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, fmt.Errorf("Email / Senha incorreto")
	}

	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return nil, err
	}

	return &dtos.OutputAuthenticationUserDto{Token: t, User: struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}(struct {
		ID   string
		Name string
	}{ID: user.ID, Name: user.Name})}, nil
}
