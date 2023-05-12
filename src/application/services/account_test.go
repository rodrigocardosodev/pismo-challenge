package services_test

import (
	"context"
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
	ctx := context.Background()

	account := mock_models.NewMockAccountInterface(ctrl)
	repository := mock_ports.NewMockIAccountRepository(ctrl)
	repository.EXPECT().GetByID(ctx, gomock.Any()).Return(account, nil)

	service := services.NewAccountService(repository)

	result, err := service.GetByID(ctx, 1)
	require.Nil(t, err)
	require.Equal(t, account, result)
}

func TestAccountService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	account := mock_models.NewMockAccountInterface(ctrl)
	repository := mock_ports.NewMockIAccountRepository(ctrl)
	repository.EXPECT().Create(ctx, gomock.Any()).Return(account, nil)
	repository.EXPECT().GetByDocumentNumber(ctx, gomock.Any()).Return(nil, nil)

	service := services.NewAccountService(repository)

	result, err := service.Create(ctx, "05803828343")
	require.Nil(t, err)
	require.Equal(t, account, result)

	result, err = service.Create(ctx, "0580382834")
	require.NotNil(t, err)
	require.Nil(t, result)
	require.Equal(t, "cpf must have 11 digits", err.Error())

	result, err = service.Create(ctx, "12345678900")
	require.NotNil(t, err)
	require.Nil(t, result)
	require.Equal(t, "invalid cpf", err.Error())

	repository.EXPECT().GetByDocumentNumber(ctx, gomock.Any()).Return(account, nil)
	result, err = service.Create(ctx, "05803828343")
	require.NotNil(t, err)
	require.Nil(t, result)
	require.Equal(t, "account already exists", err.Error())
}
