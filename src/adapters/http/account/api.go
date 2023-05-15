package account

import (
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

func (svc *HTTPAccountAdapter) CreateAccount(c *gin.Context) {
	var accountRequest dtos.AccountRequest
	if err := c.ShouldBindJSON(&accountRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	account, err := svc.service.Create(c, accountRequest.DocumentNumber)
	if err == models.ErrCpfMustHaveOnlyDigits || err == models.ErrCpfMustHave11Digits || err == models.ErrInvalidCpf || err == models.ErrAccountAlreadyExists {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, account)
}

func (svc *HTTPAccountAdapter) GetAccountById(c *gin.Context) {
	accountID := c.Param("account_id")
	parsedAccountID, err := strconv.ParseInt(accountID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid account id"})
		return
	}
	account, err := svc.service.GetByID(c, parsedAccountID)
	if err == models.ErrAccountNotFound {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, account)
}

func (svc *HTTPAccountAdapter) GetAccountBalance(c *gin.Context) {
	var accountBalanceResponse dtos.AccountBalanceResponse
	accountID := c.Param("account_id")
	parsedAccountID, err := strconv.ParseInt(accountID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid account id"})
		return
	}
	account, err := svc.service.GetAccountBalance(c, parsedAccountID)
	if err == models.ErrAccountNotFound {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	accountBalanceResponse.ID = account.GetID()
	accountBalanceResponse.Balance = account.GetBalance()
	accountBalanceResponse.DocumentNumber = account.GetDocumentNumber()

	c.JSON(200, accountBalanceResponse)
}
