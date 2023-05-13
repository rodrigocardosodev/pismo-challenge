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

func TestTransactionService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	transaction := mock_models.NewMockTransactionInterface(ctrl)
	repository := mock_ports.NewMockITransactionRepository(ctrl)
	repository.EXPECT().Create(ctx, gomock.Any()).Return(transaction, nil)

	service := services.NewTransactionService(repository)

	result, err := service.Create(ctx, 1, models.PAGAMENTO, 1000)
	require.Nil(t, err)
	require.Equal(t, transaction, result)

	result, err = service.Create(ctx, 1, models.SAQUE, 0)
	require.NotNil(t, err)
	require.Nil(t, result)

	result, err = service.Create(ctx, 1, 5, 1000)
	require.NotNil(t, err)
	require.Nil(t, result)
	require.Equal(t, "invalid operation type", err.Error())

	result, err = service.Create(ctx, 1, models.PAGAMENTO, -1000)
	require.NotNil(t, err)
	require.Nil(t, result)
	require.Equal(t, "invalid amount", err.Error())
}
