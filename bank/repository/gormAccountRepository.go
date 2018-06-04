package repository

import (
	"github.com/jakubjanuzik/bank/account"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type GormAccountRepository struct {
	DB *gorm.DB
}

func (Repository *GormAccountRepository) GetByNumber(number string) (*account.Account, error) {
	var acc account.Account
	Repository.DB.Where("number = ?", number).First(&acc)
	return &acc, nil
}

func (Repository *GormAccountRepository) Save(account *account.Account) error {
	Repository.DB.Create(&account)
	return nil
}

func (Repository *GormAccountRepository) Update(accountObj *account.Account) error {
	// var acc account.Account
	// Repository.DB.First(&acc, accountObj.ID)
	// fmt.Printf("^^^ %v", accountObj.Balance)
	Repository.DB.Model(accountObj).Update("balance", accountObj.Balance)
	return nil
}

func (Repository *GormAccountRepository) GetAll() ([]account.Account, error) {
	accounts := make([]account.Account, 0)
	Repository.DB.Find(&accounts)
	return accounts, nil
}
