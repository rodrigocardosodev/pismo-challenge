package models

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

const (
	COMPRA_A_VISTA   = 1
	COMPRA_PARCELADA = 2
	SAQUE            = 3
	PAGAMENTO        = 4
)

var (
	ErrInvalidOperationType         = errors.New("invalid operation type")
	ErrInvalidAmountByOperationType = errors.New("invalid amount by operation type")
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type TransactionInterface interface {
	GetID() int64
	GetAccountID() int64
	GetOperationTypeId() int8
	GetAmount() float64
	SetID(int64)
	SetOperationTypeId(int8)
	SetAmount(float64)
	SetEventDate(string)
	IsValid() error
}

type Transaction struct {
	ID              int64   `json:"id" valid:"-"`
	AccountId       int64   `json:"account_id" valid:"required"`
	OperationTypeId int8    `json:"operation_type_id" valid:"required,numeric"`
	Amount          float64 `json:"amount" valid:"required"`
	EventDate       string  `json:"event_date" valid:"-"`
}

func (t *Transaction) GetID() int64 {
	return t.ID
}

func (t *Transaction) GetAccountID() int64 {
	return t.AccountId
}

func (t *Transaction) GetOperationTypeId() int8 {
	return t.OperationTypeId
}

func (t *Transaction) GetAmount() float64 {
	return t.Amount
}

func (t *Transaction) SetAmount(amount float64) {
	t.Amount = amount
}

func (t *Transaction) SetOperationTypeId(operationTypeId int8) {
	t.OperationTypeId = operationTypeId
}

func (t *Transaction) SetID(id int64) {
	t.ID = id
}

func (t *Transaction) SetEventDate(eventDate string) {
	t.EventDate = eventDate
}

func (t *Transaction) IsValid() error {
	_, err := govalidator.ValidateStruct(t)
	if err != nil {
		return err
	}

	return nil
}

func NewTransaction(accountId int64, operationId int8, amount float64) TransactionInterface {
	return &Transaction{
		AccountId:       accountId,
		OperationTypeId: operationId,
		Amount:          amount,
	}
}
