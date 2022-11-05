package repository

import (
	"doce-panda/domain/user/entity"
	"doce-panda/infra/db/gorm"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddressRepository_Find_Success(t *testing.T) {
	db := gorm.NewDbTest()
	defer db.Close()

	user, _ := entity.NewUser(entity.User{
		Name:           "John Doe",
		Email:          "john.doe@test.com",
		Gender:         "male",
		Password:       "p123",
		PhoneNumber:    "11981297480",
		DocumentNumber: "48358626860",
	})

	userRepository := UserRepositoryDb{Db: db}

	err := userRepository.Create(*user)

	address, err := entity.NewAddress(entity.Address{
		City:         "Suzano",
		State:        "São Paulo",
		Street:       "Rua Bandeirantes",
		Number:       "1140",
		ZipCode:      "08694180",
		Neighborhood: "Jardim revista",
		IsMain:       false,
	})

	addressRepository := AddressRepositoryDb{Db: db}

	err = addressRepository.Create(*address)

	require.Nil(t, err)

	addressFound, err := addressRepository.FindById(address.ID)

	require.Equal(t, address.ID, addressFound.ID)

	userFound, err := userRepository.FindById(user.ID)
	fmt.Println(userFound)

}
