package account

import (
	"context"
	"database/sql"

	"github.com/rodrigocardosodev/pismo-challenge/src/application/models"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/ports"
)

type AccountRepository struct {
	DB *sql.DB
}

func (a *AccountRepository) GetByID(ctx context.Context, id int64) (models.AccountInterface, error) {
	var account models.Account
	err := a.DB.QueryRowContext(ctx, "SELECT id, document_number, amount FROM accounts WHERE id = $1", id).Scan(&account.ID, &account.DocumentNumber, &account.Amount)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (a *AccountRepository) Create(ctx context.Context, account models.AccountInterface) (models.AccountInterface, error) {
	var id int64
	err := a.DB.QueryRowContext(ctx, "INSERT INTO accounts (document_number, amount) VALUES ($1, $2) RETURNING id", account.GetDocumentNumber(), account.GetAmount()).Scan(&id)
	if err != nil {
		return nil, err
	}

	account.SetID(id)

	return account, nil
}

func NewAccountRepository(db *sql.DB) ports.IAccountRepository {
	return &AccountRepository{
		DB: db,
	}
}
