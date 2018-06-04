package generator

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"

	"github.com/jakubjanuzik/bank/account"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type GormAccountRepository struct {
	DB *gorm.DB
}

func (generator *GormAccountRepository) Next() string {
	var acc account.Account
	generator.DB.Last(&acc)
	number, _ := strconv.Atoi(acc.Number)
	number++
	return fmt.Sprintf("%026d", number)
}
