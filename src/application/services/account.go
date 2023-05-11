package services

import (
	"context"

	"github.com/rodrigocardosodev/pismo-challenge/src/application/models"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/ports"
)

type AccountService struct {
	AccountRepository ports.IAccountRepository
}

func (a *AccountService) Create(ctx context.Context, documentNumber string) (models.AccountInterface, error) {
	account := models.NewAccount(documentNumber)
	err := account.IsValid()
	if err != nil {
		return nil, err
	}

	return a.AccountRepository.Create(ctx, account)
}

func (a *AccountService) GetByID(ctx context.Context, id int64) (models.AccountInterface, error) {
	return a.AccountRepository.GetByID(ctx, id)
}

func NewAccountService(accountRepository ports.IAccountRepository) *AccountService {
	return &AccountService{
		AccountRepository: accountRepository,
	}
}
