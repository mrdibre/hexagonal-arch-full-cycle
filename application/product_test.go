package application

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(testing *testing.T) {
	product := Product{}

	product.Name = "Hello"
	product.Status = DISABLED
	product.Price = 10

	err := product.Enable()

	require.Nil(testing, err)

	product.Price = 0
	err = product.Enable()

	require.Equal(testing, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(testing *testing.T) {
	product := Product{}

	product.Name = "Hello"
	product.Status = ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(testing, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(testing, "the price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(testing *testing.T) {
	product := Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "hello"
	product.Status = DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(testing, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(testing, "the status must be enabled or disabled", err.Error())

	product.Status = ENABLED
	_, err = product.IsValid()
	require.Nil(testing, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(testing, "the price must be greater or equal zero", err.Error())

}
