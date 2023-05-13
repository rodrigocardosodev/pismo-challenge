package models

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

var (
	ErrAccountNotFound      = errors.New("account not found")
	ErrAccountAlreadyExists = errors.New("account already exists")

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
}

type Account struct {
	ID             int64  `json:"id" valid:"-"`
	DocumentNumber string `json:"document_number" valid:"required"`
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

func NewAccount(documentNumber string) AccountInterface {
	return &Account{
		DocumentNumber: documentNumber,
	}
}
