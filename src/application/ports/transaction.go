package ports

import (
	"context"

	"github.com/rodrigocardosodev/pismo-challenge/src/application/models"
)

type ITransactionWriter interface {
	Create(ctx context.Context, Transaction models.TransactionInterface) (models.TransactionInterface, error)
}

type ITransactionRepository interface {
	ITransactionWriter
}
