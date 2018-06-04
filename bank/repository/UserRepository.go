package repository

import (
	"github.com/jakubjanuzik/bank/account"
	"github.com/jakubjanuzik/bank/user"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type UserRepository struct {
	DB *gorm.DB
}

func (Repository *UserRepository) AddUserToAcc(user *user.User, acc *account.Account) {
	Repository.DB.Model(&user).Association("Accounts").Append(acc)
}
