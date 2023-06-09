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
	tx, err := a.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var id int64
	var event_date string
	err = a.DB.QueryRowContext(ctx, "INSERT INTO transactions (account_id, operation_id, amount) VALUES ($1, $2, $3) RETURNING id, event_date", transaction.GetAccountID(), transaction.GetOperationTypeId(), transaction.GetAmount()).Scan(&id, &event_date)
	if err != nil {
		if err.Error() == "pq: account not found" {
			return nil, models.ErrAccountNotFound
		}
		return nil, err
	}

	transaction.SetID(id)
	transaction.SetEventDate(event_date)

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return transaction, nil
}

func (a *TransactionRepository) GetBalanceByAccountID(ctx context.Context, accountId int64) (float64, error) {
	tx, err := a.DB.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var balance float64
	err = a.DB.QueryRowContext(ctx, "SELECT COALESCE(SUM(amount), 0) FROM transactions WHERE account_id = $1", accountId).Scan(&balance)
	if err == sql.ErrNoRows {
		return 0, models.ErrAccountNotFound
	}

	if err != nil {
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return balance, nil
}

func NewTransactionRepository(db *sql.DB) ports.ITransactionRepository {
	return &TransactionRepository{
		DB: db,
	}
}
