package account_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	http_account "github.com/rodrigocardosodev/pismo-challenge/src/adapters/primary/http/account"
	mock_models "github.com/rodrigocardosodev/pismo-challenge/src/application/mocks/models"
	mock_services "github.com/rodrigocardosodev/pismo-challenge/src/application/mocks/services"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/models"
	"github.com/stretchr/testify/require"
)

func TestHTTPAccountAdapter_CreateAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	account := mock_models.NewMockAccountInterface(ctrl)

	t.Run("Success CreateAccount", func(t *testing.T) {
		service := mock_services.NewMockIAccountService(ctrl)
		service.EXPECT().Create(gomock.Any(), gomock.Any()).Return(account, nil)

		adapter := http_account.NewHTTPAccountAdapter(service)

		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.POST("/accounts", adapter.CreateAccount)
		requestBody := strings.NewReader(`{"document_number": "12345678900"}`)
		httpRequest := httptest.NewRequest(http.MethodPost, "/accounts", requestBody)
		httpRequest.Header.Set("Content-Type", "application/json")

		responseWriter := httptest.NewRecorder()
		router.ServeHTTP(responseWriter, httpRequest)

		require.Equal(t, http.StatusCreated, responseWriter.Code)
	})

	t.Run("Error invalid body CreateAccount", func(t *testing.T) {
		service := mock_services.NewMockIAccountService(ctrl)
		adapter := http_account.NewHTTPAccountAdapter(service)

		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.POST("/accounts", adapter.CreateAccount)
		requestBody := strings.NewReader(`"document_number": "12345678900"}`)
		httpRequest := httptest.NewRequest(http.MethodPost, "/accounts", requestBody)
		httpRequest.Header.Set("Content-Type", "application/json")

		responseWriter := httptest.NewRecorder()
		router.ServeHTTP(responseWriter, httpRequest)

		require.Equal(t, http.StatusBadRequest, responseWriter.Code)
		require.Equal(t, `{"error":"invalid request body account"}`, responseWriter.Body.String())
	})

	t.Run("Error account already exists CreateAccount", func(t *testing.T) {
		service := mock_services.NewMockIAccountService(ctrl)
		service.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, models.ErrAccountAlreadyExists)

		adapter := http_account.NewHTTPAccountAdapter(service)

		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.POST("/accounts", adapter.CreateAccount)
		requestBody := strings.NewReader(`{"document_number": "12345678900"}`)
		httpRequest := httptest.NewRequest(http.MethodPost, "/accounts", requestBody)
		httpRequest.Header.Set("Content-Type", "application/json")

		responseWriter := httptest.NewRecorder()
		router.ServeHTTP(responseWriter, httpRequest)

		require.Equal(t, http.StatusBadRequest, responseWriter.Code)
	})

	t.Run("Error invalid cpf CreateAccount", func(t *testing.T) {
		service := mock_services.NewMockIAccountService(ctrl)
		service.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, models.ErrInvalidCpf)

		adapter := http_account.NewHTTPAccountAdapter(service)

		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.POST("/accounts", adapter.CreateAccount)
		requestBody := strings.NewReader(`{"document_number": "12345678900"}`)
		httpRequest := httptest.NewRequest(http.MethodPost, "/accounts", requestBody)
		httpRequest.Header.Set("Content-Type", "application/json")

		responseWriter := httptest.NewRecorder()
		router.ServeHTTP(responseWriter, httpRequest)

		require.Equal(t, http.StatusBadRequest, responseWriter.Code)
		require.Equal(t, `{"error":"invalid cpf"}`, responseWriter.Body.String())
	})

	t.Run("Error cpf must have 11 digits CreateAccount", func(t *testing.T) {
		service := mock_services.NewMockIAccountService(ctrl)
		service.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, models.ErrCpfMustHave11Digits)

		adapter := http_account.NewHTTPAccountAdapter(service)

		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.POST("/accounts", adapter.CreateAccount)
		requestBody := strings.NewReader(`{"document_number": "12345678900"}`)
		httpRequest := httptest.NewRequest(http.MethodPost, "/accounts", requestBody)
		httpRequest.Header.Set("Content-Type", "application/json")

		responseWriter := httptest.NewRecorder()
		router.ServeHTTP(responseWriter, httpRequest)

		require.Equal(t, http.StatusBadRequest, responseWriter.Code)
		require.Equal(t, `{"error":"cpf must have 11 digits"}`, responseWriter.Body.String())
	})

	t.Run("Error CPF must have only digits CreateAccount", func(t *testing.T) {
		service := mock_services.NewMockIAccountService(ctrl)
		service.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, models.ErrCpfMustHaveOnlyDigits)

		adapter := http_account.NewHTTPAccountAdapter(service)

		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.POST("/accounts", adapter.CreateAccount)
		requestBody := strings.NewReader(`{"document_number": "asdfghjklqw"}`)
		httpRequest := httptest.NewRequest(http.MethodPost, "/accounts", requestBody)
		httpRequest.Header.Set("Content-Type", "application/json")

		responseWriter := httptest.NewRecorder()
		router.ServeHTTP(responseWriter, httpRequest)
		require.Equal(t, http.StatusBadRequest, responseWriter.Code)
		require.Equal(t, `{"error":"cpf must have only digits"}`, responseWriter.Body.String())
	})

	t.Run("Error internal server error CreateAccount", func(t *testing.T) {
		service := mock_services.NewMockIAccountService(ctrl)
		service.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New("internal server error"))

		adapter := http_account.NewHTTPAccountAdapter(service)

		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.POST("/accounts", adapter.CreateAccount)
		requestBody := strings.NewReader(`{"document_number": "asdfghjklqw"}`)
		httpRequest := httptest.NewRequest(http.MethodPost, "/accounts", requestBody)
		httpRequest.Header.Set("Content-Type", "application/json")

		responseWriter := httptest.NewRecorder()
		router.ServeHTTP(responseWriter, httpRequest)
		require.Equal(t, http.StatusInternalServerError, responseWriter.Code)
		require.Equal(t, `{"error":"internal server error"}`, responseWriter.Body.String())
	})
}

