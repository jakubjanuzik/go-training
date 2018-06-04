package user

import (
	"github.com/jakubjanuzik/bank/account"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Accounts  []account.Account `gorm:"many2many:user_accounts;"`
}
