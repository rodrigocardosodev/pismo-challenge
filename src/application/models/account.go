package models

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

var (
	ErrAccountNotFound      = errors.New("account not found")
	ErrAccountAlreadyExists = errors.New("account already exists")
	ErrInvalidAccountID     = errors.New("invalid account id")
	ErrInvalidBodyAccount   = errors.New("invalid request body account")

	ErrCpfMustHave11Digits   = errors.New("cpf must have 11 digits")
	ErrCpfMustHaveOnlyDigits = errors.New("cpf must have only digits")
	ErrInvalidCpf            = errors.New("invalid cpf")
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type AccountInterface interface {
	GetID() int64
	GetDocumentNumber() string
	SetID(int64)
	IsValid() error
	SetBalance(float64)
	GetBalance() float64
}

type Account struct {
	ID             int64   `json:"id" valid:"-"`
	DocumentNumber string  `json:"document_number" valid:"required"`
	Balance        float64 `json:"balance,omitempty" valid:"-"`
}

func (a *Account) GetID() int64 {
	return a.ID
}

func (a *Account) GetDocumentNumber() string {
	return a.DocumentNumber
}

func (a *Account) SetID(id int64) {
	a.ID = id
}

func (a *Account) IsValid() error {
	_, err := govalidator.ValidateStruct(a)
	if err != nil {
		return err
	}

	return nil
}

func (a *Account) SetBalance(balance float64) {
	a.Balance = balance
}

func (a *Account) GetBalance() float64 {
	return a.Balance
}

func NewAccount(documentNumber string) AccountInterface {
	return &Account{
		DocumentNumber: documentNumber,
	}
}
