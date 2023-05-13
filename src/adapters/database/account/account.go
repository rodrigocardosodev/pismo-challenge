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
	tx, err := a.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var account models.Account
	err = a.DB.QueryRowContext(ctx, "SELECT id, document_number FROM accounts WHERE id = $1", id).Scan(&account.ID, &account.DocumentNumber)
	if err == sql.ErrNoRows {
		return nil, models.ErrAccountNotFound
	}
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &account, nil
}

func (a *AccountRepository) GetByDocumentNumber(ctx context.Context, documentNumber string) (models.AccountInterface, error) {
	tx, err := a.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var account models.Account
	err = a.DB.QueryRowContext(ctx, "SELECT id, document_number FROM accounts WHERE document_number = $1", documentNumber).Scan(&account.ID, &account.DocumentNumber)
	// Verify if errors is not found account
	if err == sql.ErrNoRows {
		return nil, models.ErrAccountNotFound
	}
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &account, nil
}

func (a *AccountRepository) Create(ctx context.Context, account models.AccountInterface) (models.AccountInterface, error) {
	tx, err := a.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	var id int64
	err = a.DB.QueryRowContext(ctx, "INSERT INTO accounts (document_number) VALUES ($1) RETURNING id", account.GetDocumentNumber()).Scan(&id)

	if err != nil {
		return nil, err
	}

	account.SetID(id)

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return account, nil
}

func NewAccountRepository(db *sql.DB) ports.IAccountRepository {
	return &AccountRepository{
		DB: db,
	}
}
