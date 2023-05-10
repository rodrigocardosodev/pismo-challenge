package services

import (
	"github.com/rodrigocardosodev/pismo-challenge/src/application/models"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/ports"
)

type AccountService struct {
	AccountRepository ports.IAccountRepository
}

func (a *AccountService) Create(documentNumber string) (models.AccountInterface, error) {
	account := models.NewAccount(documentNumber)

	return a.AccountRepository.Create(account)
}

func (a *AccountService) GetByID(id int64) (models.AccountInterface, error) {
	return a.AccountRepository.GetByID(id)
}
