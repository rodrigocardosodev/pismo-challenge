package database

import (
	"database/sql"

	"github.com/rodrigocardosodev/pismo-challenge/src/application/models"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/ports"
)

type AccountRepository struct {
	DB *sql.DB
}

func (a *AccountRepository) GetByID(id int64) (models.AccountInterface, error) {
	var account models.Account
	err := a.DB.QueryRow("SELECT id, document_number, amount FROM accounts WHERE id = $1", id).Scan(&account.ID, &account.DocumentNumber, &account.Amount)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (a *AccountRepository) Create(account models.AccountInterface) (models.AccountInterface, error) {
	var id int64
	err := a.DB.QueryRow("INSERT INTO accounts (document_number, amount) VALUES ($1, $2) RETURNING id", account.GetDocumentNumber(), account.GetAmount()).Scan(&id)
	if err != nil {
		return nil, err
	}

	account.SetID(id)

	return account, nil
}

func (a *AccountRepository) UpdateAmount(id int64, amount int64) (models.AccountInterface, error) {
	account, err := a.GetByID(id)
	if err != nil {
		return nil, err
	}

	account.SetAmount(amount)

	_, err = a.DB.Exec("UPDATE accounts SET amount = $1 WHERE id = $2", account.GetAmount(), account.GetID())
	if err != nil {
		return nil, err
	}

	return account, nil
}

func NewAccountRepository(db *sql.DB) ports.IAccountRepository {
	return &AccountRepository{
		DB: db,
	}
}
