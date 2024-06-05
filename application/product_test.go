package application_test

import (
	"testing"

	"github.com/DerivedPuma7/go-hexagonal/application"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello";
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()

	require.Nil(t, err)
}

func TestProduct_EnableError(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello";
	product.Status = application.DISABLED
	product.Price = 0

	err := product.Enable()

	require.Equal(t, "price must be greater than zero to enable product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello";
	product.Status = application.DISABLED
	product.Price = 0

	err := product.Disable()

	require.Nil(t, err)
}

func TestProduct_DisableError(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello";
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Disable()

	require.Equal(t, "price must be zero in order to disable a product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.New().String()
	product.Name = "Hello";
	product.Status = application.DISABLED
	product.Price = 10

	isValid, err := product.IsValid()
	require.Nil(t, err, "Mensagem de erro que facilita o entendimento da falha do teste")
	require.True(t, isValid)

	product.Status = "Invalid"
	_, err = product.IsValid()
	require.Equal(t, "status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err,)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "price must be greater or equal zero", err.Error())

	product.Price = 10
	product.ID = "any"
	_, err = product.IsValid()
	require.Equal(t, "ID: any does not validate as uuidv4", err.Error())
}

func TestProduct_GetName(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.New().String()
	product.Name = "Hello";
	product.Status = application.DISABLED
	product.Price = 10

	name := product.GetName()

	require.Equal(t, "Hello", name)
}

func TestProduct_GetID(t *testing.T) {
	uuid := uuid.New().String()
	product := application.Product{}
	product.ID = uuid
	product.Name = "Hello";
	product.Status = application.DISABLED
	product.Price = 10

	uuidReturn := product.GetID()

	require.Equal(t, uuid, uuidReturn)
}

func TestProduct_GetStatus(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.New().String()
	product.Name = "Hello";
	product.Status = application.DISABLED
	product.Price = 10

	status := product.GetStatus()

	require.Equal(t, "disabled", status)
}

func TestProduct_GetPrice(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.New().String()
	product.Name = "Hello";
	product.Status = application.DISABLED
	product.Price = 10

	price := product.GetPrice()

	require.Equal(t, 10.0, price)
}
