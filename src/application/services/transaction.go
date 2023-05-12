package services

import (
	"context"

	"github.com/rodrigocardosodev/pismo-challenge/src/application/models"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/ports"
)

type TransactionService struct {
	TransactionRepository ports.ITransactionRepository
}

func NewTransactionService(transactionRepository ports.ITransactionRepository) *TransactionService {
	return &TransactionService{
		TransactionRepository: transactionRepository,
	}
}

func (t *TransactionService) Create(ctx context.Context, accountId int64, operationId int8, amount uint64) (models.TransactionInterface, error) {

	transaction := models.NewTransaction(accountId, operationId, amount)
	err := transaction.IsValid()
	if err != nil {
		return nil, err
	}

	switch operationId {
	case models.SAQUE & models.COMPRA_A_VISTA & models.COMPRA_PARCELADA:
		transaction.SetAmount(amount)
	case models.COMPRA_A_VISTA:
		transaction.SetAmount(amount)
	case models.COMPRA_PARCELADA:
		transaction.SetAmount(amount)
	case models.PAGAMENTO:
		transaction.SetAmount(amount)
	default:
		return nil, models.ErrInvalidOperation
	}

	transaction, err = t.TransactionRepository.Create(ctx, transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