func TestHTTPAccountAdapter_GetAccountByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	account := models.NewAccount("95694057082")
	account.SetID(1)

	t.Run("Success GetAccountByID", func(t *testing.T) {
		service := mock_services.NewMockIAccountService(ctrl)
		service.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(account, nil)

		adapter := http_account.NewHTTPAccountAdapter(service)

		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.GET("/accounts/:account_id", adapter.GetAccountById)
		httpRequest := httptest.NewRequest(http.MethodGet, "/accounts/1", nil)
		httpRequest.Header.Set("Content-Type", "application/json")

		responseWriter := httptest.NewRecorder()
		router.ServeHTTP(responseWriter, httpRequest)

		require.Equal(t, http.StatusOK, responseWriter.Code)
		require.Equal(t, `{"id":1,"document_number":"95694057082"}`, responseWriter.Body.String())
	})

	t.Run("Error account not found GetAccountByID", func(t *testing.T) {
		service := mock_services.NewMockIAccountService(ctrl)
		service.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(nil, models.ErrAccountNotFound).AnyTimes()

		adapter := http_account.NewHTTPAccountAdapter(service)

		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.GET("/accounts/:account_id", adapter.GetAccountById)
		httpRequest := httptest.NewRequest(http.MethodGet, "/accounts/1", nil)
		httpRequest.Header.Set("Content-Type", "application/json")

		responseWriter := httptest.NewRecorder()
		router.ServeHTTP(responseWriter, httpRequest)

		require.Equal(t, http.StatusNotFound, responseWriter.Code)
		require.Equal(t, `{"error":"account not found"}`, responseWriter.Body.String())
	})

	t.Run("Error invalid account id GetAccountByID", func(t *testing.T) {
		service := mock_services.NewMockIAccountService(ctrl)
		service.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(nil, nil).AnyTimes()
		adapter := http_account.NewHTTPAccountAdapter(service)

		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.GET("/accounts/:account_id", adapter.GetAccountById)
		httpRequest := httptest.NewRequest(http.MethodGet, "/accounts/asd", nil)
		httpRequest.Header.Set("Content-Type", "application/json")

		responseWriter := httptest.NewRecorder()
		router.ServeHTTP(responseWriter, httpRequest)

		require.Equal(t, http.StatusBadRequest, responseWriter.Code)
		require.Equal(t, `{"error":"invalid account id"}`, responseWriter.Body.String())
	})

	t.Run("Error invalid account id GetAccountByID", func(t *testing.T) {
		service := mock_services.NewMockIAccountService(ctrl)
		service.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(nil, errors.New("internal server error")).AnyTimes()
		adapter := http_account.NewHTTPAccountAdapter(service)

		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.GET("/accounts/:account_id", adapter.GetAccountById)
		httpRequest := httptest.NewRequest(http.MethodGet, "/accounts/1", nil)
		httpRequest.Header.Set("Content-Type", "application/json")

		responseWriter := httptest.NewRecorder()
		router.ServeHTTP(responseWriter, httpRequest)

		require.Equal(t, http.StatusInternalServerError, responseWriter.Code)
	})
}

