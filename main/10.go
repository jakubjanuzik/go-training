package main

import (
	"fmt"
)

type user struct {
	name string
	mail string
}

type admin struct {
	user
	level int
}

func (user *user) printInfo() {
	fmt.Printf("%v, %v", user.name, user.mail)
}

func (admin *admin) printInfo() {
	fmt.Printf("ADMIN:  %v, %v", admin.name, admin.mail)
}

func main10() {
	myAdmin := admin{
		user:  user{"Jan", "Kowalski@gmail.com"},
		level: 1,
	}

	myAdmin.printInfo()
	myAdmin.user.printInfo()
}
