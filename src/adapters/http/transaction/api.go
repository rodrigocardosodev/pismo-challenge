package transaction

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/dtos"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/models"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/services"
)

type HTTPTransactionAdapter struct {
	service services.ITrasactionService
}

func NewHTTPTransactionAdapter(service services.ITrasactionService) *HTTPTransactionAdapter {
	return &HTTPTransactionAdapter{service: service}
}

func (svc *HTTPTransactionAdapter) CreateTransaction(c *gin.Context) {
	var transactionRequest dtos.TransactionRequest
	if err := c.ShouldBindJSON(&transactionRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}
	account, err := svc.service.Create(c, transactionRequest.AccountID, transactionRequest.OperationTypeID, transactionRequest.Amount)
	if err == models.ErrInvalidOperationType || err == models.ErrInvalidAmountByOperationType {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err == models.ErrAccountNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, account)
}
