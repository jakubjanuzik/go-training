package repository

import (
	"database/sql"

	"github.com/jakubjanuzik/bank/account"
)

type MySqlAccountRepository struct {
	Db *sql.DB
}

func (Repository *MySqlAccountRepository) GetByNumber(number string) (*account.Account, error) {
	var (
		accountID      int
		accountNumber  string
		accountBalance int
	)

	rows, _ := Repository.Db.Query("select * from account where account.number = ?", number)

	for rows.Next() {
		err := rows.Scan(&accountID, &accountNumber, &accountBalance)
		if err != nil {
			return nil, err
		}
	}
	account := account.Account{ID: accountID, Balance: accountBalance, Number: accountNumber}
	return &account, nil
}

func (Repository *MySqlAccountRepository) Save(account *account.Account) error {
	tx, err := Repository.Db.Begin()
	stmt, _ := tx.Prepare("insert into account values(null,?,?)")
	response, _ := tx.Stmt(stmt).Exec(account.Number, account.Balance)

	_, err = response.LastInsertId()

	if err != nil {
		return err
	}
	tx.Commit()

	return nil
}

func (Repository *MySqlAccountRepository) Update(account *account.Account) error {
	tx, err := Repository.Db.Begin()
	stmt, _ := tx.Prepare("UPDATE account SET number = ?, balance = ? WHERE account.id = ?")

	response, _ := tx.Stmt(stmt).Exec(account.Number, account.Balance, account.ID)
	tx.Commit()

	_, err = response.LastInsertId()

	if err != nil {
		return err
	}

	return nil
}

func (Repository *MySqlAccountRepository) GetAll() ([]account.Account, error) {
	var (
		accountID      int
		accountNumber  string
		accountBalance int
	)

	accounts := make([]account.Account, 0)

	rows, _ := Repository.Db.Query("select * from account")

	for rows.Next() {
		err := rows.Scan(&accountID, &accountNumber, &accountBalance)
		if err != nil {
			return nil, err
		}
		account := account.Account{ID: accountID, Balance: accountBalance, Number: accountNumber}
		accounts = append(accounts, account)
	}
	return accounts, nil
}
