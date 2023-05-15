package account

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/dtos"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/models"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/services"
)

type HTTPAccountAdapter struct {
	service services.IAccountService
}

func NewHTTPAccountAdapter(service services.IAccountService) *HTTPAccountAdapter {
	return &HTTPAccountAdapter{service: service}
}

// AddAccount godoc
//
//	@Summary		Create account
//	@Description	Create account
//	@Tags			Accounts
//	@Accept			json
//	@Produce		json
//	@Param			account	body		dtos.AccountRequest	true	"Add account"
//	@Success		201		{object}	dtos.AccountResponse
//	@Failure		400		{object}	dtos.HTTPError
//	@Failure		404		{object}	dtos.HTTPError
//	@Failure		500		{object}	dtos.HTTPError
//	@Router			/accounts [post]
func (svc *HTTPAccountAdapter) CreateAccount(c *gin.Context) {
	var accountRequest dtos.AccountRequest
	if err := c.ShouldBindJSON(&accountRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	account, err := svc.service.Create(c, accountRequest.DocumentNumber)
	if err == models.ErrCpfMustHaveOnlyDigits || err == models.ErrCpfMustHave11Digits || err == models.ErrInvalidCpf || err == models.ErrAccountAlreadyExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, account)
}

// ShowAccount godoc
//
//	@Summary		Get account by ID
//	@Description	Get account by ID
//	@Tags			Accounts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Account ID"
//	@Success		200	{object}	dtos.AccountResponse
//	@Failure		400	{object}	dtos.HTTPError
//	@Failure		404	{object}	dtos.HTTPError
//	@Failure		500	{object}	dtos.HTTPError
//	@Router			/accounts/{id} [get]
func (svc *HTTPAccountAdapter) GetAccountById(c *gin.Context) {
	accountID := c.Param("account_id")
	parsedAccountID, err := strconv.ParseInt(accountID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid account id"})
		return
	}
	account, err := svc.service.GetByID(c, parsedAccountID)
	if err == models.ErrAccountNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

// ShowBalanceAccount godoc
//
//	@Summary		Get balance account by Account ID
//	@Description	get balance account by account id
//	@Tags			Accounts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Account ID"
//	@Success		200	{object}	dtos.AccountBalanceResponse
//	@Failure		400	{object}	dtos.HTTPError
//	@Failure		404	{object}	dtos.HTTPError
//	@Failure		500	{object}	dtos.HTTPError
//	@Router			/accounts/{id}/balance [get]
func (svc *HTTPAccountAdapter) GetAccountBalance(c *gin.Context) {
	var accountBalanceResponse dtos.AccountBalanceResponse
	accountID := c.Param("account_id")
	parsedAccountID, err := strconv.ParseInt(accountID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid account id"})
		return
	}
	account, err := svc.service.GetAccountBalance(c, parsedAccountID)
	if err == models.ErrAccountNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	accountBalanceResponse.ID = account.GetID()
	accountBalanceResponse.Balance = account.GetBalance()
	accountBalanceResponse.DocumentNumber = account.GetDocumentNumber()

	c.JSON(http.StatusOK, accountBalanceResponse)
}
