package account

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/dtos"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/models"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/services"
)

type HTTPAccountAdapter struct {
	service *services.AccountService
}

func NewHTTPAccountAdapter(service *services.AccountService) *HTTPAccountAdapter {
	return &HTTPAccountAdapter{service: service}
}

func (svc *HTTPAccountAdapter) CreateAccount(c *gin.Context) {
	var accountRequest dtos.AccountRequest
	if err := c.ShouldBindJSON(&accountRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	account, err := svc.service.Create(c, accountRequest.DocumentNumber)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
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
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, account)
}
