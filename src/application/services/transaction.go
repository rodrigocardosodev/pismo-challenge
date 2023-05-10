package services

import (
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

func (t *TransactionService) Create(accountId int64, operationId int, amount int64) (models.TransactionInterface, error) {
	transaction := models.NewTransaction(accountId, operationId, amount)
	err := transaction.IsValid()
	if err != nil {
		return nil, err
	}

	if operationId == models.SAQUE || operationId == models.COMPRA_A_VISTA {
		transaction.SetAmount(-amount)
	}

	if operationId == models.PAGAMENTO {
		transaction.SetAmount(amount)
	}

	transaction, err = t.TransactionRepository.Create(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
