package repository

import (
	"doce-panda/domain/user/entity"
	"doce-panda/infra/gorm/user/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

type UserRepositoryDb struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryDb {
	return &UserRepositoryDb{Db: db}
}

func (r UserRepositoryDb) FindById(ID string) (*entity.User, error) {
	var user model.User

	r.Db.Preload("Addresses").First(&user, "id = ?", ID)

	if user.ID == "" {
		return nil, fmt.Errorf("O usuário não foi encontrado")
	}

	var addresses []entity.Address

	for _, addressModel := range user.Addresses {
		address, err := entity.NewAddress(entity.Address{
			ID:           addressModel.ID,
			City:         addressModel.City,
			State:        addressModel.State,
			Street:       addressModel.Street,
			Number:       addressModel.Number,
			ZipCode:      addressModel.ZipCode,
			Neighborhood: addressModel.Neighborhood,
			IsMain:       addressModel.IsMain,
			UserID:       addressModel.UserID,
			CreatedAt:    addressModel.CreatedAt,
			UpdatedAt:    addressModel.UpdatedAt,
		})

		if err != nil {
			return nil, err
		}

		addresses = append(addresses, *address)
	}

	return entity.NewUser(entity.User{
		ID:             user.ID,
		Name:           user.Name,
		Gender:         user.Gender,
		Password:       user.Password,
		PhoneNumber:    user.PhoneNumber,
		DocumentNumber: user.DocumentNumber,
		RewardPoints:   user.RewardPoints,
		Email:          user.Email,
		Addresses:      addresses,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	})
}

func (r UserRepositoryDb) FindByEmail(email string) (*entity.User, error) {
	var user model.User

	r.Db.First(&user, "email = ?", email)

	if user.ID == "" {
		return nil, fmt.Errorf("O usuário não foi encontrado")
	}

	return entity.NewUser(entity.User{
		ID:             user.ID,
		Name:           user.Name,
		Gender:         user.Gender,
		Password:       user.Password,
		PhoneNumber:    user.PhoneNumber,
		DocumentNumber: user.DocumentNumber,
		RewardPoints:   user.RewardPoints,
		Email:          user.Email,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	})
}

func (r UserRepositoryDb) Delete(ID string) error {
	user := entity.User{ID: ID}

	err := r.Db.Delete(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (r UserRepositoryDb) FindAll() (*[]entity.User, error) {
	var usersModel []model.User

	err := r.Db.Find(&usersModel).Error

	if err != nil {
		return nil, err
	}

	var users []entity.User

	for _, userModel := range usersModel {
		user, err := entity.NewUser(entity.User{
			ID:             userModel.ID,
			Name:           userModel.Name,
			Gender:         userModel.Gender,
			Password:       userModel.Password,
			PhoneNumber:    userModel.PhoneNumber,
			DocumentNumber: userModel.DocumentNumber,
			RewardPoints:   userModel.RewardPoints,
			Email:          userModel.Email,
			CreatedAt:      userModel.CreatedAt,
			UpdatedAt:      userModel.UpdatedAt,
		})

		if err != nil {
			return nil, err
		}

		users = append(users, *user)
	}

	return &users, nil
}

func (r UserRepositoryDb) Update(user entity.User) error {
	userModel := model.User{
		ID:             user.ID,
		Name:           user.Name,
		Gender:         user.Gender,
		PhoneNumber:    user.PhoneNumber,
		DocumentNumber: user.DocumentNumber,
		RewardPoints:   user.RewardPoints,
		Email:          user.Email,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}

	err := r.Db.Save(&userModel).Error

	if err != nil {
		return err
	}

	return nil
}

func (r UserRepositoryDb) Create(user entity.User) error {
	userModel := model.User{
		ID:             user.ID,
		Name:           user.Name,
		Gender:         user.Gender,
		Password:       user.Password,
		PhoneNumber:    user.PhoneNumber,
		DocumentNumber: user.DocumentNumber,
		RewardPoints:   user.RewardPoints,
		Email:          user.Email,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}

	err := r.Db.Create(&userModel).Error

	if err != nil {
		return err
	}

	return nil
}
