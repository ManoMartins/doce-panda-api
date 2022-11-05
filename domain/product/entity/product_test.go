package entity_test

import (
	"doce-panda/domain/product/entity"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProductValidate_EmptyProps_ErrorRequiredFields(t *testing.T) {
	product, _ := entity.NewProduct(entity.ProductProps{
		Name:         "",
		PriceInCents: 0,
		Description:  "",
		Flavor:       "",
		Quantity:     0,
	})

	err := product.Validate()

	require.NotNil(t, err)
}

func TestProductValidate_FillProps_Success(t *testing.T) {
	product, _ := entity.NewProduct(entity.ProductProps{
		Name:         "Bolo de pote",
		PriceInCents: 750,
		Description:  "Bolo de pote sabor de chocolate",
		Flavor:       "chocolate",
		Quantity:     5,
	})

	err := product.Validate()

	require.Nil(t, err)
}

func TestProduct_CreateProduct_Success(t *testing.T) {
	product, _ := entity.NewProduct(entity.ProductProps{
		Name:         "Bolo de pote",
		PriceInCents: 750,
		Description:  "Bolo de pote sabor de chocolate",
		Flavor:       "chocolate",
		Quantity:     5,
	})

	require.Equal(t, product.Name, "Bolo de pote")
	require.Equal(t, product.PriceInCents, 750)
	require.Equal(t, product.Description, "Bolo de pote sabor de chocolate")
	require.Equal(t, product.Flavor, "chocolate")
	require.Equal(t, product.Quantity, 5)
}