func TestHTTPAccountAdapter_GetAccountBalance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	account := models.NewAccount("95694057082")
	account.SetID(1)

	t.Run("Success GetAccountBalance", func(t *testing.T) {
		service := mock_services.NewMockIAccountService(ctrl)
		service.EXPECT().GetAccountBalance(gomock.Any(), gomock.Any()).Return(account, nil)

		adapter := http_account.NewHTTPAccountAdapter(service)

		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.GET("/accounts/:account_id/balance", adapter.GetAccountBalance)
		httpRequest := httptest.NewRequest(http.MethodGet, "/accounts/1/balance", nil)
		httpRequest.Header.Set("Content-Type", "application/json")

		responseWriter := httptest.NewRecorder()
		router.ServeHTTP(responseWriter, httpRequest)

		require.Equal(t, http.StatusOK, responseWriter.Code)
		require.Equal(t, `{"id":1,"document_number":"95694057082","balance":0}`, responseWriter.Body.String())
	})

	t.Run("Error account not found GetAccountBalance", func(t *testing.T) {
		service := mock_services.NewMockIAccountService(ctrl)
		service.EXPECT().GetAccountBalance(gomock.Any(), gomock.Any()).Return(nil, models.ErrAccountNotFound).AnyTimes()

		adapter := http_account.NewHTTPAccountAdapter(service)
		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.GET("/accounts/:account_id/balance", adapter.GetAccountBalance)
		httpRequest := httptest.NewRequest(http.MethodGet, "/accounts/1/balance", nil)
		httpRequest.Header.Set("Content-Type", "application/json")

		responseWriter := httptest.NewRecorder()
		router.ServeHTTP(responseWriter, httpRequest)

		require.Equal(t, http.StatusNotFound, responseWriter.Code)
		require.Equal(t, `{"error":"account not found"}`, responseWriter.Body.String())
	})

	t.Run("Error invalid account id GetAccountBalance", func(t *testing.T) {
		service := mock_services.NewMockIAccountService(ctrl)
		service.EXPECT().GetAccountBalance(gomock.Any(), gomock.Any()).Return(nil, nil).AnyTimes()

		adapter := http_account.NewHTTPAccountAdapter(service)
		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.GET("/accounts/:account_id/balance", adapter.GetAccountBalance)
		httpRequest := httptest.NewRequest(http.MethodGet, "/accounts/asd/balance", nil)
		httpRequest.Header.Set("Content-Type", "application/json")

		responseWriter := httptest.NewRecorder()
		router.ServeHTTP(responseWriter, httpRequest)

		require.Equal(t, http.StatusBadRequest, responseWriter.Code)
		require.Equal(t, `{"error":"invalid account id"}`, responseWriter.Body.String())
	})

	t.Run("Error internal server error GetAccountBalance", func(t *testing.T) {
		service := mock_services.NewMockIAccountService(ctrl)
		service.EXPECT().GetAccountBalance(gomock.Any(), gomock.Any()).Return(nil, errors.New("internal server error")).AnyTimes()

		adapter := http_account.NewHTTPAccountAdapter(service)
		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.GET("/accounts/:account_id/balance", adapter.GetAccountBalance)
		httpRequest := httptest.NewRequest(http.MethodGet, "/accounts/1/balance", nil)
		httpRequest.Header.Set("Content-Type", "application/json")

		responseWriter := httptest.NewRecorder()
		router.ServeHTTP(responseWriter, httpRequest)

		require.Equal(t, http.StatusInternalServerError, responseWriter.Code)
	})
}
