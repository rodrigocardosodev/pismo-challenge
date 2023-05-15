package services

import (
	"context"
	"strconv"
	"strings"

	"github.com/rodrigocardosodev/pismo-challenge/src/application/models"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/ports"
)

type AccountService struct {
	AccountRepository     ports.IAccountRepository
	TransactionRepository ports.ITransactionRepository
}

type IAccountService interface {
	Create(ctx context.Context, documentNumber string) (models.AccountInterface, error)
	GetByID(ctx context.Context, id int64) (models.AccountInterface, error)
	GetAccountBalance(ctx context.Context, accountID int64) (models.AccountInterface, error)
}

func (a *AccountService) Create(ctx context.Context, documentNumber string) (models.AccountInterface, error) {
	err := a.isValidCPF(documentNumber)
	if err != nil {
		return nil, err
	}
	accountExistsWithDocument, _ := a.AccountRepository.GetByDocumentNumber(ctx, documentNumber)
	if accountExistsWithDocument != nil {
		return nil, models.ErrAccountAlreadyExists
	}

	account := models.NewAccount(documentNumber)
	err = account.IsValid()
	if err != nil {
		return nil, err
	}

	result, err := a.AccountRepository.Create(ctx, account)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *AccountService) GetByID(ctx context.Context, id int64) (models.AccountInterface, error) {
	return a.AccountRepository.GetByID(ctx, id)
}

func (a *AccountService) GetAccountBalance(ctx context.Context, accountID int64) (models.AccountInterface, error) {
	account, err := a.AccountRepository.GetByID(ctx, accountID)
	if err != nil {
		return nil, err
	}

	balance, err := a.TransactionRepository.GetBalanceByAccountID(ctx, accountID)
	if err != nil {
		return nil, err
	}

	account.SetBalance(balance)

	return account, nil
}

func (a *AccountService) isValidCPF(cpf string) error {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	if len(cpf) != 11 {
		return models.ErrCpfMustHave11Digits
	}

	for _, char := range cpf {
		if char < '0' || char > '9' {
			return models.ErrCpfMustHaveOnlyDigits
		}
	}

	if a.allEqual(cpf) {
		return models.ErrInvalidCpf
	}

	firstDigit := a.calculateDigit(cpf[:9])
	secondDigit := a.calculateDigit(cpf[:10])

	if cpf[9] != firstDigit || cpf[10] != secondDigit {
		return models.ErrInvalidCpf
	}

	return nil
}

func (a *AccountService) allEqual(cpf string) bool {
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			return false
		}
	}
	return true
}

func (a *AccountService) calculateDigit(cpf string) byte {
	sum := 0
	multiplier := len(cpf) + 1

	for _, char := range cpf {
		digit, _ := strconv.Atoi(string(char))
		sum += digit * multiplier
		multiplier--
	}

	remainder := sum % 11
	if remainder < 2 {
		return '0'
	}

	return byte(11 - remainder + '0')
}

func NewAccountService(accountRepository ports.IAccountRepository, transactionRepository ports.ITransactionRepository) IAccountService {
	return &AccountService{
		AccountRepository:     accountRepository,
		TransactionRepository: transactionRepository,
	}
}
