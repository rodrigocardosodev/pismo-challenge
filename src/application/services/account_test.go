package services_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_models "github.com/rodrigocardosodev/pismo-challenge/src/application/mocks/models"
	mock_ports "github.com/rodrigocardosodev/pismo-challenge/src/application/mocks/ports"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/services"
	"github.com/stretchr/testify/require"
)

func TestAccountService_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	account := mock_models.NewMockAccountInterface(ctrl)
	repository := mock_ports.NewMockIAccountRepository(ctrl)
	repository.EXPECT().GetByID(gomock.Any()).Return(account, nil)

	service := services.NewAccountService(repository)

	result, err := service.GetByID(1)
	require.Nil(t, err)
	require.Equal(t, account, result)
}

func TestAccountService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	account := mock_models.NewMockAccountInterface(ctrl)
	repository := mock_ports.NewMockIAccountRepository(ctrl)
	repository.EXPECT().Create(gomock.Any()).Return(account, nil)

	service := services.NewAccountService(repository)

	result, err := service.Create("557.242.030-14")
	require.Nil(t, err)
	require.Equal(t, account, result)

	_, err = service.Create("123.456.789-02")
	require.Equal(t, "cpf inválido", err.Error())
}
