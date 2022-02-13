package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/matheusr42/go-hexagonal/application"
	mock_application "github.com/matheusr42/go-hexagonal/application/mocks"
	"github.com/matheusr42/go-hexagonal/cli"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	productStatus := application.ENABLED

	product := application.Product{
		Name:   "Product",
		Price:  10.0,
		Status: &productStatus,
		ID:     "123",
	}

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(product.ID).AnyTimes()
	productMock.EXPECT().GetName().Return(product.Name).AnyTimes()
	productMock.EXPECT().GetStatus().Return(*product.Status).AnyTimes()
	productMock.EXPECT().GetPrice().Return(product.Price).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(product.Name, product.Price).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(product.ID).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	expectedResult := "product created"
	result, err := cli.Run(service, cli.CREATE, "", product.Name, product.Price)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = "product enabled"
	result, err = cli.Run(service, cli.ENABLE, product.ID, "", product.Price)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf("product: %s", product.GetName())
	result, err = cli.Run(service, cli.GET, product.ID, "", product.Price)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)
}
