package application_test

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := Product{
		Price:  10,
		Name:   "Hello World",
		Status: ENABLED,
	}
	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than equal to zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := Product{
		Price:  0,
		Name:   "Hello World",
		Status: ENABLED,
	}
	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be equal to zero to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := Product{
		ID:     uuid.NewV4().String(),
		Price:  0,
		Name:   "Hello World",
		Status: ENABLED,
	}
	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "invalid"
	_, err = product.IsValid()
	require.Equal(t, "invalid product status", err.Error())

	product.Status = ENABLED
	product.Price = -1
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal to zero", err.Error())

	product.Price = 0
	product.Name = ""
	_, err = product.IsValid()
	require.Equal(t, "Name: non zero value required", err.Error())
}
