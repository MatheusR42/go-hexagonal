package application_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/matheusr42/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	status := application.DISABLED

	product := application.Product{
		Name:   "Example",
		Status: &status,
		Price:  10,
	}

	err := product.Enable()
	require.Nil(t, err)
	require.Equal(t, application.ENABLED, product.GetStatus())

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	status := application.DISABLED

	product := application.Product{
		Name:   "Example",
		Status: &status,
		Price:  0,
	}

	err := product.Disable()
	require.Nil(t, err)
	require.Equal(t, application.DISABLED, product.GetStatus())

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero to disable product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	application.Init()
	status := application.DISABLED

	product := application.Product{
		ID:     uuid.New().String(),
		Name:   "Example",
		Status: &status,
		Price:  0,
	}

	valid, err := product.IsValid()
	require.Equal(t, true, valid)
	require.Nil(t, err)

	product.Price = -10
	valid, err = product.IsValid()
	require.Equal(t, false, valid)
	require.NotNil(t, err)
}
