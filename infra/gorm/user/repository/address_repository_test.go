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

	address, err := entity.NewAddress(entity.AddressProps{
		City:         "Suzano",
		State:        "São Paulo",
		Street:       "Rua Bandeirantes",
		Number:       "1140",
		ZipCode:      "08694180",
		Neighborhood: "Jardim revista",
		IsMain:       false,
		UserId:       user.ID,
	})

	addressRepository := AddressRepositoryDb{Db: db}

	_, err = addressRepository.Create(*address)

	require.Nil(t, err)

	addressFound, err := addressRepository.Find(address.ID)

	require.Equal(t, address.ID, addressFound.ID)

	userFound, err := userRepository.Find(user.ID)
	fmt.Println(userFound)

}
