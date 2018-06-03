package main

import (
	"fmt"
)

var fullName = "Jan Kowalski"
var firstName, lastName = "Mark", "Nowak"
var age int

func main1() {
	age = 21

	var score = 40
	text := "Go lang"

	fmt.Printf("%v %v", score, text)
}
