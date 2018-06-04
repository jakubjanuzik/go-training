package account

// Account Konto
type Account struct {
	ID      int `gorm:"primary_key`
	Number  string
	Balance int
}

func (account *Account) Deposit(funds int) {
	account.Balance += funds
}

func (account *Account) Withdraw(funds int) {
	account.Balance -= funds
}
