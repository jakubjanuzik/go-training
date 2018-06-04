package repository

import (
	"errors"

	"github.com/jakubjanuzik/bank/account"
)

// AccountRepository Repozytorium
type AccountRepository interface {
	GetByNumber(number string) (*account.Account, error)

	Save(account *account.Account) error

	Update(account *account.Account) error

	GetAll() ([]account.Account, error)
}

// MapAccountRepository mapy
type MapAccountRepository struct {
	Accounts map[string]*account.Account
}

func (repository *MapAccountRepository) GetByNumber(number string) (*account.Account, error) {
	account, found := repository.Accounts[number]
	if !found {
		return nil, errors.New("Account not found")
	}
	return account, nil
}

func (repository *MapAccountRepository) Save(account *account.Account) error {
	repository.Accounts[account.Number] = account
	return nil
}

func (repository *MapAccountRepository) Update(account *account.Account) error {
	return repository.Save(account)
}

func (repository *MapAccountRepository) GetAll() ([]account.Account, error) {
	accounts := make([]account.Account, 0, len(repository.Accounts))
	for _, account := range repository.Accounts {
		accounts = append(accounts, *account)
	}

	return accounts, nil
}
