package main

var address *int

func main6() {
	age := 40
	println(age)

	println(&age)

	address = &age

	println(address)

	*address = 39

	println(*address)
	println(age)
}
