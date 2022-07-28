package cli_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/mrdibre/hexagonal-arch-go/adapters/cli"
	"github.com/mrdibre/hexagonal-arch-go/application"
	mock_application "github.com/mrdibre/hexagonal-arch-go/application/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	productName   = "Product Test"
	productPrice  = 25.99
	productStatus = application.ENABLED
	productId     = "abc"
)

func setUp(ctrl *gomock.Controller) (*mock_application.MockProductInterface, *mock_application.MockProductServiceInterface) {
	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	return productMock, service
}

func TestRun_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	_, service := setUp(ctrl)

	resultExpected := fmt.Sprintf(
		"Product ID %s with the name %s has been created with the price %f and status %s",
		productId,
		productName,
		productPrice,
		productStatus,
	)

	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}

func TestRun_Enable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	_, service := setUp(ctrl)

	resultExpected := fmt.Sprintf("Product %s has been enabled", productName)
	result, err := cli.Run(service, "enable", productId, "", 0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}

func TestRun_Disable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	_, service := setUp(ctrl)

	resultExpected := fmt.Sprintf("Product %s has been disabled", productName)
	result, err := cli.Run(service, "disable", productId, "", 0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}

func TestRun_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	_, service := setUp(ctrl)

	resultExpected := fmt.Sprintf(
		"Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
		productId,
		productName,
		productPrice,
		productStatus,
	)

	result, err := cli.Run(service, "whatever", productId, "", 0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}
