package cli_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"myrnova/hexagonal/adapters/cli"
	mockapplication "myrnova/hexagonal/application/mocks"
	"testing"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product Name"
	productPrice := 25.99
	productStatus := "enabled"
	productId := "abc"

	productMock := mockapplication.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetId().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mockapplication.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID #%s with name %s has been created with price %.2f", productId, productName, productPrice)
	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Equal(t, resultExpected, result)
	require.Nil(t, err)

	resultExpected = fmt.Sprintf("Product #%s has been enabled", productName)
	result, err = cli.Run(service, "enable", productId, "", 0)
	require.Equal(t, resultExpected, result)
	require.Nil(t, err)

	resultExpected = fmt.Sprintf("Product #%s has been disabled", productName)
	result, err = cli.Run(service, "disable", productId, "", 0)
	require.Equal(t, resultExpected, result)
	require.Nil(t, err)

	resultExpected = fmt.Sprintf("Product ID %s\nName %s\nPrice %.2f\nStatus %s",
		productId, productName, productPrice, productStatus)
	result, err = cli.Run(service, "get", productId, "", 0)
	require.Equal(t, resultExpected, result)
	require.Nil(t, err)
}
