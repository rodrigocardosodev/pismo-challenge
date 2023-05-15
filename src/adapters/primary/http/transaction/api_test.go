package transaction_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	http_transaction "github.com/rodrigocardosodev/pismo-challenge/src/adapters/primary/http/transaction"
	mock_models "github.com/rodrigocardosodev/pismo-challenge/src/application/mocks/models"
	mock_services "github.com/rodrigocardosodev/pismo-challenge/src/application/mocks/services"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/models"
	"github.com/stretchr/testify/require"
)

func TestHTTPTransactionAdapter_CreateTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	transaction := mock_models.NewMockTransactionInterface(ctrl)

	t.Run("Success CreateTransaction", func(t *testing.T) {
		service := mock_services.NewMockITrasactionService(ctrl)

		adapter := http_transaction.NewHTTPTransactionAdapter(service)
		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.POST("/transactions", adapter.CreateTransaction)
		service.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(transaction, nil).AnyTimes()

		requestBody := strings.NewReader(`{"account_id": 1, "operation_type_id": 1, "amount": 1000}`)
		httpRequest := httptest.NewRequest(http.MethodPost, "/transactions", requestBody)
		httpRequest.Header.Set("Content-Type", "application/json")

		responseWriter := httptest.NewRecorder()
		router.ServeHTTP(responseWriter, httpRequest)

		require.Equal(t, http.StatusCreated, responseWriter.Code)
	})

	t.Run("Error account not exists CreateTransaction", func(t *testing.T) {
		service := mock_services.NewMockITrasactionService(ctrl)

		adapter := http_transaction.NewHTTPTransactionAdapter(service)
		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.POST("/transactions", adapter.CreateTransaction)
		service.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, models.ErrAccountNotFound).AnyTimes()

		requestBody := strings.NewReader(`{"account_id": 1, "operation_type_id": 1, "amount": 1000}`)
		httpRequest := httptest.NewRequest(http.MethodPost, "/transactions", requestBody)
		httpRequest.Header.Set("Content-Type", "application/json")

		responseWriter := httptest.NewRecorder()
		router.ServeHTTP(responseWriter, httpRequest)

		require.Equal(t, http.StatusNotFound, responseWriter.Code)
		require.Equal(t, `{"error":"account not found"}`, responseWriter.Body.String())
	})

	t.Run("Error invalid operation type CreateTransaction", func(t *testing.T) {
		service := mock_services.NewMockITrasactionService(ctrl)

		adapter := http_transaction.NewHTTPTransactionAdapter(service)
		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.POST("/transactions", adapter.CreateTransaction)
		service.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, models.ErrInvalidOperationType).AnyTimes()

		requestBody := strings.NewReader(`{"account_id": 1, "operation_type_id": 5, "amount": 1000}`)
		httpRequest := httptest.NewRequest(http.MethodPost, "/transactions", requestBody)
		httpRequest.Header.Set("Content-Type", "application/json")

		responseWriter := httptest.NewRecorder()
		router.ServeHTTP(responseWriter, httpRequest)

		require.Equal(t, http.StatusBadRequest, responseWriter.Code)
		require.Equal(t, `{"error":"invalid operation type"}`, responseWriter.Body.String())
	})

	t.Run("Error invalid amount CreateTransaction", func(t *testing.T) {
		service := mock_services.NewMockITrasactionService(ctrl)
		adapter := http_transaction.NewHTTPTransactionAdapter(service)
		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.POST("/transactions", adapter.CreateTransaction)
		service.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, models.ErrInvalidAmountByOperationType).AnyTimes()

		requestBody := strings.NewReader(`{"account_id": 1, "operation_type_id": 1, "amount": 0}`)
		httpRequest := httptest.NewRequest(http.MethodPost, "/transactions", requestBody)
		httpRequest.Header.Set("Content-Type", "application/json")

		responseWriter := httptest.NewRecorder()
		router.ServeHTTP(responseWriter, httpRequest)
		require.Equal(t, http.StatusBadRequest, responseWriter.Code)
		require.Equal(t, `{"error":"invalid amount by operation type"}`, responseWriter.Body.String())
	})

	t.Run("Error default in CreateTransaction", func(t *testing.T) {
		service := mock_services.NewMockITrasactionService(ctrl)

		adapter := http_transaction.NewHTTPTransactionAdapter(service)
		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.POST("/transactions", adapter.CreateTransaction)
		service.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("a weird error")).AnyTimes()

		requestBody := strings.NewReader(`{"account_id": 1, "operation_type_id": 1, "amount": 1000}`)
		httpRequest := httptest.NewRequest(http.MethodPost, "/transactions", requestBody)
		httpRequest.Header.Set("Content-Type", "application/json")

		responseWriter := httptest.NewRecorder()
		router.ServeHTTP(responseWriter, httpRequest)

		require.Equal(t, http.StatusInternalServerError, responseWriter.Code)
	})
}
