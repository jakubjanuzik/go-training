package main

import (
	"fmt"
)

type money float32

type account struct {
	number  string
	balance money
}

func (acc *account) deposit(funds money) {
	acc.balance += funds
}

func main9() {
	myAccount := account{"0001", 12}
	otherAccount := account{balance: 0, number: "1234"}
	myAccount.balance = 333

	// yetAnotherAccount := account{balance: 2000, number: "3000"}

	yetAnotherAccount := otherAccount
	yetAnotherAccount.balance = 213213213

	otherAccount.deposit(1000)

	user := struct {
		name string
	}{"Kowalski"}

	fmt.Println(user)
	fmt.Println(myAccount)
	fmt.Println(otherAccount)
	fmt.Println(yetAnotherAccount)
}
