package ports

import (
	"context"

	"github.com/rodrigocardosodev/pismo-challenge/src/application/models"
)

type IAccountReader interface {
	GetByID(ctx context.Context, id int64) (models.AccountInterface, error)
	GetByDocumentNumber(ctx context.Context, documentNumber string) (models.AccountInterface, error)
}

type IAccountWriter interface {
	Create(ctx context.Context, account models.AccountInterface) (models.AccountInterface, error)
}

type IAccountRepository interface {
	IAccountReader
	IAccountWriter
}
