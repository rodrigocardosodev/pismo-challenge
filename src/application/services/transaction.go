package services

import (
	"context"

	"github.com/rodrigocardosodev/pismo-challenge/src/application/models"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/ports"
)

type TransactionService struct {
	TransactionRepository ports.ITransactionRepository
}

type ITrasactionService interface {
	Create(ctx context.Context, accountId int64, operationId uint8, amount float64) (models.TransactionInterface, error)
}

func NewTransactionService(transactionRepository ports.ITransactionRepository) ITrasactionService {
	return &TransactionService{
		TransactionRepository: transactionRepository,
	}
}

func (t *TransactionService) Create(ctx context.Context, accountId int64, operationId uint8, amount float64) (models.TransactionInterface, error) {

	transaction := models.NewTransaction(accountId, operationId, amount)
	err := transaction.IsValid()
	if err != nil {
		return nil, err
	}

	switch operationId {
	case models.COMPRA_A_VISTA:
		if amount < 0 {
			transaction.SetAmount(amount)
		} else {
			return nil, models.ErrInvalidAmountByOperationType
		}
	case models.COMPRA_PARCELADA:
		if amount < 0 {
			transaction.SetAmount(amount)
		} else {
			return nil, models.ErrInvalidAmountByOperationType
		}
	case models.SAQUE:
		if amount < 0 {
			transaction.SetAmount(amount)
		} else {
			return nil, models.ErrInvalidAmountByOperationType
		}
	case models.PAGAMENTO:
		if amount > 0 {
			transaction.SetAmount(amount)
		} else {
			return nil, models.ErrInvalidAmountByOperationType
		}
	default:
		return nil, models.ErrInvalidOperationType
	}

	transaction, err = t.TransactionRepository.Create(ctx, transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
