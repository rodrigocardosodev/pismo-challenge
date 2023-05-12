package account

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/dtos"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/services"
)

type HTTPTransactionAdapter struct {
	service *services.TransactionService
}

func NewHTTPTransactionAdapter(service *services.TransactionService) *HTTPTransactionAdapter {
	return &HTTPTransactionAdapter{service: service}
}

func (svc *HTTPTransactionAdapter) CreateTransaction(c *gin.Context) {
	var transactionRequest dtos.TransactionRequest
	if err := c.ShouldBindJSON(&transactionRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	account, err := svc.service.Create(c, transactionRequest.AccountID, transactionRequest.OperationTypeID, transactionRequest.Amount)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, account)
}
