package application_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"myrnova/hexagonal/application"
	mockapplication "myrnova/hexagonal/application/mocks"
	"testing"
)

func TestProductService_Get(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mockapplication.NewMockProductInterface(ctrl)
	persistence := mockapplication.NewMockProductPersistentInterface(ctrl)

	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).Times(1)

	productService := application.ProductService{
		Persistence: persistence,
	}

	result, err := productService.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)

	persistence.EXPECT().Get(gomock.Any()).Return(nil, errors.New("error getting product")).Times(1)

	result, err = productService.Get("abc")
	require.Equal(t, "error getting product", err.Error())
	require.Nil(t, result)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mockapplication.NewMockProductInterface(ctrl)
	persistence := mockapplication.NewMockProductPersistentInterface(ctrl)

	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).Times(1)

	productService := application.ProductService{
		Persistence: persistence,
	}

	result, err := productService.Create("abc", 1)
	require.Nil(t, err)
	require.Equal(t, product, result)

	persistence.EXPECT().Save(gomock.Any()).Return(nil, errors.New("error getting product")).Times(1)

	result, err = productService.Create("abc", 2)
	require.Equal(t, "error getting product", err.Error())
	require.Nil(t, result)

	result, err = productService.Create("abc", -1)
	require.Equal(t, "the price must be greater or equal to zero", err.Error())
	require.Nil(t, result)
}

func TestProductService_Enable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mockapplication.NewMockProductInterface(ctrl)
	persistence := mockapplication.NewMockProductPersistentInterface(ctrl)

	product.EXPECT().Enable().Return(nil).Times(1)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).Times(1)

	productService := ProductService{
		Persistence: persistence,
	}

	result, err := productService.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

	product.EXPECT().Enable().Return(errors.New("new error")).Times(1)
	result, err = productService.Enable(product)
	require.Equal(t, "new error", err.Error())
	require.Nil(t, result)

}

func TestProductService_Disable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mockapplication.NewMockProductInterface(ctrl)
	persistence := mockapplication.NewMockProductPersistentInterface(ctrl)

	product.EXPECT().Disable().Return(nil).Times(1)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).Times(1)

	productService := application.ProductService{
		Persistence: persistence,
	}

	result, err := productService.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

	product.EXPECT().Disable().Return(errors.New("new error")).Times(1)
	result, err = productService.Disable(product)
	require.Equal(t, "new error", err.Error())
	require.Nil(t, result)

}
