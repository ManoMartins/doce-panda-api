package repository

import (
	"doce-panda/domain/product/entity"
	"doce-panda/infra/db/gorm"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProductRepository_Create_Success(t *testing.T) {
	db := gorm.NewDbTest()
	defer db.Close()

	product, _ := entity.NewProduct(entity.ProductProps{
		Name:         "Bolo de pote",
		PriceInCents: 750,
		Description:  "Um bolo de pote de chocolate",
		Flavor:       "chocolate",
		Quantity:     5,
	})

	productRepository := ProductRepositoryDb{Db: db}

	_, err := productRepository.Create(*product)

	require.Nil(t, err)

	productFound, err := productRepository.Find(product.ID)

	require.Equal(t, product.ID, productFound.ID)
}

func TestProductRepository_Update_Success(t *testing.T) {
	db := gorm.NewDbTest()
	defer db.Close()

	product, _ := entity.NewProduct(entity.ProductProps{
		Name:         "Bolo de pote",
		PriceInCents: 750,
		Description:  "Um bolo de pote de chocolate",
		Flavor:       "chocolate",
		Quantity:     5,
	})

	productRepository := ProductRepositoryDb{Db: db}

	_, err := productRepository.Create(*product)

	require.Nil(t, err)

	productFound, err := productRepository.Find(product.ID)

	productFound.PriceInCents = 950

	productUpdate := entity.Product{
		ID:           productFound.ID,
		Name:         productFound.Name,
		PriceInCents: 950,
		Status:       entity.StatusEnum(productFound.Status),
		Description:  productFound.Description,
		Flavor:       productFound.Flavor,
		Quantity:     productFound.Quantity,
	}
	_, err = productRepository.Update(productUpdate)

	require.Equal(t, 950, productFound.PriceInCents)
}

func TestProductRepository_FindAll_Success(t *testing.T) {
	db := gorm.NewDbTest()
	defer db.Close()

	product, _ := entity.NewProduct(entity.ProductProps{
		Name:         "Bolo de pote",
		PriceInCents: 750,
		Description:  "Um bolo de pote de chocolate",
		Flavor:       "chocolate",
		Quantity:     5,
	})
	product2, _ := entity.NewProduct(entity.ProductProps{
		Name:         "Bolo de pote",
		PriceInCents: 950,
		Description:  "Um bolo de pote de chocolate",
		Flavor:       "ninho",
		Quantity:     5,
	})
	product3, _ := entity.NewProduct(entity.ProductProps{
		Name:         "Donuts",
		PriceInCents: 1000,
		Description:  "Um bolo de pote de chocolate",
		Flavor:       "chocolate",
		Quantity:     5,
	})

	productRepository := ProductRepositoryDb{Db: db}

	_, err := productRepository.Create(*product)
	_, err = productRepository.Create(*product2)
	_, err = productRepository.Create(*product3)

	require.Nil(t, err)

	productsFound, err := productRepository.FindAll()

	require.NotEmpty(t, productsFound)
}
