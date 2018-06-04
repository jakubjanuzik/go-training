package service

import (
	"fmt"
)

type ConsoleAccountService struct {
	Service AccountServiceInterface
}

func (consoleService *ConsoleAccountService) CreateAccount() string {
	accountNumber := consoleService.Service.CreateAccount()
	fmt.Printf("Created new account: %v", accountNumber)

	return accountNumber
}

func (consoleService *ConsoleAccountService) DepositFunds(number string, funds int) error {
	err := consoleService.Service.DepositFunds(number, funds)
	if err != nil {
		fmt.Printf("Error while depositing funds to %v\n", number)
		return err
	}

	fmt.Printf("%v deposited successfully to %v\n", funds, number)
	return nil
}

func (consoleService *ConsoleAccountService) WithdrawFunds(number string, funds int) error {
	err := consoleService.Service.WithdrawFunds(number, funds)
	if err != nil {
		fmt.Printf("Error while withdrawing funds from %v\n", number)
		return err
	}

	fmt.Printf("%v withdrawn successfully from %v\n", funds, number)
	return nil
}

func (consoleService *ConsoleAccountService) PrintReport() {
	consoleService.Service.PrintReport()
}
