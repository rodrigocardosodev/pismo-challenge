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
	repository.EXPECT().Create(ctx, gomock.Any()).Return(transaction, nil).AnyTimes()

	service := services.NewTransactionService(repository)

	t.Run("should create transaction", func(t *testing.T) {
		result, err := service.Create(ctx, 1, models.PAGAMENTO, 1000)
		require.Nil(t, err)
		require.Equal(t, transaction, result)
	})

	t.Run("should return amount not equal 0 error", func(t *testing.T) {
		result, err := service.Create(ctx, 1, models.SAQUE, 0)
		require.NotNil(t, err)
		require.Nil(t, result)
		require.Equal(t, "amount: non zero value required", err.Error())
	})

	t.Run("should return amount not equal 0 error", func(t *testing.T) {
		result, err := service.Create(ctx, 1, 5, 1000)
		require.NotNil(t, err)
		require.Nil(t, result)
		require.Equal(t, "invalid operation type", err.Error())
	})

	t.Run("should return error from positive amount in SAQUE", func(t *testing.T) {
		result, err := service.Create(ctx, 1, models.SAQUE, 1000)
		require.NotNil(t, err)
		require.Nil(t, result)
		require.Equal(t, "invalid amount by operation type", err.Error())
	})

	t.Run("should return error from positive amount in COMPRA_A_VISTA", func(t *testing.T) {
		result, err := service.Create(ctx, 1, models.COMPRA_A_VISTA, 1000)
		require.NotNil(t, err)
		require.Nil(t, result)
		require.Equal(t, "invalid amount by operation type", err.Error())
	})

	t.Run("should return error from negative amount in COMPRA_PARCELADA", func(t *testing.T) {
		result, err := service.Create(ctx, 1, models.COMPRA_PARCELADA, 1000)
		require.NotNil(t, err)
		require.Nil(t, result)
		require.Equal(t, "invalid amount by operation type", err.Error())
	})

	t.Run("should return error from negative amount in PAGAMENTO", func(t *testing.T) {
		result, err := service.Create(ctx, 1, models.PAGAMENTO, -1000)
		require.NotNil(t, err)
		require.Nil(t, result)
		require.Equal(t, "invalid amount by operation type", err.Error())
	})
}
