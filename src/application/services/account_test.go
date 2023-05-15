package services_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	mock_models "github.com/rodrigocardosodev/pismo-challenge/src/application/mocks/models"
	mock_ports "github.com/rodrigocardosodev/pismo-challenge/src/application/mocks/ports"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/models"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/services"
	"github.com/stretchr/testify/require"
)

func TestAccountService_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	account := mock_models.NewMockAccountInterface(ctrl)
	accountRepository := mock_ports.NewMockIAccountRepository(ctrl)
	transactionRepository := mock_ports.NewMockITransactionRepository(ctrl)
	accountRepository.EXPECT().GetByID(ctx, gomock.Any()).Return(account, nil)

	service := services.NewAccountService(accountRepository, transactionRepository)

	result, err := service.GetByID(ctx, 1)
	require.Nil(t, err)
	require.Equal(t, account, result)
}

func TestAccountService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	account := mock_models.NewMockAccountInterface(ctrl)
	accountRepository := mock_ports.NewMockIAccountRepository(ctrl)
	transactionRepository := mock_ports.NewMockITransactionRepository(ctrl)
	accountRepository.EXPECT().Create(ctx, gomock.Any()).Return(account, nil)
	accountRepository.EXPECT().GetByDocumentNumber(ctx, gomock.Any()).Return(nil, nil)

	service := services.NewAccountService(accountRepository, transactionRepository)

	result, err := service.Create(ctx, "95694057082")
	require.Nil(t, err)
	require.Equal(t, account, result)

	result, err = service.Create(ctx, "9569405708")
	require.NotNil(t, err)
	require.Nil(t, result)
	require.Equal(t, "cpf must have 11 digits", err.Error())

	result, err = service.Create(ctx, "12345678900")
	require.NotNil(t, err)
	require.Nil(t, result)
	require.Equal(t, "invalid cpf", err.Error())

	accountRepository.EXPECT().GetByDocumentNumber(ctx, gomock.Any()).Return(account, nil)
	result, err = service.Create(ctx, "95694057082")
	require.NotNil(t, err)
	require.Nil(t, result)
	require.Equal(t, "account already exists", err.Error())
}

func TestAccountService_GetBalanceAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	account := models.NewAccount("95694057082")

	accountRepository := mock_ports.NewMockIAccountRepository(ctrl)
	accountRepository.EXPECT().GetByID(ctx, gomock.Any()).Return(account, nil)
	transactionRepository := mock_ports.NewMockITransactionRepository(ctrl)
	transactionRepository.EXPECT().GetBalanceByAccountID(ctx, gomock.Any()).Return(1000.0, nil)

	service := services.NewAccountService(accountRepository, transactionRepository)

	result, err := service.GetAccountBalance(ctx, 1)
	require.Nil(t, err)
	require.Equal(t, 1000.0, result.GetBalance())
}
