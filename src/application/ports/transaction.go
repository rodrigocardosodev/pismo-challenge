package ports

import (
	"context"

	"github.com/rodrigocardosodev/pismo-challenge/src/application/models"
)

type ITransactionWriter interface {
	Create(ctx context.Context, Transaction models.TransactionInterface) (models.TransactionInterface, error)
}

type ITransactionReader interface {
	GetBalanceByAccountID(ctx context.Context, accountId int64) (float64, error)
}

type ITransactionRepository interface {
	ITransactionWriter
	ITransactionReader
}
