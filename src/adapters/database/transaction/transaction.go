package transaction

import (
	"context"
	"database/sql"

	"github.com/rodrigocardosodev/pismo-challenge/src/application/models"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/ports"
)

type TransactionRepository struct {
	DB *sql.DB
}

func (a *TransactionRepository) Create(ctx context.Context, transaction models.TransactionInterface) (models.TransactionInterface, error) {
	var id int64
	var event_date string
	err := a.DB.QueryRow("INSERT INTO transactions (account_id, operation_id, amount) VALUES ($1, $2, $3) RETURNING id, event_date", transaction.GetAccountID(), transaction.GetOperationTypeId(), transaction.GetAmount()).Scan(&id, &event_date)
	if err != nil {
		return nil, err
	}

	transaction.SetID(id)
	transaction.SetEventDate(event_date)

	return transaction, nil
}

func NewTransactionRepository(db *sql.DB) ports.ITransactionRepository {
	return &TransactionRepository{
		DB: db,
	}
}