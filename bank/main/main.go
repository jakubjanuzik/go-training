package main

import (
	"sync"

	"github.com/jakubjanuzik/bank/account"
	"github.com/jakubjanuzik/bank/rest"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/jakubjanuzik/bank/user"
)

var group sync.WaitGroup

func InitDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "bank.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&account.Account{})
	db.AutoMigrate(&user.User{})
	return db
}

func main() {
	rest.Start()
	// db, _ := sql.Open("mysql", "root:root@/go")
	// defer db.Close()
	// repository := bank.MapAccountRepository{Accounts: make(map[string]*bank.Account)}
	// generator := bank.InMemoryAccountNumberGenerator{}
	// repository := repository.MySqlAccountRepository{Db: db}
	// generator := generator.DatabaseNumberGenerator{Db: db}
	// var Db = InitDb()
	// accRepo := repository.GormAccountRepository{DB: Db}
	// generator := generator.GormAccountRepository{DB: Db}
	// userRepo := repository.UserRepository{DB: Db}

	// accountService := service.AccountService{Repository: &accRepo, Generator: &generator, UserRepository: userRepo}
	// consoleService := service.ConsoleAccountService{Service: &accountService}
	// lockerService := service.LockerAccountService{Service: &consoleService, Locker: sync.RWMutex{}}
	// accountNumber := accountService.CreateAccount()
	// println(accountNumber)
	// accountService.DepositFunds(accountNumber, 1000)
	// accountService.WithdrawFunds(accountNumber, 500)
	// var myUser user.User
	// Db.Create(&user.User{FirstName: "Kuba", LastName: "Januzik"})
	// Db.First(&myUser)
	// acc, _ := accountService.Repository.GetByNumber(accountNumber)
	// accountService.AddUserToAcc(&myUser, acc)
	// accountService.WithdrawFunds(accountNumber, 500)
	// group.Add(5)
	// go func() {
	// 	for counter := 0; counter < 10000; counter++ {
	// 		lockerService.DepositFunds(accountNumber, 1000)
	// 	}
	// 	group.Done()

	// }()
	// go func() {
	// 	for counter := 0; counter < 10000; counter++ {
	// 		lockerService.DepositFunds(accountNumber, 100)
	// 	}
	// 	group.Done()
	// }()

	// go func() {
	// 	for counter := 0; counter < 10000; counter++ {
	// 		lockerService.WithdrawFunds(accountNumber, 300)
	// 	}
	// 	group.Done()
	// }()

	// go func() {
	// 	for counter := 0; counter < 10000; counter++ {
	// 		lockerService.WithdrawFunds(accountNumber, 10)
	// 	}
	// 	group.Done()
	// }()

	// go func() {
	// 	for counter := 0; counter < 10000; counter++ {
	// 		lockerService.WithdrawFunds(accountNumber, 50)
	// 	}
	// 	group.Done()
	// }()

	// group.Wait()
	// lockerService.PrintReport()
}
