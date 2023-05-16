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

	isNegativeOperation := operationId == models.COMPRA_A_VISTA || operationId == models.COMPRA_PARCELADA || operationId == models.SAQUE
	isPositiveOperation := operationId == models.PAGAMENTO

	if isNegativeOperation && amount < 0 {
		transaction.SetAmount(amount)
	} else if isPositiveOperation && amount > 0 {
		transaction.SetAmount(amount)
	} else if !isNegativeOperation && !isPositiveOperation {
		return nil, models.ErrInvalidOperationType
	} else {
		return nil, models.ErrInvalidAmountByOperationType
	}

	transaction, err = t.TransactionRepository.Create(ctx, transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
