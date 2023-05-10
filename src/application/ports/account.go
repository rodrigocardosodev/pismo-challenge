package ports

import (
	"github.com/rodrigocardosodev/pismo-challenge/src/application/models"
)

type IAccountReader interface {
	GetByID(id int64) (models.AccountInterface, error)
}

type IAccountWriter interface {
	Create(account models.AccountInterface) (models.AccountInterface, error)
}

type IAccountRepository interface {
	IAccountReader
	IAccountWriter
}
