package cli_test

import (
	"fmt"
	"testing"

	"github.com/DerivedPuma7/go-hexagonal/adapters/cli"
	mocks "github.com/DerivedPuma7/go-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var productName = "Product Test"
var productPrice = 25.99
var productStatus = "ENABLED"
var productId = "abc"


func buildMocks(ctrl *gomock.Controller) (*mocks.MockProductInterface, *mocks.MockProductServiceInterface) {
	productMock := mocks.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()

	service := mocks.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	return productMock, service
}


func TestRun_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	_, service := buildMocks(ctrl)
	expectedCreateResult := fmt.Sprintf(`Product ID %s with the name %s has been created with price %f and status %s`, productId, productName, productPrice, productStatus)

	result, err := cli.Run(service, "create", "", productName, productPrice)

	require.Nil(t, err)
	require.Equal(t, expectedCreateResult, result)
}

func TestRun_Enable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	_, service := buildMocks(ctrl)
	expectedEnableResult := fmt.Sprintf(`Product %s has been enabled`, productName)

	result, err := cli.Run(service, "enable", productId, "", 0)

	require.Nil(t, err)
	require.Equal(t, expectedEnableResult, result)
}

func TestRun_Disable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	_, service := buildMocks(ctrl)
	expectedEnableResult := fmt.Sprintf(`Product %s has been disabled`, productName)

	result, err := cli.Run(service, "disable", productId, "", 0)

	require.Nil(t, err)
	require.Equal(t, expectedEnableResult, result)
}

func TestRun_Default(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	_, service := buildMocks(ctrl)
	expectedEnableResult := fmt.Sprintf("Product ID: %s \nName: %s \nPrice: %f \nStatus: %s", productId, productName, productPrice, productStatus)

	result, err := cli.Run(service, "", productId, "", 0)

	require.Nil(t, err)
	require.Equal(t, expectedEnableResult, result)
}