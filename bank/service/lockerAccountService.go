package service

import (
	"sync"
)

type LockerAccountService struct {
	Service AccountServiceInterface
	Locker  sync.RWMutex
}

func (lockerService *LockerAccountService) CreateAccount() string {
	lockerService.Locker.Lock()
	accountNumber := lockerService.Service.CreateAccount()
	lockerService.Locker.Unlock()

	return accountNumber
}

func (lockerService *LockerAccountService) DepositFunds(number string, funds int) error {
	lockerService.Locker.Lock()
	err := lockerService.Service.DepositFunds(number, funds)
	lockerService.Locker.Unlock()
	return err
}

func (lockerService *LockerAccountService) WithdrawFunds(number string, funds int) error {
	lockerService.Locker.Lock()
	err := lockerService.Service.WithdrawFunds(number, funds)
	lockerService.Locker.Unlock()
	return err
}

func (lockerService *LockerAccountService) PrintReport() {
	lockerService.Locker.RLock()
	lockerService.Service.PrintReport()
	lockerService.Locker.RUnlock()
}
