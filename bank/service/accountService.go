package service

import (
	"fmt"

	"github.com/jakubjanuzik/bank/account"
	"github.com/jakubjanuzik/bank/generator"
	"github.com/jakubjanuzik/bank/repository"
	"github.com/jakubjanuzik/bank/user"
)

type AccountServiceInterface interface {
	CreateAccount() string
	DepositFunds(number string, funds int) error
	WithdrawFunds(number string, funds int) error
	PrintReport()
}

// AccountService serwis
type AccountService struct {
	Repository     repository.AccountRepository
	Generator      generator.AccountNumberGenerator
	UserRepository repository.UserRepository
}

// CreateAccount tworzymy konto
func (accountService *AccountService) CreateAccount() string {
	accountNumber := accountService.Generator.Next()
	accountService.Repository.Save(&account.Account{Number: accountNumber})

	return accountNumber
}

// DepositFunds depozytujemy
func (accountService *AccountService) DepositFunds(number string, funds int) error {
	return accountService.process(number, func(account *account.Account) {
		account.Deposit(funds)
	})
}

// WithdrawFunds wyciagamy
func (accountService *AccountService) WithdrawFunds(number string, funds int) error {
	return accountService.process(number, func(account *account.Account) {
		account.Withdraw(funds)
	})
}

func (accountService *AccountService) process(number string, callback func(account *account.Account)) error {
	account, err := accountService.Repository.GetByNumber(number)
	if err != nil {
		return err
	}
	callback(account)
	accountService.Repository.Update(account)
	return nil
}

// PrintReport report
func (accountService *AccountService) PrintReport() {
	accounts, _ := accountService.Repository.GetAll()

	for _, account := range accounts {
		fmt.Printf("%v %v\n", account.Number, account.Balance)
	}
}

func (accountService *AccountService) AddUserToAcc(user *user.User, acc *account.Account) {
	accountService.UserRepository.AddUserToAcc(user, acc)
}
