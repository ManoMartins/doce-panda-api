package repository

import (
	"doce-panda/domain/user/entity"
	"doce-panda/infra/db/gorm"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserRepository_Find_Success(t *testing.T) {
	db := gorm.NewDbTest()
	defer db.Close()

	user, _ := entity.NewUser(entity.UserProps{
		Name:           "John Doe",
		Email:          "john.doe@test.com",
		Gender:         "male",
		Password:       "p123",
		PhoneNumber:    "11981297480",
		DocumentNumber: "48358626860",
	})

	userRepository := UserRepositoryDb{Db: db}

	_, err := userRepository.Create(*user)

	require.Nil(t, err)

	userFound, err := userRepository.Find(user.ID)

	require.Equal(t, user.ID, userFound.ID)
}
