package models

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

const (
	COMPRA_A_VISTA   = 1
	COMPRA_PARCELADA = 2
	SAQUE            = 3
	PAGAMENTO        = 4
)

type TransactionInterface interface {
	GetID() int64
	GetAccountID() int64
	GetOperationTypeId() int64
	GetAmount() int64
	SetID(int64)
	SetOperationTypeId(int64)
	SetAmount(int64)
	SetEventDate(string)
	IsValid() error
}

type Transaction struct {
	ID              int64  `json:"id" valid:"-"`
	AccountId       int64  `json:"account_id" valid:"required"`
	OperationTypeId int64  `json:"operation_type_id" valid:"required,numeric"`
	Amount          int64  `json:"amount" valid:"required,numeric"`
	EventDate       string `json:"event_date" valid:"-"`
}

func (t *Transaction) GetID() int64 {
	return t.ID
}

func (t *Transaction) GetAccountID() int64 {
	return t.AccountId
}

func (t *Transaction) GetOperationTypeId() int64 {
	return t.OperationTypeId
}

func (t *Transaction) GetAmount() int64 {
	return t.Amount
}

func (t *Transaction) SetAmount(amount int64) {
	t.Amount = amount
}

func (t *Transaction) SetOperationTypeId(operationTypeId int64) {
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

func NewTransaction(accountId int64, operationId int, amount int64) TransactionInterface {
	return &Transaction{
		AccountId:       accountId,
		OperationTypeId: int64(operationId),
		Amount:          amount,
	}
}
