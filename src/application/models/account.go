package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/rodrigocardosodev/pismo-challenge/src/utils"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type AccountInterface interface {
	GetID() int64
	GetDocumentNumber() string
	GetAmount() int64
	SetAmount(int64)
	SetID(int64)
}

type Account struct {
	ID             int64  `json:"id" valid:"required"`
	DocumentNumber string `json:"document_number" valid:"required"`
	Amount         int64  `json:"amount" valid:"-"`
}

func (a *Account) GetID() int64 {
	return a.ID
}

func (a *Account) GetDocumentNumber() string {
	return a.DocumentNumber
}

func (a *Account) GetAmount() int64 {
	return a.Amount
}

func (a *Account) SetAmount(amount int64) {
	a.Amount = amount
}

func (a *Account) SetID(id int64) {
	a.ID = id
}

func (a *Account) ValidateDocumentNumber() error {
	err := utils.IsValidCPF(a.DocumentNumber)
	if err != nil {
		return err
	}

	return nil
}

func NewAccount(documentNumber string) AccountInterface {
	return &Account{
		DocumentNumber: documentNumber,
		Amount:         0,
	}
}
