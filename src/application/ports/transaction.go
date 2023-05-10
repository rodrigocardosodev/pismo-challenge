package ports

import (
	"github.com/rodrigocardosodev/pismo-challenge/src/application/models"
)

type ITransactionWriter interface {
	Create(Transaction models.TransactionInterface) (models.TransactionInterface, error)
}

type ITransactionRepository interface {
	ITransactionWriter
}
