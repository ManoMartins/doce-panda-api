package find

import (
	"doce-panda/domain/product/entity"
	"doce-panda/infra/db/gorm"
	"doce-panda/infra/gorm/product/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindProduct_Success(t *testing.T) {
	db := gorm.NewDbTest()
	defer db.Close()

	productRepository := repository.ProductRepositoryDb{Db: db}
	usecase := NewFindProductUseCase(productRepository)

	product, err := entity.NewProduct(entity.ProductProps{
		Name:         "Bolo de pote",
		PriceInCents: 750,
		Description:  "Um bolo de pote",
		Flavor:       "chocolate",
		Quantity:     10,
	})

	input := InputFindProductDto{ID: product.ID}

	_, err = productRepository.Create(*product)

	require.Nil(t, err)

	productFound, err := usecase.Execute(input)

	require.Equal(t, product.ID, productFound.ID)
}
